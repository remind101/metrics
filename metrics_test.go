package metrics_test

import (
	"errors"
	"fmt"
	"testing"

	metrics "."
)

func setMetricsDrain(d metrics.Drainer) func() {
	original := metrics.DefaultDrain
	metrics.DefaultDrain = d
	return func() {
		metrics.DefaultDrain = original
	}
}

func setMetricsSource(s string) func() {
	original := metrics.Source
	metrics.Source = s
	return func() {
		metrics.Source = original
	}
}

func TestLocalStoreDrain(t *testing.T) {
	cleanup := setMetricsDrain(&metrics.LocalStoreDrain{})
	defer cleanup()

	store := metrics.DefaultDrain.(*metrics.LocalStoreDrain).Store()
	const key = "user.signup"

	// increment our key twice
	for i := 0; i < 2; i++ {
		metrics.Count(key, 1)
	}
	if len(store[key]) != 2 {
		t.Error(
			"For", key,
			"expected", 2,
			"got", len(store[key]),
		)
	}

	metrics.Measure(key, 127, "ms")
	if len(store[key]) != 3 {
		t.Error(
			"For", key,
			"expected", 3,
			"got", len(store[key]),
		)
	}

	metricsMap := make(map[string]int)
	if metrics, ok := store[key]; ok {
		for _, metric := range metrics {
			if _, ok := metricsMap[metric.Type()]; ok {
				metricsMap[metric.Type()] += 1
			} else {
				metricsMap[metric.Type()] = 1
			}
		}
	}

	if metricsMap["count"] != 2 {
		t.Error(
			"For metric", "count",
			"expected", 2,
			"got", metricsMap["count"],
		)
	}

	if metricsMap["measure"] != 1 {
		t.Error(
			"For metric", "measure",
			"expected", 1,
			"got", metricsMap["measure"],
		)
	}
}

func TestLocalStoreDrainFlush(t *testing.T) {
	cleanup := setMetricsDrain(&metrics.LocalStoreDrain{})
	defer cleanup()
	const key = "key"

	metrics.Count(key, 1)
	if len(metrics.DefaultDrain.(*metrics.LocalStoreDrain).Store()) != 1 {
		t.Error("Error recording metric")
	}

	metrics.DefaultDrain.(*metrics.LocalStoreDrain).Flush()
	if len(metrics.DefaultDrain.(*metrics.LocalStoreDrain).Store()[key]) != 0 {
		t.Error("Error flushing LocalStoreDrain")
	}

}

var UnexpectedFuncCall = errors.New("unexpected function call")

type stub func(string, int64) error

var expect = func(t *testing.T, en string, ev int64) stub {
	return func(n string, v int64) error {
		if got, want := n, en; got != want {
			return fmt.Errorf("metric name: expected %s; got %s", got, want)
		}
		if got, want := v, ev; got != want {
			return fmt.Errorf("metric value: expected %d; got %d", got, want)
		}
		return nil
	}
}

type MockStatsdClient struct {
	IncrFunc   stub
	GaugeFunc  stub
	TimingFunc stub
}

func (c *MockStatsdClient) Incr(n string, v int64) error {
	if c.IncrFunc != nil {
		return c.IncrFunc(n, v)
	} else {
		return UnexpectedFuncCall
	}
}

func (c *MockStatsdClient) Gauge(n string, v int64) error {
	if c.GaugeFunc != nil {
		return c.GaugeFunc(n, v)
	} else {
		return UnexpectedFuncCall
	}
}

func (c *MockStatsdClient) Timing(n string, v int64) error {
	if c.TimingFunc != nil {
		return c.TimingFunc(n, v)
	} else {
		return UnexpectedFuncCall
	}
}

func TestStatsdDrain(t *testing.T) {
	mc := &MockStatsdClient{}

	d, err := metrics.NewStatsdDrain(mc, "{{.Name}}.source__{{.Source}}__")
	if err != nil {
		t.Fatal(err)
	}

	drainCleanup := setMetricsDrain(d)
	defer drainCleanup()

	sourceCleanup := setMetricsSource("test")
	defer sourceCleanup()

	mc.IncrFunc = expect(t, "requests.count.source__test__", 5)
	if err := metrics.Count("requests.count", 5); err != nil {
		t.Error(err)
	}
	mc.GaugeFunc = expect(t, "requests.sample.source__test__", 50)
	if err := metrics.Sample("requests.sample", 50, ""); err != nil {
		t.Error(err)
	}

	mc.GaugeFunc = expect(t, "requests.measure.source__test__", 500)
	if err := metrics.Measure("requests.measure", 500, ""); err != nil {
		t.Error(err)
	}

	mc.TimingFunc = expect(t, "requests.timing.source__test__", 5000)
	if err := metrics.Measure("requests.timing", 5000, "ms"); err != nil {
		t.Error(err)
	}

	mc.TimingFunc = expect(t, "requests.timing.source__test__", 101)
	if err := metrics.Measure("requests.timing", 100.925, "ms"); err != nil {
		t.Error(err)
	}
}
