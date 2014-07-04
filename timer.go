package metrics

import "time"

// Timer is an implementation of the Metric interface for timing things.
type Timer struct {
	name  string
	start time.Time
	end   time.Time
	value interface{}
}

// NewTimer returns a new Timer metric.
func NewTimer(metric string) *Timer {
	return &Timer{name: metric, start: time.Now()}
}

// Methods to implement the Metric interface
func (t *Timer) Name() string  { return t.name }
func (t *Timer) Type() string  { return "measure" }
func (t *Timer) Units() string { return "ms" }
func (t *Timer) Value() interface{} {
	if t.value == nil {
		t.value = t.Milliseconds()
	}
	return t.value
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
