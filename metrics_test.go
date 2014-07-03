package metrics

import "testing"

func TestCount(t *testing.T) {
	Count("user.signup", 1)
}

func TestSample(t *testing.T) {
	Sample("goroutine", 1, "")
}

func TestMeasure(t *testing.T) {
	Measure("request.time.2xx", 12.14, "ms")
}

func TestTime(t *testing.T) {
	timer := Time("request.time")
	timer.Done()
}
