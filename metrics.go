// Package metrics is a go library for sampling, counting and timing go code
// to be output in the l2met format.
package metrics

import (
	"fmt"
	"log"
	"os"
)

type Namespace string

// Logger is the logger to use when printing metrics. By default, metrics are printed
// to Stdout.
var Logger = log.New(os.Stdout, "", 0)

// The root namespace.
var root Namespace = ""

// Count logs a count metric.
func (n Namespace) Count(metric string, v interface{}) {
	n.print(MetricCount, metric, v, "")
}

// Sample logs a sample metric.
func (n Namespace) Sample(metric string, v interface{}, units string) {
	n.print(MetricSample, metric, v, units)
}

// Measure logs a measurement metric.
func (n Namespace) Measure(metric string, v interface{}, units string) {
	n.print(MetricMeasure, metric, v, units)
}

// print prints a metric type to the logger.
func (n Namespace) print(t MetricType, metric string, v interface{}, units string) {
	if n != "" {
		metric = fmt.Sprintf("%s.%s", n, metric)
	}
	m := &Metric{Name: metric, Type: t, Value: v, Units: units}
	m.Print()
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
