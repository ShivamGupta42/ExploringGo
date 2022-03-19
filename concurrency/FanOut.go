package concurrency

func FanOut(input <-chan int) <-chan int {
	readChan := make(chan int)

	go func() {
		for i := range input {
			readChan <- i
		}
		close(readChan)
	}()
	return readChan
}
