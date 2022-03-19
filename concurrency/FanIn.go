package concurrency

import (
	"math/rand"
	"time"
)

func GenerateRandom(n int) <-chan int {
	rand.Seed(time.Now().UnixNano())

	randVals := make(chan int)

	go func() {
		for k := 0; k < n; k++ {
			randVals <- rand.Intn(50)
		}
		close(randVals)
	}()

	return randVals
}

func FanIn(chan1, chan2 <-chan int) <-chan int {
	fanInChan := make(chan int)

	go func() {
		for {
			fanInChan <- <-chan1
		}
	}()

	go func() {
		for {
			fanInChan <- <-chan2
		}
	}()

	return fanInChan
}
