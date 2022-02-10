package interfaces

import "testing"

func TestUnicornCeo(t *testing.T) {
	typeShivam := Shivam{charisma: 100, confidence: 100, greatness: 100}
	UnicornCeo(&typeShivam)
}
