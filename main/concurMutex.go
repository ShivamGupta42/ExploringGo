package main

import (
	"fmt"
	"sync"
)

func deposit(balance *int, amount int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	mutex.Lock()
	*balance += amount
	mutex.Unlock()
	wg.Done()
}

func withdraw(balance *int, amount int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	mutex.Lock()
	*balance -= amount
	mutex.Unlock()
	wg.Done()
}

func mutexBank() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	balance := 100
	wg.Add(2)
	go deposit(&balance, 50, &mutex, &wg)
	go withdraw(&balance, 100, &mutex, &wg)
	wg.Wait()
	fmt.Println(balance)
}
