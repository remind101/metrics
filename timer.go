package metrics

import "time"

// Timer is an implementation of the Metric interface for timing things.
type Timer struct {
	*metric
	start time.Time
	end   time.Time
}

// NewTimer returns a new Timer metric.
func NewTimer(name string) *Timer {
	return &Timer{
		metric: &metric{
			name:  name,
			typ:   "measure",
			units: "ms",
		},
		start: time.Now(),
	}
}

// Value returns the difference between start and end in milliseconds.
func (t *Timer) Value() interface{} {
	if t.metric.value == nil {
		t.metric.value = t.Milliseconds()
	}
	return t.metric.value
}

// Milliseconds returns the number of milliseconds elapsed.
func (t *Timer) Milliseconds() int64 {
	return t.end.Sub(t.start).Nanoseconds() / int64(time.Millisecond)
}

// Stop stops the timer.
func (t *Timer) Stop() {
	t.end = time.Now()
}

// Done stops the timer and drains it.
func (t *Timer) Done() {
	t.Stop()
	drain(t)
}
