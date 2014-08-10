package metrics

import (
	"fmt"
	"log"
	"os"
)

// Ensure implementations implement the Drain interface.
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

func (d *LogDrain) logger() *log.Logger {
	if d.Logger == nil {
		d.Logger = log.New(os.Stdout, "", 0)
	}
	return d.Logger
}

// NullDrain is a Drain implementation that does nothing.
type NullDrain struct{}

// Drain implements the Drainer interface.
func (d *NullDrain) Drain(m Metric) error { return nil }
