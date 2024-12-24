package parser_test

import (
	"calculator/internal/pkg"
	"calculator/internal/pkg/parser"
	"testing"
)

func TestParserExpression(t *testing.T) {

	testCases := []struct {
		expression string         // Математическое выражение
		nums       []float64      // Ожидаемый список чисел
		ops        []pkg.Operator // Ожидаемый список операторов

	}{
		{"-2*(-2)", []float64{-2.0, -2.0}, []pkg.Operator{pkg.NewOperator("*", 0, 5)}},
		{"2+2*2", []float64{2.0, 2.0, 2.0}, []pkg.Operator{pkg.NewOperator("*", 1, 1), pkg.NewOperator("+", 0, 0)}},
		{"-2(-2)", []float64{-2.0, -2.0}, []pkg.Operator{pkg.NewOperator("*", 0, 5)}},
		{"(-2)(-2)", []float64{-2.0, -2.0}, []pkg.Operator{pkg.NewOperator("*", 0, 5)}},
		{"(-2)^(-2)", []float64{-2.0, -2.0}, []pkg.Operator{pkg.NewOperator("^", 0, 2)}},
	}

	for i, testCase := range testCases {
		nums, _, _ := parser.ParserExpression(testCase.expression)

		for n := 0; i < len(nums); i++ {
			if nums[n] != testCase.nums[n] {
				t.Fatalf("#%s: %d", testCase.expression, i)
			}
		}

	}

}
