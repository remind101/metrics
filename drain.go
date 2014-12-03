package metrics

import (
	"log"
	"os"
)

// Ensure implementations implement the Drainer interface.
var (
	_ Drainer = &LogDrain{}
	_ Drainer = &NullDrain{}
)

// Drainer is an interface that can drain a metric to it's output.
type Drainer interface {
	Drain(Metric) error
}

// LogDrain is a Drainer implementation that logs the metrics to Stdout in
// l2met format.
type LogDrain struct {
	// Formatter to use to format the metric into a string before outputting.
	Formatter Formatter

	DrainFunc func(string)
	Logger    *log.Logger
}

// Drain logs the metric to Stdout.
func (d *LogDrain) Drain(m Metric) error {
	s := d.formatter().Format(m)
	d.drain(s)
	return nil
}

func (d *LogDrain) drain(s string) {
	if d.DrainFunc == nil {
		d.DrainFunc = func(s string) {
			d.logger().Println(s)
		}
	}

	d.DrainFunc(s)
}

func (d *LogDrain) formatter() Formatter {
	if d.Formatter == nil {
		d.Formatter = DefaultFormatter
	}

	return d.Formatter
}

func (d *LogDrain) logger() *log.Logger {
	if d.Logger == nil {
		d.Logger = log.New(os.Stdout, "", 0)
	}
	return d.Logger
}

// NullDrain is a Drainer implementation that does nothing.
type NullDrain struct{}

// Drain implements the Drainer interface.
func (d *NullDrain) Drain(m Metric) error { return nil }

type LocalStoreDrain struct {
	store map[string][]Metric
}

func (d *LocalStoreDrain) Store() map[string][]Metric {
	if d.store == nil {
		d.store = make(map[string][]Metric)
	}
	return d.store
}

func (d *LocalStoreDrain) Flush() {
	if d.store != nil {
		d.store = nil
	}
}

// Drain records metrics to the local store.
func (d *LocalStoreDrain) Drain(m Metric) error {
	var metrics []Metric
	if existingMetrics, ok := d.Store()[m.Name()]; ok {
		metrics = existingMetrics
	}
	d.Store()[m.Name()] = append(metrics, m)
	return nil
}
