// Package metrics is a go library for sampling, counting and timing go code
// to be output in the l2met format.
package metrics

import (
	"fmt"
	"os"
)

type Namespace string

var (
	// Drain is the Drainer that will be used to drain metrics.
	Drain Drainer = &LogDrain{}

	// The root source that these metrics are coming from.
	Source = os.Getenv("DYNO")

	// The root namespace.
	root Namespace = ""
)

// Count logs a count metric.
func (n Namespace) Count(metric string, v interface{}) {
	n.drain("count", n.prefix(metric), v, "")
}

// Sample logs a sample metric.
func (n Namespace) Sample(metric string, v interface{}, units string) {
	n.drain("sample", n.prefix(metric), v, units)
}

// Measure logs a measurement metric.
func (n Namespace) Measure(metric string, v interface{}, units string) {
	n.drain("measure", n.prefix(metric), v, units)
}

// Time starts a timer and returns it.
func (n Namespace) Time(metric string) *Timer {
	return NewTimer(n.prefix(metric))
}

// prefix prefixes the namespace onto the metric name.
func (n Namespace) prefix(metric string) string {
	if n != "" {
		return fmt.Sprintf("%s.%s", n, metric)
	} else {
		return metric
	}
}

// drain drains the metric to the Drainer.
func (n Namespace) drain(t, metric string, v interface{}, units string) {
	drain(&coreMetric{name: metric, typ: t, value: v, units: units})
}

// Count logs a count metric in the root namespace.
func Count(metric string, v interface{}) {
	root.Count(metric, v)
}

// Sample logs a sample metric in the root namespace.
func Sample(metric string, v interface{}, units string) {
	root.Sample(metric, v, units)
}

// Measure logs a measurement metric in the root namespace.
func Measure(metric string, v interface{}, units string) {
	root.Measure(metric, v, units)
}

// Time starts a timer and returns it.
func Time(metric string) *Timer {
	return root.Time(metric)
}

// drain drains the metric using the configured Drainer.
func drain(m Metric) error {
	return Drain.Drain(m)
}
