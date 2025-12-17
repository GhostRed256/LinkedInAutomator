package stealth

import (
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// HumanType types text into an element with realistic delays and typos
func (e *Engine) HumanType(el *rod.Element, text string) error {
	// Click to focus?
	// Often good to move mouse there first
	if err := e.MoveToElement(el.Page(), el); err != nil {
		return err // Element might be gone or hidden
	}

	// Click
	el.Click(proto.InputMouseButtonLeft, 1)
	e.SleepRandom(100*time.Millisecond, 300*time.Millisecond)

	for _, char := range text {
		// Occasional error (5% chance)
		if e.rnd.Float64() < 0.05 {
			// Type a wrong character
			wrongChar := string(char + 1) // simple wrong char
			el.Input(wrongChar)
			e.SleepRandom(50*time.Millisecond, 150*time.Millisecond)

			// Backspace - using explicit backspace character
			el.Input("\b")
			e.SleepRandom(100*time.Millisecond, 200*time.Millisecond)
		}

		el.Input(string(char))

		// Variable delay between keystrokes
		// Average 100ms, but vary
		baseDelay := 80
		variance := 50
		delay := baseDelay + e.rnd.Intn(variance) - (variance / 2)
		if delay < 20 {
			delay = 20
		}

		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
	return nil
}
