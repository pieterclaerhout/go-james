package common

import (
	"fmt"
	"math"
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
	log.Info("\n"+prefix, t.formatDuration(time.Since(t.start)))
}

// formatDuration formats a duration as HH:MM:SS
func (t *Timer) formatDuration(duration time.Duration) string {

	durationAsFloat := duration.Seconds()

	hours := math.Floor(durationAsFloat / 3600.0)
	minutes := math.Floor((durationAsFloat - (hours * 3600.0)) / 60)
	seconds := int64(durationAsFloat) % 60

	return fmt.Sprintf("%02.0f:%02.0f:%02d", hours, minutes, seconds)

}
