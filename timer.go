package metrics

import "time"

// DefaultNowFunc is the default function returning a time.Time for "now".
var DefaultNowFunc = func() time.Time {
	return time.Now()
}

// Timer is an implementation of the Metric interface for timing things.
type Timer struct {
	// NowFunc is a function to return a time.Time representing "now". Zero value
	// is DefaultNowFunc.
	NowFunc func() time.Time

	*metric
	start time.Time
	end   time.Time
}

// NewTimer returns a new Timer metric.
func NewTimer(name string) *Timer {
	t := &Timer{
		metric: &metric{
			name:  name,
			typ:   "measure",
			units: "ms",
		},
	}
	t.Start()

	return t
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

// Start the timer.
func (t *Timer) Start() {
	t.start = t.now()
}

// Stop stops the timer.
func (t *Timer) Stop() {
	t.end = t.now()
}

// Done stops the timer and drains it.
func (t *Timer) Done() {
	t.Stop()
	drain(t)
}

func (t *Timer) now() time.Time {
	if t.NowFunc == nil {
		t.NowFunc = DefaultNowFunc
	}

	return t.NowFunc()
}
