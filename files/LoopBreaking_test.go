package files

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestForeverLoop(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(2)

	go func(val int) {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		ch1 <- val
	}(123)

	go func(val int) {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
		ch2 <- val
	}(123)

	ForeverLoop(&ch1, &ch2)
	wg.Wait()
	fmt.Println("Test executed")
}
