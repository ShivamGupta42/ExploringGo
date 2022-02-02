package slices

import "testing"

func acceptAndValidate(t *testing.T, want interface{}, got interface{}) {
	t.Helper()
	if got != want {
		t.Fail()
	}
}

func TestSum(t *testing.T) {

	t.Run("testing sum with 4 elements", func(t *testing.T) {
		want := 10
		got := Sum([]int{1, 2, 3, 4})
		acceptAndValidate(t, want, got)
	})

}
