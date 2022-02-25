package fuzzing

import "testing"

func TestSwap(t *testing.T) {
	data := Swap()
	if data[0] != 10 || data[1] != 5 {
		t.Fail()
	}
}
