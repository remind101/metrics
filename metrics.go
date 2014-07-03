// Package metrics is a go library for sampling, counting and timing go code
// to be output in the l2met format.
package metrics

import (
	"log"
	"os"
)

// Logger is the logger to use when printing metrics. By default, metrics are printed
// to Stdout.
var Logger = log.New(os.Stdout, "", 0)

func Count(metric string, v interface{}) {
	printMetric(MetricCount, metric, v, "")
}

func Sample(metric string, v interface{}, units string) {
	printMetric(MetricSample, metric, v, units)
}

func Measure(metric string, v interface{}, units string) {
	printMetric(MetricMeasure, metric, v, units)
}

func printMetric(t MetricType, metric string, v interface{}, units string) {
	m := &Metric{Name: metric, Type: t, Value: v, Units: units}
	m.Print()
}
