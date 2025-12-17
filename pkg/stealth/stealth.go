package stealth

import (
	"math/rand"
	"time"
)

// Engine holds the stealth configuration and state
type Engine struct {
	rnd   *rand.Rand
	lastX float64
	lastY float64
}

func New() *Engine {
	return &Engine{
		rnd:   rand.New(rand.NewSource(time.Now().UnixNano())),
		lastX: 0,
		lastY: 0,
	}
}

// RandomInt returns a random int between min and max
func (e *Engine) RandomInt(min, max int) int {
	return min + e.rnd.Intn(max-min+1)
}

// RandomDuration returns a random duration between min and max
func (e *Engine) RandomDuration(min, max time.Duration) time.Duration {
	delta := max - min
	if delta <= 0 {
		return min
	}
	return min + time.Duration(e.rnd.Int63n(int64(delta)))
}

// SleepRandom sleeps for a random duration
func (e *Engine) SleepRandom(min, max time.Duration) {
	time.Sleep(e.RandomDuration(min, max))
}
