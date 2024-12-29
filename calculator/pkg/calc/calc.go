package calc

import (
	"calculator/pkg/parser"
	"fmt"
	"slices"
)

// Вычисляет математическое выражение. Поддерживает такие операции как +, -, *, /, ^, (, ).
func Calc(expression string) (string, error) {

	nums, ops, err := parser.ParserExpression(expression)
	if err != nil {
		return "", err
	}

	if len(nums) != 1 && len(ops) != 0 {
		for {
			op := ops[0]
			I := op.Index

			operation, err := op.MakeOperation(nums)
			if err != nil {
				return "", err
			}

			nums = slices.Insert(removeElement(removeElement(nums, I), I), I, operation)
			op.FixIndices(ops)
			ops = ops[1:]

			if len(ops) == 0 {
				result := convertNum(nums[0])
				return result, nil
			}
		}
	}

	result := convertNum(nums[0])
	return result, nil
}

// Приводит число в нужный вид
func convertNum(num float64) string {
	numString := fmt.Sprintf("%.3f", num)

	cnt := 0
	for i := len(numString) - 1; i >= 0; i-- {
		sym := string(numString[i])
		if sym == "." {
			cnt++
			break
		}

		if sym == "0" {
			cnt++
		} else {
			break
		}
	}

	return numString[:(len(numString) - cnt)]
}

// Убирает элемент из слайса по индексу
func removeElement(slice []float64, s int) []float64 {
	return append(slice[:s], slice[s+1:]...)
}
