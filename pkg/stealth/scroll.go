package stealth

import (
	"time"

	"github.com/go-rod/rod"
)

// HumanScroll performs a human-like scroll on the page
func (e *Engine) HumanScroll(page *rod.Page) {
	// Scroll a random amount
	amount := e.rnd.Intn(500) + 200 // 200-700 pixels
	if e.rnd.Intn(2) == 0 {
		// Occasionally scroll up
		amount = -amount / 3
	}

	// Break into steps
	steps := 10
	stepSize := float64(amount) / float64(steps)

	for i := 0; i < steps; i++ {
		// e.Mouse.Scroll? Rod doesn't have direct Mouse.Scroll easily exposed on Input?
		// Page.Mouse.Scroll(x, y, steps)
		page.Mouse.Scroll(0, stepSize, 1)

		// Random micro delay
		e.SleepRandom(10*time.Millisecond, 50*time.Millisecond)
	}
}
