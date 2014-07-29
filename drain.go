package metrics

import (
	"fmt"
	"log"
	"os"

	"github.com/remind101/metrics/vendor/g2s"
)

// Insure interfaces are implemented.
var (
	_ Drainer = &LogDrain{}
	_ Drainer = &StatsDrain{}
)

// Drainer is an interface that can drain a metric to it's output.
type Drainer interface {
	Drain(Metric) error
}

// LogDrain is a Drainer implementation that logs the metrics to Stdout in
// l2met format.
type LogDrain struct {
	Logger *log.Logger
}

// Drain logs the metric to Stdout.
func (d *LogDrain) Drain(m Metric) error {
	s := fmt.Sprintf("%s#%s=%v%s", m.Type(), m.Name(), m.Value(), m.Units())

	if Source != "" {
		s = fmt.Sprintf("source=%s %s", Source, s)
	}

	d.logger().Println(s)
	return nil
}

// SetPrefix sets the prefix for the logger.
func (d *LogDrain) SetPrefix(prefix string) {
	d.logger().SetPrefix(prefix)
}

func (d *LogDrain) logger() *log.Logger {
	if d.Logger == nil {
		d.Logger = log.New(os.Stdout, "", 0)
	}
	return d.Logger
}

// StatsDrain is a Drain implementation that emits the metrics to StatsD.
type StatsDrain struct {
	statter g2s.Statter
}

// NewStatsDrain returns a new StatsDDrain.
func NewStatsDrain(endpoint string) *StatsDrain {
	s, err := g2s.Dial("udp", endpoint)
	if err != nil {
		s = g2s.Noop()
	}
	return &StatsDrain{statter: s}
}

// Drain drains the metrics to StatsD.
func (d *StatsDrain) Drain(m Metric) error {
	name := m.Name()

	if Source != "" {
		name = fmt.Sprintf("%s.%s", Source, name)
	}

	switch m.Type() {
	case "sample", "measure":
		d.statter.Gauge(1.0, name, fmt.Sprintf("%v", m.Value()))
	case "count":
		if value, ok := m.Value().(int); ok {
			d.statter.Counter(1.0, name, value)
		}
	default:
	}
	return nil
}
