package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/* IDEA 💡
   When multiple threads are in contention
   And you want to communicate back to the winning thread
*/

var wg sync.WaitGroup

func Race(raceUpdates, quit chan string, i int) {

	raceUpdates <- fmt.Sprintf("%d Car started racing", i)
	for {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		quit <- fmt.Sprintf("%d Car finished racing", i)
		fmt.Println(<-quit)
		wg.Done()
	}
}

func RaceOrchestrator() {
	raceUpdates := make(chan string)
	quit := make(chan string)

	wg.Add(1)

	for i := 0; i < 3; i++ {
		go Race(raceUpdates, quit, i)
	}

	for {
		select {
		case update := <-raceUpdates:
			fmt.Println(update)
		case quitMsg := <-quit:
			fmt.Println(quitMsg)
			quit <- "You won Boy"
			wg.Wait()
			return
		}
	}
}
