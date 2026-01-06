package division

import (
	"testing"
)

type myTest struct {
	a        int
	b        int
	resInt   int
	resFloat float64
}

var tests = []myTest{
	{a: 1, b: 2, resInt: 0, resFloat: 0.5},
	{a: 7, b: 14, resInt: 0, resFloat: 0.5},
	{a: 3, b: 3, resInt: 1, resFloat: 1.0},
	{a: 6, b: 6, resInt: 1, resFloat: 1.0},
	{a: 5, b: 2, resInt: 2, resFloat: 2.5},
	{a: 5, b: 4, resInt: 1, resFloat: 1.2},
}

func TestAll(t *testing.T) {
	t.Parallel()

	for _, test := range tests {
		intRes := intDiv(test.a, test.b)
		if intRes != test.resInt {
			t.Errorf("Expected %d, got %d", test.resInt, intRes)
		}
		floatResult := floatDiv(test.a, test.b)
		if floatResult != test.resFloat {
			t.Errorf("Expected %f, got %f", test.resFloat, floatResult)

		}
	}
}
