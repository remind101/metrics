package metrics

import "testing"

func Test_Timer_Done(t *testing.T) {
	timer := NewTimer("request.time")
	timer.Done()
}
