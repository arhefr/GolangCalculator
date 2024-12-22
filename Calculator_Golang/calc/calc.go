package calc

import (
	"Calculator_Golang/pkg"
	"slices"
	"strings"
)

// Вычисляет строчное математическое выражение. Поддерживает такие операции как +, -, *, /, а также скобки ( ).
func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	nums, ops, err := pkg.ParserExpression(expression)
	if err != nil {
		return 0.0, err
	}

	pkg.SortOperators(ops)
	for len(ops) != 0 { // Пошагово выполняем операции, пока они не закончатся
		op := ops[0]
		I := op.Index

		operation, err := op.MakeOperation(nums)
		if err != nil {
			return 0.0, err
		}

		nums = slices.Insert(pkg.RemoveElement(pkg.RemoveElement(nums, I), I), I, operation) // Заменяем оперируемые числа на результат операции
		op.FixIndices(ops)
		ops = ops[1:] // Убираем использованный оператор

	}

	if len(nums) != 1 {
		return 0.0, pkg.ErrIncorrectExpression
	}

	return nums[0], nil
}
