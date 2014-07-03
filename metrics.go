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

func (n Namespace) Count(metric string, v interface{}) {
	n.print(MetricCount, metric, v, "")
}

func (n Namespace) Sample(metric string, v interface{}, units string) {
	n.print(MetricSample, metric, v, units)
}

func (n Namespace) Measure(metric string, v interface{}, units string) {
	n.print(MetricMeasure, metric, v, units)
}

func (n Namespace) print(t MetricType, metric string, v interface{}, units string) {
	if n != "" {
		metric = fmt.Sprintf("%s.%s", n, metric)
	}
	m := &Metric{Name: metric, Type: t, Value: v, Units: units}
	m.Print()
}

func Count(metric string, v interface{}) {
	root.Count(metric, v)
}

func Sample(metric string, v interface{}, units string) {
	root.Sample(metric, v, units)
}

func Measure(metric string, v interface{}, units string) {
	root.Measure(metric, v, units)
}
