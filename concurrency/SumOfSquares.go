package concurrency

import (
	"fmt"
)

func sumOfSquaresHelper(c, quit chan int) {
	updates := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			updates <- i * i
		}
	}()

	for {
		select {
		case val := <-updates:
			c <- val
		case <-quit:
			return
		}
	}
}

func SumOfSquares() {
	mychannel := make(chan int)
	quitchannel := make(chan int)
	sum := 0
	go func() {
		for i := 1; i < 6; i++ {
			sum += <-mychannel
		}
		fmt.Println(sum)
		quitchannel <- 1
	}()
	sumOfSquaresHelper(mychannel, quitchannel)
}
