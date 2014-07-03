package metrics

import "testing"

func TestMetricString(t *testing.T) {
	tests := []struct {
		metric   *Metric
		expected string
	}{
		{
			&Metric{Name: "request.time", Type: MetricMeasure, Value: 120.12, Units: "ms"},
			"measure#request.time=120.12ms",
		},
		{
			&Metric{Name: "goroutine", Type: MetricCount, Value: 1},
			"count#goroutine=1",
		},
	}

	for i, test := range tests {
		if test.metric.String() != test.expected {
			t.Errorf("%i: Want %v; Got %v", i, test.expected, test.metric.String())
		}
	}
}
