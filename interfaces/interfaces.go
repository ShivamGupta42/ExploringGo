package interfaces

import "fmt"

type Startup bool

type Entrepreneur interface {
	Enterprise() Startup
}

type Shivam struct {
	charisma   int
	confidence int
	greatness  int
}

func (s *Shivam) Enterprise() Startup {
	return true
}

func UnicornCeo(name Entrepreneur) Startup {
	if name.Enterprise() {
		fmt.Println("Phod diya")

	} else {
		fmt.Println("ta ta tanna")
	}
	return name.Enterprise()
}
