package main

import (
	"fmt"
	"sync"
	"time"
)

func printTable(wg *sync.WaitGroup, number int) {
	for i := 1; i <= 12; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
		time.Sleep(50 * time.Millisecond)
	}
	wg.Done()
}

func concurMultiplication() {
	var wg sync.WaitGroup
	wg.Add(12)

	for number := 1; number <= 12; number++ {
		go printTable(&wg, number)
	}
	wg.Wait()
}
