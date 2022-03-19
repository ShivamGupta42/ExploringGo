package concurrency

import "testing"

func TestFibGenerator(t *testing.T) {
	fibChan := FibGenerator(5)

	t.Run("Testing Correctness & Capacity Constraint", func(t *testing.T) {
		a := <-fibChan
		b := <-fibChan
		c := <-fibChan
		d := <-fibChan
		e := <-fibChan

		if a != 0 && b != 1 && c != 1 && d != 2 && e != 3 {
			t.Errorf("Incorrect data %d %d %d %d %d", a, b, c, d, e)
		}

		if _, isOpen := <-fibChan; isOpen {
			t.Errorf("Channel wasn't closed")
		}
	})

}
