package tests

import "testing"

func TestRepeat(t *testing.T) {

	t.Run("repeat 3 times", func(t *testing.T) {
		want := "aaa"
		got := Repeat("a")
		acceptParamsAndValidate(t, want, got)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("A")
	}
}
