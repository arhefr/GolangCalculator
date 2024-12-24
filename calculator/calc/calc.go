package calc

import (
	"calculator/internal/pkg"
	"calculator/internal/pkg/parser"
	"log"
	"slices"
	"time"
)

// Вычисляет математическое выражение. Поддерживает такие операции как +, -, *, /, ^, (, ).
func Calc(expression string) (string, error) {
	start := time.Now()

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

			nums = slices.Insert(pkg.RemoveElement(pkg.RemoveElement(nums, I), I), I, operation)
			op.FixIndices(ops)
			ops = ops[1:]

			if len(ops) == 0 {
				result := parser.FixNum(nums[0])
				log.Println(time.Since(start).Microseconds(), "мкс |", expression, ":", result)
				return result, nil
			}
		}
	}

	result := parser.FixNum(nums[0])
	log.Println(time.Since(start).Microseconds(), "мкс |", expression, ":", result)
	return result, nil
}
