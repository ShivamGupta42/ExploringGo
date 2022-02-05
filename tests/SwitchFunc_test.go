package tests

import "testing"

var acceptParamsAndValidate = func(t testing.TB, want string, got string) {
	t.Helper() //Why am i able to call this on a pointer variable
	if got != want {
		t.Errorf("Wanted %s , Got %s", want, got)
	}
}

func TestGoSwitch(t *testing.T) {

	t.Run("testing eng", func(t *testing.T) {
		want := "Hi"
		got := GoSwitch("eng")
		acceptParamsAndValidate(t, want, got)
	})

	t.Run("testing frn", func(t *testing.T) {
		want := "Oui"
		got := GoSwitch("frn")
		acceptParamsAndValidate(t, want, got)
	})

}
