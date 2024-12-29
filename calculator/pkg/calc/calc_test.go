package calc_test

import (
	calc "calculator/pkg/calc"
	"testing"
)

func TestCalc(t *testing.T) {

	for i, testCase := range testCases {
		res, err := calc.Calc(testCase.expression)

		if res != testCase.expectedResult {
			t.Fatalf("#%s: %s, %d", testCase.name, testCase.expression, i)

		} else if err != nil {
			if !testCase.expectedError {
				t.Fatalf("#%s: %s, %d", testCase.name, testCase.expression, i)
			}
		}
	}

}

var testCases = []struct {
	name           string // Тип теста
	expression     string // Математическое выражение
	expectedResult string // Ожидаемый результат
	expectedError  bool   // Ожидание ошибки
}{
	{
		name:           "operators",
		expression:     "(2+1)(2-1)(2*1)(2/1)(2^1)",
		expectedResult: "24",
		expectedError:  false,
	},
	{
		name:           "convertation",
		expression:     "   (2--2)* 5 (2+-1) / 3   ",
		expectedResult: "6.667",
		expectedError:  false,
	},
	{
		name:           "convertation",
		expression:     "(((1)))",
		expectedResult: "1",
		expectedError:  false,
	},
	{
		name:           "hard priority",
		expression:     "(((3*3)*3(((1/3)/(1/3))^(1/10)))/(27))+5(7+3)-70/(3+4)",
		expectedResult: "41",
		expectedError:  false,
	},
	{
		name:           "incorrect expression",
		expression:     "*2+(",
		expectedResult: "",
		expectedError:  true,
	},
	{
		name:           "incorrect expression",
		expression:     "0/0",
		expectedResult: "",
		expectedError:  true,
	},
	{
		name:           "incorrect expression",
		expression:     "math expression",
		expectedResult: "",
		expectedError:  true,
	},
}
