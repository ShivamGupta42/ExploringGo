package pointers

import "fmt"

type DivisionError struct {
	Msg        string
	IntA, IntB int
}

func (d *DivisionError) Error() string {
	return "Division Error"
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionError{
			Msg:  fmt.Sprintf("cannot divide '%d' by zero", a),
			IntA: a, IntB: b,
		}
	}
	return a / b, nil
}
