package main

import "fmt"

func main() {
	/* MergeSort programs
	concurMergeSortExecute()
	seqMergeSortExecute()
	*/

	/*
		executeWaitGroup()
	*/

	//mutexBank()

	//runtimePackOpr()

	//fmt.Println(stringAlgos.Swap("Hello", "World"))

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answeree"]
	fmt.Println("The value:", v, "Present?", ok)

}
