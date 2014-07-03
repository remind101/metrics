package metrics

import (
	"errors"
	"time"
)

var (
	ErrTimerStopped = errors.New("this timer has already been stopped")
)

type Timer struct {
	// A Timer is a Metric.
	*Metric

	// The time that this Timer started.
	Start time.Time

	// The time that this Timer ended.
	End time.Time

	stopped bool
}

// NewTimer returns a new Timer.
func NewTimer(metric string) *Timer {
	return &Timer{
		Metric: &Metric{Type: Measurement, Name: metric, Units: "ms"},
		Start:  time.Now(),
	}
}

// Duration returns the duration of this timer.
func (t *Timer) Duration() time.Duration {
	return t.End.Sub(t.Start)
}

// Stop sets End on the Timer and calculates the value for the Metric.
func (t *Timer) Stop() {
	t.Value = t.Duration() * time.Millisecond
	t.stopped = true
}

// Done stops the timer and prints it.
func (t *Timer) Done() error {
	if t.stopped {
		return ErrTimerStopped
	}

	t.Stop()
	t.Print()

	return nil
}
