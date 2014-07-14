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
	Sample("runtime.MemStats.Alloc", r.MemStats.Alloc, "")
	Sample("runtime.MemStats.Frees", r.MemStats.Frees, "")
	Sample("runtime.MemStats.HeapAlloc", r.MemStats.HeapAlloc, "")
	Sample("runtime.MemStats.HeapIdle", r.MemStats.HeapIdle, "")
	Sample("runtime.MemStats.HeapObjects", r.MemStats.HeapObjects, "")
	Sample("runtime.MemStats.HeapReleased", r.MemStats.HeapReleased, "")
	Sample("runtime.MemStats.HeapSys", r.MemStats.HeapSys, "")
	Sample("runtime.MemStats.LastGC", r.MemStats.LastGC, "")
	Sample("runtime.MemStats.Lookups", r.MemStats.Lookups, "")
	Sample("runtime.MemStats.Mallocs", r.MemStats.Mallocs, "")
	Sample("runtime.MemStats.MCacheInuse", r.MemStats.MCacheInuse, "")
	Sample("runtime.MemStats.MCacheSys", r.MemStats.MCacheSys, "")
	Sample("runtime.MemStats.MSpanInuse", r.MemStats.MSpanInuse, "")
	Sample("runtime.MemStats.MSpanSys", r.MemStats.MSpanSys, "")
	Sample("runtime.MemStats.NextGC", r.MemStats.NextGC, "")
	Sample("runtime.MemStats.NumGC", r.MemStats.NumGC, "")
	Sample("runtime.MemStats.StackInuse", r.MemStats.StackInuse, "")
}

// Runtime enters into a loop, sampling and outputing the runtime stats periodically.
func Runtime() {
	c := time.Tick(5 * time.Second)
	for _ = range c {
		r := NewRuntimeSample()
		r.Drain()
	}
}
