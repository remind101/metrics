package metrics

import (
	"fmt"
	"time"
)

func init() {
	Source = ""
	DefaultDrain.(*LogDrain).DrainFunc = func(s string) {
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

func ExampleTime() {
	t := Time("request.time")
	t.NowFunc = func() time.Time {
		return t.start.Add(527 * time.Millisecond)
	}
	t.Done()
	// Output:
	// measure#request.time=527ms
}
