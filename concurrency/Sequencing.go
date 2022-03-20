package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

// CookInfo Struct with a waiting channel
type CookInfo struct {
	foodCooked     string
	waitForPartner chan bool
}

func CookFood(name string) <-chan CookInfo {

	cookChannel := make(chan CookInfo)
	wait := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			cookChannel <- CookInfo{fmt.Sprintf("%s %s", name, "Done"), wait}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

			<-wait
		}
	}()

	return cookChannel
}

func FanInCookInfo(channel1, channel2 <-chan CookInfo) <-chan CookInfo {
	channel := make(chan CookInfo)

	go func() {
		for {
			channel <- <-channel1
		}
	}()

	go func() {
		for {
			channel <- <-channel2
		}
	}()

	return channel
}

func Orchestrate() {
	gameChannel := FanInCookInfo(CookFood("Player 1 : "), CookFood("Player 2 :"))

	for round := 0; round < 3; round++ {
		food1 := <-gameChannel
		fmt.Println(food1.foodCooked)

		food2 := <-gameChannel
		fmt.Println(food2.foodCooked)

		food1.waitForPartner <- true
		food2.waitForPartner <- true

		fmt.Printf("Done with round %d\n", round+1)
	}

	fmt.Println("Done with the competition")
}
