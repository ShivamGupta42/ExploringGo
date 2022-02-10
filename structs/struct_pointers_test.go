package structs

import "testing"

func TestStructTest_SUM(t *testing.T) {
	test := StructTest{1, 2}

	t.Run("testing access to struct variables using pointers", func(t *testing.T) {
		got := (&test).SUM()
		if got != 3 {
			t.Fail()
		}
	})

}
