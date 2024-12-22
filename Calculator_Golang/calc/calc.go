package calc

import (
	"Calculator_Golang/pkg"
	"fmt"
	"log"
	"slices"
	"time"
)

// Вычисляет строчное математическое выражение. Поддерживает такие операции как +, -, *, /, а также скобки ( ).
func Calc(expression string) (float64, error) {
	start := time.Now()

	expression = pkg.FixExpression(expression)
	log.Println("Пример: ", expression)

	nums, ops, err := pkg.ParserExpression(expression)
	fmt.Println(nums, ops)
	if err != nil && len(expression) < 3 {
		log.Println(time.Since(start).Milliseconds(), "мс. ", "Ошибка: ", err)
		log.Println("")
		return 0.0, err
	}

	pkg.SortOperators(ops)
	for len(ops) != 0 { // Пошагово выполняем операции, пока они не закончатся
		op := ops[0]
		I := op.Index

		operation, err := op.MakeOperation(nums)
		if err != nil {
			log.Println(time.Since(start).Milliseconds(), "мс. ", "Ошибка: ", err)
			log.Println("")
			return 0.0, err
		}

		nums = slices.Insert(pkg.RemoveElement(pkg.RemoveElement(nums, I), I), I, operation) // Заменяем оперируемые числа на результат операции
		op.FixIndices(ops)
		ops = ops[1:] // Убираем использованный оператор

		fmt.Println(nums, ops)
	}

	if len(nums) != 1 {
		log.Println(time.Since(start).Milliseconds(), "мс. ", "Ошибка: ", err)
		log.Println("")
		return 0.0, pkg.ErrIncorrectExpression
	}

	log.Println(time.Since(start).Milliseconds(), "мс. ", "Результат: ", nums[0])
	return nums[0], nil
}
