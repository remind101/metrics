package metrics

import (
	"runtime"
	"time"
)

// RuntimeSample represents a sampling of the runtime stats.
type RuntimeSample struct {
	*runtime.MemStats
	NumGoroutine int
}

// NewRuntimeSample samples the current runtime and returns a RuntimeSample.
func NewRuntimeSample() *RuntimeSample {
	r := &RuntimeSample{MemStats: &runtime.MemStats{}}
	runtime.ReadMemStats(r.MemStats)
	r.NumGoroutine = runtime.NumGoroutine()
	return r
}

// Drain drains all of the metrics.
func (r *RuntimeSample) Drain() {
	Sample("goroutine", r.NumGoroutine, "")
	Sample("memory.allocated", r.MemStats.Alloc, "")
	Sample("memory.mallocs", r.MemStats.Mallocs, "")
	Sample("memory.frees", r.MemStats.Frees, "")
	Sample("memory.heap", r.MemStats.HeapAlloc, "")
	Sample("memory.stack", r.MemStats.StackInuse, "")
}

// Runtime enters into a loop, sampling and outputing the runtime stats periodically.
func Runtime() {
	c := time.Tick(5 * time.Second)
	for _ = range c {
		r := NewRuntimeSample()
		r.Drain()
	}
}
