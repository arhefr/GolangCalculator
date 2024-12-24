package calc_test

import (
	c "calculator/calc"
	"testing"
)

func TestCalc(t *testing.T) {

	testCases := []struct {
		name        string // Тип теста
		expression  string // Математическое выражение
		result      string // Решение
		expectedErr bool   // Ожидание ошибки
	}{
		{"operators", "1+1", "2", false},
		{"operators", "1-1", "0", false},
		{"operators", "1*1", "1", false},
		{"operators", "1/3", "0.333", false},
		{"operators", "64^(1/3)", "4", false},

		{"simple priority", "2+2*2", "6", false},
		{"hard priority", "((1+2)(5(7+3)-70/(3+4)(1+2))-(8-1))+(10(5-1(2+3)))", "53", false},

		{"convert expression", "1     +     1", "2", false},
		{"convert expression", "((1))", "1", false},
		{"convert expression", "5((1))", "5", false},
		{"convert operators", "1--1", "2", false},
		{"convert operators", "1+-1", "0", false},

		{"invalid expression", "1+", "", true},
		{"invalid expression", "1*(", "", true},
		{"invalid expression", "a", "", true},
		{"invalid expression", "1/0", "", true},
	}

	for i, testCase := range testCases {
		res, err := c.Calc(testCase.expression)

		if res != testCase.result {
			t.Fatalf("#%s: %s, %d", testCase.name, testCase.expression, i)

		} else if err != nil {
			if !testCase.expectedErr {
				t.Fatalf("#%s: %s, %d", testCase.name, testCase.expression, i)
			}
		}
	}

}
