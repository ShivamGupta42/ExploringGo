package concurrency

import "testing"

func TestFanIn(t *testing.T) {
	chan1 := GenerateRandom(5)
	chan2 := GenerateRandom(5)

	fanChan := FanIn(chan1, chan2)

	k := 0
	for i := 0; i < 10; i++ {
		k++
		<-fanChan
	}

	if k != 10 {
		t.Fail()
	}
}
