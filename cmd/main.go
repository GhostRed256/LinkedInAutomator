package main

import (
	"os"
	"time"

	"linkin-automator/internal/auth"
	"linkin-automator/internal/browser"
	"linkin-automator/internal/config"
	"linkin-automator/internal/connection"
	"linkin-automator/internal/messaging"
	"linkin-automator/internal/search"
	"linkin-automator/internal/storage" // Added messaging
	"linkin-automator/pkg/logger"
	"log/slog"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	// 2. Initialize Logger
	logger.Init(cfg.App.LogLevel)

	// 3. Initialize Browser
	b, err := browser.New(cfg)
	if err != nil {
		slog.Error("Failed to create browser", "error", err)
		os.Exit(1)
	}
	defer b.Close()

	// Open initial page
	if err := b.Open("https://www.google.com"); err != nil {
		slog.Error("Failed to open browser", "error", err)
	}

	// 4. Authenticate
	authenticator := auth.New(b, cfg)
	if err := authenticator.Login(); err != nil {
		slog.Error("Authentication failed", "error", err)
		os.Exit(1)
	}

	// 5. Initialize Storage
	store := storage.New()

	// 6. Search
	searcher := search.New(b)
	connector := connection.New(b, store)

	keywords := cfg.Search.Keywords
	slog.Info("Starting search and connect loop", "keywords", keywords)

	profiles, err := searcher.Run(search.SearchOptions{
		Keywords: keywords,
		Limit:    cfg.Search.Limit,
	})
	if err != nil {
		slog.Error("Search failed", "error", err)
	}

	slog.Info("Found profiles", "count", len(profiles))

	// 7. Connect Loop
	for _, p := range profiles {
		// Business hours check (warning only for POC)
		if !b.Stealth.IsBusinessHours() {
			slog.Warn("Operating outside business hours (Stealth risk)")
		}

		// Daily limit check
		if count := store.CountRequestsToday(); count >= 20 {
			slog.Warn("Daily limit reached, stopping")
			break
		}

		// Random Stealth Hover before action
		if b.Stealth.RandomInt(0, 10) > 7 {
			b.Stealth.RandomHover(b.Page)
		}

		if err := connector.SendConnectionRequest(p, "Hi, I'd like to connect!"); err != nil {
			slog.Error("Failed to connect", "profile", p, "error", err)
		}

		// Random delay between requests
		b.Stealth.SleepRandom(30*time.Second, 60*time.Second)
	}

	// 8. Messaging
	msgr := messaging.New(b, store)
	if err := msgr.CheckAcceptedConnections(); err != nil {
		slog.Error("Failed to check connections", "error", err)
	}

	// Example template
	if err := msgr.SendFollowUpMessages("Hi %s, thanks for connecting! Looking forward to seeing your updates."); err != nil {
		slog.Error("Failed to send follow-ups", "error", err)
	}

	slog.Info("Workflow complete")
}
