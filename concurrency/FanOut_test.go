package concurrency

import (
	"fmt"
	"testing"
)

func TestFanOut(t *testing.T) {
	randVals := GenerateRandom(5)
	chan1 := FanOut(randVals)
	chan2 := FanOut(randVals)
	fanIn := FanIn(chan1, chan2)
	var v int
	var ok bool
	for i := 0; i < 7; i++ {
		v, ok = <-fanIn
		fmt.Println(v)
	}

	if ok {
		t.Fail()
	}

}
