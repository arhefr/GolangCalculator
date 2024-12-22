package pkg

import (
	"math"
	"sort"
)

type operator struct {
	name     string // Тип оператора
	Index    int    // Его индекс в примере
	priority int    // Приоритет его выполнения
}

func newOperator(name string, index, priority int) operator {
	return operator{name, index, priority}
}

// Сортировка операторов по приоритету
func SortOperators(ops []operator) {
	sort.Slice(ops, func(i, j int) bool {
		return ops[i].priority > ops[j].priority
	})
}

// Вычисление одной операции
func (op operator) MakeOperation(nums []float64) (float64, error) { // Вычисляемая операция и его результат
	var resOp float64
	opName, opIndex := op.name, op.Index
	n1, n2 := nums[opIndex], nums[op.Index+1]

	switch opName {
	case "^":
		resOp = math.Pow(n1, n2)
	case "*":
		resOp = n1 * n2
	case "/":
		if n2 == 0.0 {
			return 0.0, ErrDivisionByZero // Делить на ноль нельзя
		}
		resOp = n1 / n2
	case "+":
		resOp = n1 + n2
	case "-":
		resOp = n1 - n2

	default:
		return 0.0, ErrIncorrectExpression // Неизвестный оператор
	}

	return resOp, nil
}

// Корректирует индексы операторов
func (op operator) FixIndices(ops []operator) {
	opIndex := op.Index
	for I, opNext := range ops[1:] {
		if opNext.Index > opIndex {
			ops[I+1].Index--
		}
	}
}
