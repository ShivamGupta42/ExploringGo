package slices

func Sum(list []int) int {

	var sum int
	for _, val := range list {
		sum += val
	}
	return sum
}
