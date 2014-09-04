package metrics

// Metric represents an individual count/sample/measurement and encapsulates information
// about it.
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

// metric is a generic implementation of the Metric interface.
type metric struct {
	name  string
	typ   string
	value interface{}
	units string
}

// Methods to implement the Metric interface
func (m *metric) Name() string       { return m.name }
func (m *metric) Type() string       { return m.typ }
func (m *metric) Value() interface{} { return m.value }
func (m *metric) Units() string      { return m.units }
