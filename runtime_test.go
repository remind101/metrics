package metrics

import "testing"

func Test_RuntimeSample_Drain(t *testing.T) {
	r := NewRuntimeSample()
	r.Drain()
}
