package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func Merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	left := MergeSort(data[:mid])
	right := MergeSort(data[mid:])
	return Merge(left, right)
}

func average(times []int64) float64 {
	length := len(times)
	var sum int64 = 0
	for i := 0; i < length; i++ {
		sum += times[i]
	}
	return float64(sum) / float64(length)
}

func runMergeSortIteration() time.Duration {
	start := time.Now()
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = rand.Intn(10)
	}
	MergeSort(data)
	return time.Since(start)
}

func seqMergeSortExecute() {
	rand.Seed(time.Now().UnixNano())
	totalRuns := 1000000
	var times = make([]int64, totalRuns)
	for i := 0; i < totalRuns; i++ {
		times[i] = runMergeSortIteration().Microseconds()
	}
	fmt.Printf("Total average time taken by sequential Merge Sort in %d runs : %f\n", totalRuns, average(times))
}
