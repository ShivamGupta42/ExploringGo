package main

import (
	"fmt"
	"time"
)

func receiveValues(s chan int) {
	for i := 0; i < 5; i++ {
		data := <-s
		fmt.Println(data)
	}
}

func main3() {
	sendLoopValues := make(chan int)

	go receiveValues(sendLoopValues)
	for i := 0; i < 5; i++ {
		sendLoopValues <- i
		time.Sleep(1000000000)
	}

}
