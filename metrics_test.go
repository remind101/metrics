package metrics_test

import (
	metrics "."
	"testing"
)

func TestLocalStoreDrain(t *testing.T) {
	original := metrics.DefaultDrain
	metrics.DefaultDrain = &metrics.LocalStoreDrain{}
	defer func() {
		metrics.DefaultDrain = original
	}()

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
