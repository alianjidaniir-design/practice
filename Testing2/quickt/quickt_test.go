package quickt

import (
	"testing"
	"testing/quick"
)

var N = 100000

func TestWithItself(t *testing.T) {
	condition := func(a, b Point2D) bool {
		return ADD(a, b) == ADD(b, a)
	}
	err := quick.Check(condition, &quick.Config{MaxCount: N})
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
func TestThree(t *testing.T) {
	condition := func(a, b, c Point2D) bool {
		return ADD(ADD(a, b), c) == ADD(a, b)
	}
	err := quick.Check(condition, &quick.Config{MaxCount: N})
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
