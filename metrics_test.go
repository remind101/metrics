package metrics

import (
	"fmt"
	"testing"
)

func init() {
	Drain.(*LogDrain).DrainFunc = func(s string) {
		fmt.Println(s)
	}
}

func ExampleCount() {
	Count("user.signup", 1)
	// Output:
	// count#user.signup=1
}

func ExampleSample() {
	Sample("goroutine", 1, "")
	// Output:
	// sample#goroutine=1
}

func ExampleMeasure() {
	Measure("request.time.2xx", 12.14, "ms")
	// Output:
	// measure#request.time.2xx=12.14ms
}

func TestTime(t *testing.T) {
	timer := Time("request.time")
	timer.Done()
}
