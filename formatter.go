package metrics

import "fmt"

// DefaultFormatter is the default Formatter to use.
var DefaultFormatter = Formatter(&l2metFormatter{})

// Formatter is an interface for formatting a metric into a string.
type Formatter interface {
	Format(Metric) string
}

// l2metFormatter is an implementation of the Formatter interface that formats the metric
// in the l2met format:
//
//	count#user.signup=1
//	measure#request.time=6ms
type l2metFormatter struct{}

func (f *l2metFormatter) Format(m Metric) string {
	s := fmt.Sprintf("%s#%s=%v%s", m.Type(), m.Name(), m.Value(), m.Units())

	if Source != "" {
		s = fmt.Sprintf("source=%s %s", Source, s)
	}

	return s
}
