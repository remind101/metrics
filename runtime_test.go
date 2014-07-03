package metrics

import "testing"

func Test_RuntimeSample_Print(t *testing.T) {
	r := NewRuntimeSample()
	r.Print()
}
