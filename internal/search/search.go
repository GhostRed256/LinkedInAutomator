package search

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"linkin-automator/internal/browser"
	"log/slog"
)

type SearchOptions struct {
	Keywords string
	Limit    int
}

type Searcher struct {
	Browser *browser.Browser
	Seen    map[string]bool
}

func New(b *browser.Browser) *Searcher {
	return &Searcher{
		Browser: b,
		Seen:    make(map[string]bool),
	}
}

// Run performs the search and returns unique profile URLs
func (s *Searcher) Run(opts SearchOptions) ([]string, error) {
	page := s.Browser.Page
	slog.Info("Starting search", "keywords", opts.Keywords)

	// Construct Search URL
	u := fmt.Sprintf("https://www.linkedin.com/search/results/people/?keywords=%s&origin=SWITCH_SEARCH_VERTICAL", url.QueryEscape(opts.Keywords))

	// Navigate with stealth
	page.MustNavigate(u)
	page.MustWaitLoad()

	// Initial wait for dynamic content
	s.Browser.Stealth.SleepRandom(3*time.Second, 5*time.Second)

	var profiles []string

	// Pagination Loop
	pageCount := 0
	for len(profiles) < opts.Limit {
		pageCount++
		slog.Info("Processing search page", "page", pageCount, "collected", len(profiles))

		// Scroll to load all results (lazy loading)
		slog.Info("Scrolling through results")
		s.Browser.Stealth.HumanScroll(page)

		// Wait for results container
		// Try multiple possible selectors for redundancy
		selectors := []string{
			".reusable-search__result-container",
			"li.reusable-search__result-container",
			".entity-result__item",
			"ul.reusable-search__entity-result-list > li",
			".search-results-container",
			"li.ps-search-result", // Older layout
			"div.search-results-container ul li",
		}

		foundSelector := ""
		for _, sel := range selectors {
			if has, _, _ := page.Has(sel); has {
				foundSelector = sel
				break
			}
		}

		if foundSelector == "" {
			slog.Warn("No results found or page structure changed (no selector matched).")
			break
		}
		slog.Info("Found results container selector", "selector", foundSelector)

		// Extract URLs from the found selector context
		elements := page.MustElements(foundSelector + " a.app-aware-link")
		slog.Info("Found primary elements", "count", len(elements))

		if len(elements) == 0 {
			// Fallback 1: specific list items
			elements = page.MustElements("ul.reusable-search__entity-result-list li a[href*='/in/']")
			slog.Info("Found fallback elements (list)", "count", len(elements))
		}

		if len(elements) == 0 {
			// Fallback 2: Global search for profile links on the page (aggressive)
			// We look for any link containing /in/ that is visible
			allLinks, _ := page.Elements("a[href*='/in/']")
			slog.Info("Found global /in/ links", "count", len(allLinks))

			// Filter for likely profile links (avoiding mini-profile duplications if possible, though 'Seen' map handles dupes)
			// We'll process all of them
			elements = allLinks
		}

		for _, el := range elements {
			href, err := el.Property("href")
			if err != nil {
				continue
			}
			val := href.String()
			slog.Info("Inspecting link", "url", val)

			// Clean URL (remove query params)
			if idx := strings.Index(val, "?"); idx != -1 {
				val = val[:idx]
			}

			// Validation: Must be a profile link, not a post or activity
			// Exclude /in/ACoAA... (mini profiles) sometimes, but usually /in/username is what we want.
			if strings.Contains(val, "/in/") && !s.Seen[val] {
				profiles = append(profiles, val)
				s.Seen[val] = true
				slog.Debug("Found profile", "url", val)

				if len(profiles) >= opts.Limit {
					break
				}
			}
		}

		if len(profiles) >= opts.Limit {
			break
		}

		// Next Page
		slog.Info("Looking for Next button")

		// Button usually has aria-label="Next" or class "artdeco-pagination__button--next"
		nextBtn, err := page.Element("button[aria-label='Next']")
		if err != nil {
			nextBtn, err = page.Element(".artdeco-pagination__button--next")
		}

		if err != nil || nextBtn == nil {
			slog.Info("No more pages available (Next button not found)")
			break
		}

		if disabled, _ := nextBtn.Attribute("disabled"); disabled != nil {
			slog.Info("Next button is disabled")
			break
		}

		// Scroll to button
		if err := s.Browser.Stealth.MoveToElement(page, nextBtn); err != nil {
			slog.Warn("Could not move to next button, trying direct scroll")
			nextBtn.ScrollIntoView()
		}

		nextBtn.MustClick()
		s.Browser.Stealth.SleepRandom(3*time.Second, 6*time.Second)
		page.MustWaitLoad()
	}

	slog.Info("Search complete", "total_found", len(profiles))
	return profiles, nil
}
