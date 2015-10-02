package metrics

import (
	"bytes"
	"errors"
	"log"
	"math"
	"os"
	"text/template"
)

// Ensure implementations implement the Drainer interface.
var (
	_ Drainer = &LogDrain{}
	_ Drainer = &NullDrain{}
	_ Drainer = &StatsdDrain{}
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

var ValueInvalidErr = errors.New("value must be one of [int, uint, int32, uint32, int64, uint64]")
var ValueOverflowErr = errors.New("value for uint64 too large to be converted to int64")
var InvalidMetricTypeErr = errors.New("metric type must be one of [count, sample, measure]")

type StatsdClient interface {
	Incr(name string, count int64) error
	Gauge(name string, value int64) error
	Timing(name string, ms int64) error
}

// StatsdDrain is a Drainer that records metrics to a statsd server.
type StatsdDrain struct {
	client   StatsdClient
	template *template.Template
}

// NewStatsdDrain takes a statsd client and a template string. The template
// string is used to construct the metric name with a Metric as its context.
func NewStatsdDrain(c StatsdClient, tmpl string) (*StatsdDrain, error) {
	t, err := template.New("stat").Parse(tmpl)
	if err != nil {
		return nil, err
	}

	return &StatsdDrain{client: c, template: t}, nil
}

// Drain records metrics to a statsd server.
func (d *StatsdDrain) Drain(m Metric) error {
	name, err := d.name(m)
	if err != nil {
		return err
	}

	value, err := vtoi(m.Value())
	if err != nil {
		return err
	}

	switch m.Type() {
	case "count":
		return d.client.Incr(name, value)
	case "sample", "measure":
		if m.Units() == "ms" {
			return d.client.Timing(name, value)
		} else {
			return d.client.Gauge(name, value)
		}
	default:
		return InvalidMetricTypeErr
	}
}

// name constructs a metric name with the StatsdDrain template.
func (d *StatsdDrain) name(m Metric) (string, error) {
	b := new(bytes.Buffer)
	if err := d.template.Execute(b, m); err != nil {
		return "", err
	}
	return b.String(), nil
}

// value coerces an interface value to an int64.
// The metric value MUST be an int or int64 or an error will be returned.
func vtoi(v interface{}) (int64, error) {
	switch i := v.(type) {
	case int:
		return int64(i), nil
	case uint:
		return int64(i), nil
	case int32:
		return int64(i), nil
	case uint32:
		return int64(i), nil
	case int64:
		return i, nil
	case uint64:
		if i <= math.MaxInt64 {
			return int64(i), nil
		} else {
			return 0, ValueOverflowErr
		}
	default:
		return 0, ValueInvalidErr
	}
}
