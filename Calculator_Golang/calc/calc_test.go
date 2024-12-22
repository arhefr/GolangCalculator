package calc_test

import (
	c "Calculator_Golang/calc"
	"fmt"
	"testing"
)

type Test struct {
	expression  string
	result      float64
	expectedErr bool
}

var Tests = []Test{
	{"2**2**2**2", 256.0, false},
	{"2+2*(10-8)", 6.0, false},
	{"2+2*2", 6.0, false},
	{"((1+4) * (1+2) +10) *4", 100.0, false},
	{"(70/7) * 10 /((3+2) * (3+7)) -2", 0.0, false},
	{"((7+1) / (2+2) * 4) / 8 * (32 - ((4+12)*2)) -1", -1.0, false},
	{"((1+2)*(5*(7+3) - 70 / (3+4) * (1+2)) - (8-1)) + (10 * (5-1 * (2+3)))", 53.0, false},

	{"2+", 0.0, true},
	{"2+(9-8", 0.0, true},
	{"2+aboba42", 0.0, true},
}

func TestCalc(t *testing.T) {
	for i, Test := range Tests {
		res, err := c.Calc(Test.expression)
		if res != Test.result {
			fmt.Println(Test.result, res)
			t.Fatalf("#%d", i)
		} else if err != nil {
			if !Test.expectedErr {
				fmt.Println(Test.result, res)
				t.Fatalf("#%d", i)
			}
		}
	}
}
