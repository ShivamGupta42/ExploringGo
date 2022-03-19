package concurrency

func FibGenerator(n int) <-chan int {
	fibChan := make(chan int)

	go func() {
		for i, j, k := 0, 1, 0; k < n; k++ {
			fibChan <- i
			i, j = j, i+j
		}
		close(fibChan)
	}()
	return fibChan
}
