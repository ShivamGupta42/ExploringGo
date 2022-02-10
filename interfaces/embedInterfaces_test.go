package interfaces

import (
	"fmt"
	"testing"
)

func TestEmbedInterfaces(t *testing.T) {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5
	areaIntf = sq1
	fmt.Println("Type %T", areaIntf)
}
