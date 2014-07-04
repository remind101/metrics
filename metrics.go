// Package metrics is a go library for sampling, counting and timing go code
// to be output in the l2met format.
package metrics

import "fmt"

type Namespace string

var (
	// Logger is the logger to use when printing metrics. By default, metrics are printed
	// to Stdout.
	Drain Drainer = &LogDrain{}

	// The root namespace.
	root Namespace = ""
)

// Count logs a count metric.
func (n Namespace) Count(metric string, v interface{}) {
	n.print("count", metric, v, "")
}

// Sample logs a sample metric.
func (n Namespace) Sample(metric string, v interface{}, units string) {
	n.print("sample", metric, v, units)
}

// Measure logs a measurement metric.
func (n Namespace) Measure(metric string, v interface{}, units string) {
	n.print("measure", metric, v, units)
}

// print prints a metric type to the logger.
func (n Namespace) print(t, metric string, v interface{}, units string) {
	if n != "" {
		metric = fmt.Sprintf("%s.%s", n, metric)
	}
	m := &coreMetric{name: metric, typ: t, value: v, units: units}
	Drain.Drain(m)
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
	return NewTimer(metric)
}
