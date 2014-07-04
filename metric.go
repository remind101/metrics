package metrics

type Metric interface {
	// Name returns the name of the metric (e.g. request.time.2xx)
	Name() string

	// Type returns the type of metric.
	Type() string

	// Value returns the value of the metric.
	Value() interface{}

	// Units returns the units of the metric.
	Units() string
}

// coreMetric is a generic implementation of the Metric interface.
type coreMetric struct {
	name  string
	typ   string
	value interface{}
	units string
}

// Methods to implement the Metric interface
func (m *coreMetric) Name() string       { return m.name }
func (m *coreMetric) Type() string       { return m.typ }
func (m *coreMetric) Value() interface{} { return m.value }
func (m *coreMetric) Units() string      { return m.units }
