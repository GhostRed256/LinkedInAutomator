package stealth

import (
	"log/slog"
	"time"

	"github.com/go-rod/rod"
)

// IsBusinessHours checks if current time is within 9 AM to 6 PM
func (e *Engine) IsBusinessHours() bool {
	now := time.Now()
	hour := now.Hour()
	// Simple 9-18 check
	if hour >= 9 && hour < 18 {
		return true
	}
	slog.Info("Current time is outside business hours", "hour", hour)
	return false
}

// RandomHover moves the mouse to a random interactive element on the screen
func (e *Engine) RandomHover(page *rod.Page) {
	// Find visible links or buttons
	elements, err := page.Elements("a, button")
	if err != nil || len(elements) == 0 {
		return
	}

	// Pick a random one
	idx := e.rnd.Intn(len(elements))
	el := elements[idx]

	// Only if visible
	if visible, _ := el.Visible(); visible {
		slog.Info("Stealth: Randomly hovering over element")
		_ = e.MoveToElement(page, el)
		e.SleepRandom(500*time.Millisecond, 1500*time.Millisecond)
	}
}
