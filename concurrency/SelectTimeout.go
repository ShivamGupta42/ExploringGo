package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func SelectTimeout() {
	channel := make(chan string)

	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
			time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
			channel <- "Dynamite diffused"
		}
	}()

	for {
		select {
		case msg := <-channel:
			fmt.Println(msg)
			return
		case <-time.After(time.Duration(rand.Intn(50)) * time.Millisecond):
			fmt.Println("Dynamite exploded")
			return
		}
	}

}
