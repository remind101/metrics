package metrics

import "fmt"

const (
	MetricCount   MetricType = "count"
	MetricSample  MetricType = "sample"
	MetricMeasure MetricType = "measure"
)

type MetricType string

type Metric struct {
	// The name of the metric. (e.g. request.time.2xx)
	Name string

	// The type of metric.
	Type MetricType

	// The value of the metric.
	Value interface{}

	// The units of the metric.
	Units string
}

// String returns the string representation of the metric in l2met format.
func (m *Metric) String() string {
	return fmt.Sprintf("%s#%s=%v%s", m.Type, m.Name, m.Value, m.Units)
}

// Print puts the metric onto stdout.
func (m *Metric) Print() {
	Logger.Println(m.String())
}
