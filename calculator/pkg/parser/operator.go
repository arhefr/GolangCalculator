package parser

import (
	"math"
	"sort"
)

// Структура для операторов
type Operator struct {
	name     string // Тип оператора
	Index    int    // Его индекс в примере
	priority int    // Приоритет его выполнения
}

// Инициализация класса operator
func NewOperator(name string, index, priority int) Operator {
	return Operator{name, index, priority}
}

// Сортировка операторов по приоритету
func SortOperators(ops []Operator) {
	sort.Slice(ops, func(i, j int) bool {
		return ops[i].priority > ops[j].priority
	})
}

// Вычисление одной операции
func (op Operator) MakeOperation(nums []float64) (float64, error) {
	var resOp float64
	n1, n2 := nums[op.Index], nums[op.Index+1]

	switch op.name {
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
	}

	return resOp, nil
}

// Понижает индексы операторов
func (op Operator) FixIndices(ops []Operator) {
	for I, opNext := range ops[1:] {
		if opNext.Index > op.Index {
			ops[I+1].Index--
		}
	}
}
