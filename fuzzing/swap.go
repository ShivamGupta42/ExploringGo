package fuzzing

func Swap() []int {
	a, b := 5, 10
	a, b = b, a
	return []int{a, b}
}
