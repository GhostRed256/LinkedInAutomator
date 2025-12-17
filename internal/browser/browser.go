package browser

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/stealth"

	"linkin-automator/internal/config"
	pkgStealth "linkin-automator/pkg/stealth"
)

type Browser struct {
	RodArgs *rod.Browser
	Page    *rod.Page // Active page
	Stealth *pkgStealth.Engine
	Config  *config.Config
}

func New(cfg *config.Config) (*Browser, error) {
	// Configure Launcher
	l := launcher.New().
		Headless(cfg.App.Headless).
		Devtools(false).
		Leakless(false) // Disable leakless to avoid AV false positives

		// User Agent Randomization (Basic pool for now)
	// Ideally we fetch this or have a large list.
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	}
	ua := userAgents[time.Now().Unix()%int64(len(userAgents))]
	l.Set("user-agent", ua)

	url, err := l.Launch()
	if err != nil {
		return nil, fmt.Errorf("failed to launch browser: %w", err)
	}

	b := rod.New().ControlURL(url).MustConnect()

	// Initialize Stealth Engine
	s := pkgStealth.New()

	return &Browser{
		RodArgs: b,
		Stealth: s,
		Config:  cfg,
	}, nil
}

func (b *Browser) Open(url string) error {
	var err error

	// Create a new incognito page properly?
	// For Rod, b.MustPage() creates a page in the default context.
	// To be safer, we should use Incognito if we want fresh starts,
	// but users might want session persistence (cookies).
	// For this assignment: "Persist session cookies for seamless reuse".
	// So we use the default context.

	b.Page, err = b.RodArgs.Page(proto.TargetCreateTarget{URL: url})
	if err != nil {
		// Fallback
		b.Page = b.RodArgs.MustPage(url)
	}

	// Apply Stealth (go-rod/stealth)
	// This injects scripts to mask navigator.webdriver etc.
	b.Page.MustEvalOnNewDocument(stealth.JS)

	// Set Viewport to something common
	// Randomize slightly
	w := 1920 + b.Stealth.RandomInt(-50, 50) // Need to add RandomInt to engine
	h := 1080 + b.Stealth.RandomInt(-50, 50)
	b.Page.MustSetViewport(w, h, 1, false)

	// Navigate if created generic
	// If we used MustPage(url) it's already there, but re-nav is fine.
	// Wait for load
	b.Page.MustWaitLoad()

	return nil
}

func (b *Browser) Close() {
	b.RodArgs.MustClose()
}
