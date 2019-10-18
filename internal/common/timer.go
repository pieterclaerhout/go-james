package common

import (
	"time"

	"github.com/pieterclaerhout/go-log"
)

// Timer is what can be injected into a subcommand when you need a timer
type Timer struct {
	start time.Time
}

// StartTimer starts the timer
func (t *Timer) StartTimer() {
	t.start = time.Now()
}

// PrintElapsed prints the elapsed time to stdout
func (t *Timer) PrintElapsed(prefix string) {
	log.Info(prefix, time.Since(t.start))
}
