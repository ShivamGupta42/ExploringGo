package main

import (
	"fmt"
	"sync"
	"time"
)

func executeWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		time.Sleep(time.Millisecond * 1000)
		fmt.Println("Executing the goroutine")
		wg.Done()
	}()

	wg.Wait()
}
