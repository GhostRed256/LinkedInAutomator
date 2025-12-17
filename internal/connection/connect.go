package connection

import (
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

	c.Browser.Stealth.SleepRandom(1*time.Second, 2*time.Second)

	slog.Info("Simulating profile reading with scroll")
	// Just 2 simple scrolls - don't overdo it
	c.Browser.Stealth.HumanScroll(page)
	c.Browser.Stealth.SleepRandom(500*time.Millisecond, 1*time.Second)
	c.Browser.Stealth.HumanScroll(page)
	c.Browser.Stealth.SleepRandom(500*time.Millisecond, 1*time.Second)

	// Find Connect button
	// It is often in the "pvs-profile-actions" area
	// Primary button might be "Connect", "Follow", "Message", or "Pending"

	slog.Info("Looking for Connect action")
	var connectBtn *rod.Element

	// Try to find "Connect" button first
	connectCandidates, _ := page.Elements("button")
	for _, btn := range connectCandidates {
		txt, err := btn.Text()
		if err != nil {
			continue
		}
		if txt == "Connect" || txt == "connect" {
			if visible, _ := btn.Visible(); visible {
				connectBtn = btn
				slog.Info("Found Connect button by text")
				break
			}
		}
		// Check aria-label
		if lbl, _ := btn.Attribute("aria-label"); lbl != nil && strings.Contains(*lbl, "Connect") {
			if visible, _ := btn.Visible(); visible {
				connectBtn = btn
				slog.Info("Found Connect button by aria-label")
				break
			}
		}
	}

	if connectBtn == nil {
		slog.Warn("Connect button not easily accessible, skipping this profile")
		return nil
	}

	// Click Connect
	slog.Info("Clicking Connect button")
	if err := c.Browser.Stealth.MoveToElement(page, connectBtn); err != nil {
		slog.Warn("Failed to move to connect button", "error", err)
	}

	connectBtn.MustClick()

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
		} else {
			slog.Warn("Send button disabled?")
		}
	} else {
		slog.Warn("Could not find Send button in modal")
	}

	// Record in storage
	c.Storage.AddRequest(profileURL, message != "")
	slog.Info("Connection request completed", "profile", profileURL)

	// Navigate back to search results for next profile
	c.Browser.Stealth.SleepRandom(500*time.Millisecond, 1*time.Second)
	slog.Info("Navigating back to search results")
	page.MustNavigateBack()
	page.MustWaitLoad()
	c.Browser.Stealth.SleepRandom(500*time.Millisecond, 1*time.Second)

	return nil
}
