package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"

	"linkin-automator/internal/browser"
	"linkin-automator/internal/config"
	"log/slog"
)

const (
	cookieFile = "cookies.json"
	loginURL   = "https://www.linkedin.com/login"
)

type Authenticator struct {
	Browser *browser.Browser
	Config  *config.Config
}

func New(b *browser.Browser, cfg *config.Config) *Authenticator {
	return &Authenticator{
		Browser: b,
		Config:  cfg,
	}
}

// Login performs the login flow or restores session
func (a *Authenticator) Login() error {
	slog.Info("Starting authentication flow")

	// 1. Try to restore session
	if err := a.loadCookies(); err == nil {
		slog.Info("Restored session from cookies")
		// Verify session validity by visiting feed
		a.Browser.Page.MustNavigate("https://www.linkedin.com/feed/")
		a.Browser.Page.MustWaitLoad()

		url := a.Browser.Page.MustInfo().URL
		if strings.Contains(url, "feed") && !strings.Contains(url, "login") && !strings.Contains(url, "checkpoint") {
			slog.Info("Session is valid", "url", url)
			return nil
		}
		slog.Warn("Session invalid or expired (redirected to login), initiating fresh login", "url", url)
	}

	// 2. Fresh Login
	return a.performLogin()
}

func (a *Authenticator) performLogin() error {
	page := a.Browser.Page

	slog.Info("Navigating to login page", "url", loginURL)
	page.MustNavigate(loginURL)
	page.MustWaitLoad()

	// Wait for username field
	slog.Info("Waiting for login form")

	var userInput, passInput, submitBtn *rod.Element

	// Try Selector Set 1 (Login Page)
	if has, _, _ := page.Has("#username"); has {
		userInput = page.MustElement("#username")
		passInput = page.MustElement("#password")
		submitBtn = page.MustElement("button[type=submit]")
	} else if has, _, _ := page.Has("#session_key"); has {
		// Selector Set 2 (Home Page / Alternative)
		userInput = page.MustElement("#session_key")
		passInput = page.MustElement("#session_password")
		// Button might be different
		if hasBtn, _, _ := page.Has("button[data-id='sign-in-form__submit-btn']"); hasBtn {
			submitBtn = page.MustElement("button[data-id='sign-in-form__submit-btn']")
		} else {
			// fallback generic submit
			submitBtn = page.MustElement("button[type=submit]")
		}
	} else if has, _, _ := page.Has("#password"); has {
		// Selector Set 3: "Welcome Back" (Password only)
		slog.Info("Detected 'Welcome Back' screen (Password only)")
		passInput = page.MustElement("#password")
		submitBtn = page.MustElement("button[type=submit]")
		// userInput remains nil
	} else {
		return fmt.Errorf("could not find login inputs (changed layout?)")
	}

	// Direct typing for reliability during login
	slog.Info("Entering credentials for user", "username", a.Config.LinkedIn.Username)

	// Only type username if input element was found
	if userInput != nil {
		// Ensure field is visible and focused
		// Check visibility to avoid hanging
		if visible, _ := userInput.Visible(); visible {
			userInput.MustScrollIntoView()
			userInput.MustSelectAllText().MustInput("")
			if err := userInput.Input(a.Config.LinkedIn.Username); err != nil {
				return fmt.Errorf("typing username: %w", err)
			}
		} else {
			slog.Info("Username field hidden, skipping username entry")
		}

		// Random delay between fields
		a.Browser.Stealth.SleepRandom(500*time.Millisecond, 1500*time.Millisecond)
	} else {
		slog.Info("No username input found, assuming email is already pre-filled/saved.")
	}

	if passInput != nil {
		passInput.MustWaitVisible()
		passInput.MustSelectAllText().MustInput("")
		if err := passInput.Input(a.Config.LinkedIn.Password); err != nil {
			return fmt.Errorf("typing password: %w", err)
		}
	}

	// Click Login
	slog.Info("Clicking login button")
	if err := a.Browser.Stealth.MoveToElement(page, submitBtn); err != nil {
		return fmt.Errorf("moving to submit: %w", err)
	}
	submitBtn.MustClick()

	// Wait for navigation
	// Check for success or 2FA
	slog.Info("Waiting for post-login navigation")
	page.MustWaitLoad()

	// Give it time to settle or redirect
	time.Sleep(5 * time.Second) // basic wait

	currentURL := page.MustInfo().URL
	if strings.Contains(currentURL, "challenge") || strings.Contains(currentURL, "checkpoint") {
		slog.Warn("2FA or Security Checkpoint detected!", "url", currentURL)
		slog.Warn("Please manually solve the challenge in the browser window/terminal and press Enter here to continue...")
		// In a real headless tool, we might pause or notify.
		// Since we are running potentially headed, we can wait.
		fmt.Println("Press Enter after solving CAPTCHA/2FA...")
		fmt.Scanln()

		// Refresh state
		currentURL = page.MustInfo().URL
	}

	// Strict Success Check: Must be on feed or have left the login page
	if strings.Contains(currentURL, "/login") || strings.Contains(currentURL, "uas/login") {
		return fmt.Errorf("login failed (still on login page). check credentials. url: %s", currentURL)
	}

	if !strings.Contains(currentURL, "feed") && !strings.Contains(currentURL, "linkedin.com") {
		// This might be too strict if they redirect elsewhere, but /login check above mostly covers it.
		// Let's rely on NOT being on /login
	}

	// Save cookies
	if err := a.saveCookies(); err != nil {
		slog.Error("Failed to save cookies", "error", err)
	}

	slog.Info("Login successful")
	return nil
}

func (a *Authenticator) saveCookies() error {
	cookies, err := a.Browser.Page.Browser().GetCookies()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cookies)
	if err != nil {
		return err
	}

	return os.WriteFile(cookieFile, data, 0644)
}

func (a *Authenticator) loadCookies() error {
	data, err := os.ReadFile(cookieFile)
	if err != nil {
		return err
	}

	var cookies []*proto.NetworkCookie
	if err := json.Unmarshal(data, &cookies); err != nil {
		return err
	}

	// Convert to NetworkCookieParam
	params := make([]*proto.NetworkCookieParam, len(cookies))
	for i, c := range cookies {
		params[i] = &proto.NetworkCookieParam{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Secure:   c.Secure,
			HTTPOnly: c.HTTPOnly,
			SameSite: c.SameSite,
			Expires:  c.Expires,
		}
	}

	return a.Browser.Page.Browser().SetCookies(params)
}
