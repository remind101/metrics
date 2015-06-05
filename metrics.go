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
	// DefaultDrain is the Drainer that will be used to drain metrics.
	DefaultDrain = Drainer(&LogDrain{})

	// Source is the root source that these metrics are coming from.
	Source string

	// DefaultNamespace is the default namespace to output metrics. By default,
	// no namespace.
	DefaultNamespace Namespace
)

func init() {
	hostname, _ := os.Hostname()
	prefix := os.Getenv("SOURCE")
	if prefix == "" {
		prefix = os.Getenv("DYNO")
	}
	Source = fmt.Sprintf("%s.%s", prefix, hostname)
}

// NewMetric returns a new Metric.
func (n Namespace) NewMetric(t, name string, v interface{}, units string) Metric {
	return &metric{name: n.prefix(name), typ: t, value: v, units: units}
}

// CountMetric returns a new Metric for a count.
func (n Namespace) CountMetric(name string, v interface{}) Metric {
	return n.NewMetric("count", name, v, "")
}

// Count creates a count metric and drains it.
func (n Namespace) Count(name string, v interface{}) {
	drain(n.CountMetric(name, v))
}

// SampleMetric returns a new Metric for a sample.
func (n Namespace) SampleMetric(name string, v interface{}, units string) Metric {
	return n.NewMetric("sample", name, v, units)
}

// Sample creates a sample metric and drains it.
func (n Namespace) Sample(name string, v interface{}, units string) {
	drain(n.SampleMetric(name, v, units))
}

// MeasureMetric returns a new Metric for a measure.
func (n Namespace) MeasureMetric(name string, v interface{}, units string) Metric {
	return n.NewMetric("measure", name, v, units)
}

// Measure creates a measure metric and drains it.
func (n Namespace) Measure(name string, v interface{}, units string) {
	drain(n.MeasureMetric(name, v, units))
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
	return DefaultDrain.Drain(m)
}
