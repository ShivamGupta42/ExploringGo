package main

import (
	"fmt"
	"math/rand"
	"time"
)

func concurMerge(left, right []int) []int {
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

func concurMergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	done := make(chan bool)
	mid := len(data) / 2

	var left []int
	go func() {
		left = concurMergeSort(data[:mid])
		done <- true
	}()

	right := concurMergeSort(data[mid:])
	<-done
	return concurMerge(left, right)
}

func runConcurMergeSortIteration() time.Duration {
	start := time.Now()
	data := make([]int, 10)
	for i := 0; i < 10; i++ {
		data[i] = rand.Intn(10)
	}
	concurMergeSort(data)
	return time.Since(start)
}

func concurMergeSortExecute() {
	rand.Seed(time.Now().UnixNano())
	totalRuns := 1000000
	var times = make([]int64, totalRuns)
	for i := 0; i < totalRuns; i++ {
		times[i] = runConcurMergeSortIteration().Microseconds()
	}
	fmt.Printf("Total average time taken by concurrent Merge Sort in %d runs : %f\n", totalRuns, average(times))
}
