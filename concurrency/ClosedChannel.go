package concurrency

import "fmt"

func sendCash(buyer chan int) {
	var i int
	for i = 0; i <= 3; i++ {
		buyer <- i
	}
	close(buyer)
}

func ClosedChannelOrchestrator() {
	money := make(chan int)

	go sendCash(money)

	for seller := range money {
		fmt.Println(seller)
	}
}
