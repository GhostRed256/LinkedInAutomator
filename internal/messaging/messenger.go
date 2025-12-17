package messaging

import (
	"fmt"
	"linkin-automator/internal/browser"
	"linkin-automator/internal/storage"
	"log/slog"
	"time"
)

type Messenger struct {
	Browser *browser.Browser
	Storage *storage.Storage
}

func New(b *browser.Browser, s *storage.Storage) *Messenger {
	return &Messenger{
		Browser: b,
		Storage: s,
	}
}

// CheckAcceptedConnections scans the connections page specifically
func (m *Messenger) CheckAcceptedConnections() error {
	slog.Info("Checking for new connections...")
	page := m.Browser.Page

	// Navigate to Connections page
	err := page.Navigate("https://www.linkedin.com/mynetwork/invite-connect/connections/")
	if err != nil {
		return err
	}
	page.MustWaitLoad()
	m.Browser.Stealth.SleepRandom(2*time.Second, 4*time.Second)

	// Scrape names/URLs
	// Selector for recent connections
	elements := page.MustElements(".mn-connection-card__link") // Standard class for connection links

	count := 0
	for _, el := range elements {
		href, err := el.Property("href")
		if err != nil {
			continue
		}
		val := href.String()

		// Clean URL
		if idx := 0; idx != -1 { // dummy check
			// Basic cleaning logic if needed
		}

		// Check if we tracked this as "sent"
		if m.Storage.HasPendingRequest(val) {
			slog.Info("Detected accepted connection!", "url", val)
			m.Storage.MarkAsConnected(val)
			count++
		}
	}
	slog.Info("Finished checking connections", "new_accepted", count)
	return nil
}

// SendFollowUpMessages sends a message to recently connected users
func (m *Messenger) SendFollowUpMessages(template string) error {
	targets := m.Storage.GetPendingFollowups()
	slog.Info("Sending follow-up messages", "count", len(targets))

	for _, url := range targets {
		slog.Info("Sending follow-up", "to", url)

		page := m.Browser.Page
		page.MustNavigate(url)
		page.MustWaitLoad()
		m.Browser.Stealth.SleepRandom(2*time.Second, 4*time.Second)

		// Click Message Button
		// Usually "Message" or icon
		msgBtn, err := page.ElementR("button", "Message")
		if err != nil {
			// Try aria-label
			msgBtn, err = page.Element("button[aria-label^='Message']")
		}

		if err == nil && msgBtn != nil {
			msgBtn.MustClick()
			m.Browser.Stealth.SleepRandom(1*time.Second, 2*time.Second)

			// Focus editor
			editor, err := page.Element(".msg-form__contenteditable")
			if err != nil {
				// try fallback selector
				editor, err = page.Element("div[role='textbox']")
			}

			if err == nil {
				// Template substitution (Basic)
				// Real impl would parse First Name from page
				msg := fmt.Sprintf(template, "there")

				m.Browser.Stealth.HumanType(editor, msg)
				m.Browser.Stealth.SleepRandom(1*time.Second, 2*time.Second)

				// Send
				sendBtn, err := page.Element(".msg-form__send-button")
				if err == nil {
					sendBtn.MustClick()
					slog.Info("Message sent successfully")
					m.Storage.MarkFollowupSent(url)
				}
			}
		}

		m.Browser.Stealth.SleepRandom(3*time.Second, 8*time.Second)
	}
	return nil
}
