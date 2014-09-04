// Package metrics is a go library for sampling, counting and timing go code
// to be output in the l2met format.
package metrics

import (
	"fmt"
	"os"
)

// Namespace represents a metric prefix. (e.g. "memcached.")
type Namespace string

var (
	// Drain is the Drainer that will be used to drain metrics.
	Drain = Drainer(&LogDrain{})

	// Source is the root source that these metrics are coming from.
	Source = os.Getenv("DYNO")

	// DefaultNamespace is the default namespace to output metrics. By default,
	// no namespace.
	DefaultNamespace Namespace
)

// Count logs a count metric.
func (n Namespace) Count(name string, v interface{}) {
	n.drain("count", n.prefix(name), v, "")
}

// Sample logs a sample metric.
func (n Namespace) Sample(name string, v interface{}, units string) {
	n.drain("sample", n.prefix(name), v, units)
}

// Measure logs a measurement metric.
func (n Namespace) Measure(name string, v interface{}, units string) {
	n.drain("measure", n.prefix(name), v, units)
}

// Time starts a timer and returns it.
func (n Namespace) Time(name string) *Timer {
	return NewTimer(n.prefix(name))
}

// prefix prefixes the namespace onto the metric name.
func (n Namespace) prefix(name string) string {
	if n != "" {
		return fmt.Sprintf("%s.%s", n, name)
	}

	return name
}

// drain drains the metric to the Drainer.
func (n Namespace) drain(t, name string, v interface{}, units string) {
	drain(&metric{name: name, typ: t, value: v, units: units})
}

// Count logs a count metric in the root namespace.
func Count(name string, v interface{}) {
	DefaultNamespace.Count(name, v)
}

// Sample logs a sample metric in the root namespace.
func Sample(name string, v interface{}, units string) {
	DefaultNamespace.Sample(name, v, units)
}

// Measure logs a measurement metric in the root namespace.
func Measure(name string, v interface{}, units string) {
	DefaultNamespace.Measure(name, v, units)
}

// Time starts a timer and returns it.
func Time(name string) *Timer {
	return DefaultNamespace.Time(name)
}

// drain drains the metric using the configured Drainer.
func drain(m Metric) error {
	return Drain.Drain(m)
}
