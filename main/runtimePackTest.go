package main

import (
	"fmt"
	"runtime"
)

func runtimePackOpr() {
	fmt.Printf("GOMAXPROCS is %d\n", runtime.GOMAXPROCS(3))
	fmt.Printf("GOMAXPROCS is %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("NumCPU is %d\n", runtime.NumCPU())
}
