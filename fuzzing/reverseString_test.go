package fuzzing

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	want := "egod"
	got := Reverse("doge")

	if want != got {
		t.Fail()
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {

		t.Run("Testing double reversal", func(t *testing.T) {
			want := "doge"
			got := Reverse(Reverse(want))

			if want != got {
				t.Errorf("Got: %q, Wanted: %q", got, want)
			}
		})

		t.Run("Testing reversal is also UTF-8", func(t *testing.T) {
			got := Reverse(orig)
			if utf8.ValidString(orig) && !utf8.ValidString(got) {
				t.Errorf("Invalid string produced %q", got)
			}
		})
	})
}
