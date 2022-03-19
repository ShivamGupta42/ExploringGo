package files

import "fmt"

func ForeverLoop(ch1, ch2 *chan int) {
loop:
	for {
		select {
		case <-*ch1:
			fmt.Println("Ch1 received")
			break loop
		case <-*ch2:
			fmt.Println("Ch2 received")
			break loop
		}
	}

	fmt.Println("Loop exited")
}
