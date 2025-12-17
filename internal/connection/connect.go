package connection

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-rod/rod"

	"linkin-automator/internal/browser"
	"linkin-automator/internal/storage"
	"log/slog"
)

type Connector struct {
	Browser *browser.Browser
	Storage *storage.Storage
}

func New(b *browser.Browser, s *storage.Storage) *Connector {
	return &Connector{
		Browser: b,
		Storage: s,
	}
}

// SendConnectionRequest visits profile and tries to connect
func (c *Connector) SendConnectionRequest(profileURL string, message string) error {
	if c.Storage.HasPendingRequest(profileURL) {
		slog.Info("Skipping already requested profile", "url", profileURL)
		return nil
	}

	page := c.Browser.Page
	slog.Info("Navigating to profile", "url", profileURL)
	page.MustNavigate(profileURL)
	page.MustWaitLoad()

	// Random scroll to simulate reading
	c.Browser.Stealth.SleepRandom(2*time.Second, 4*time.Second)
	c.Browser.Stealth.HumanScroll(page)
	c.Browser.Stealth.SleepRandom(1*time.Second, 2*time.Second) // Pause "reading"

	// Find Connect button
	// It is often in the "pvs-profile-actions" area
	// Primary button might be "Connect", "Follow", "Message", or "Pending"

	slog.Info("Looking for Connect action")

	// Strategy 1: Look for direct "Connect" button
	// Often: button with text "Connect" inside specific container
	// We iterate buttons to find one with text "Connect"
	// Using XPath or iterating elements

	buttons := page.MustElements("button")
	var connectBtn *rod.Element

	for _, btn := range buttons {
		// Check text
		txt, err := btn.Text()
		if err == nil && strings.TrimSpace(txt) == "Connect" {
			// Ensure it's visible?
			if visible, _ := btn.Visible(); visible {
				connectBtn = btn
				break
			}
		}
		// Check aria-label
		if lbl, _ := btn.Attribute("aria-label"); lbl != nil && strings.Contains(*lbl, "Connect") {
			if visible, _ := btn.Visible(); visible {
				connectBtn = btn
				break
			}
		}
	}

	if connectBtn == nil {
		slog.Info("Connect button not found directly, checking 'More' menu")
		// Find "More" button (usually aria-label="More actions")
		moreBtn, err := page.Element("button[aria-label='More actions']")
		if err != nil || moreBtn == nil {
			// Try searching text "More"
			// Keep it simple for POC
			return fmt.Errorf("could not find Connect or More button")
		}

		moreBtn.MustClick()
		c.Browser.Stealth.SleepRandom(500*time.Millisecond, 1000*time.Millisecond)

		// Now look for "Connect" in dropdown items
		// Dropdown items are usually in a div/ul
		// We look for any element with text "Connect" that is visible
		dropdownItems := page.MustElements("div[role='button'], li") // vague selector
		for _, item := range dropdownItems {
			txt, _ := item.Text()
			if strings.Contains(txt, "Connect") || strings.Contains(txt, "connect") {
				// Excluding "Connections" link
				if strings.Contains(txt, "Connect") && !strings.Contains(txt, "Connections") {
					connectBtn = item
					break
				}
			}
		}

		if connectBtn == nil {
			return fmt.Errorf("connect option not found in More menu (might rely on Follow?)")
		}
	}

	// Click Connect
	slog.Info("Clicking Connect")
	if err := c.Browser.Stealth.MoveToElement(page, connectBtn); err != nil {
		connectBtn.MustClick() // Fallback
	} else {
		connectBtn.MustClick()
	}

	c.Browser.Stealth.SleepRandom(1*time.Second, 2*time.Second)

	// Handle Modal
	// Modal might match "Add a note"
	// Buttons: "Add a note", "Send"

	slog.Info("Handling connection modal")

	if message != "" {
		addNoteBtn, err := page.ElementR("button", "Add a note")
		if err == nil && addNoteBtn != nil {
			addNoteBtn.MustClick()
			c.Browser.Stealth.SleepRandom(500*time.Millisecond, 1000*time.Millisecond)

			textArea, err := page.Element("textarea[name='message']")
			if err == nil {
				c.Browser.Stealth.HumanType(textArea, message)
				c.Browser.Stealth.SleepRandom(1*time.Second, 2*time.Second)
			}
		}
	}

	// Send
	sendBtn, err := page.ElementR("button", "Send")
	if err != nil {
		// Try aria-label="Send now"
		sendBtn, err = page.Element("button[aria-label='Send now']")
	}

	if err == nil && sendBtn != nil {
		// Just verify enabled
		if disabled, _ := sendBtn.Attribute("disabled"); disabled == nil {
			sendBtn.MustClick()
			slog.Info("Connection request sent")
			c.Storage.AddRequest(profileURL, message != "")
		} else {
			slog.Warn("Send button disabled?")
		}
	} else {
		slog.Warn("Could not find Send button in modal")
	}

	return nil
}
