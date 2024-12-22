package pkg

import (
	"strconv"
	"strings"
)

type symRep struct {
	sym string
	fix string
}

func newSymRep(sym, fix string) symRep {
	return symRep{sym, fix}
}

var fixSymphols = []symRep{
	newSymRep(" ", ""),
	newSymRep(",", "."),
	newSymRep("--", "+"),
	newSymRep("+-", "-"),
	newSymRep("**", "^"),
}

func FixExpression(expression string) string {
	for _, fixSym := range fixSymphols {
		expression = strings.ReplaceAll(expression, fixSym.sym, fixSym.fix)
	}

	return expression
}

// Парсинг всех чисел и операторов
func ParserExpression(expression string) ([]float64, []operator, error) { // Массив с числами и операторами
	var nums []float64
	var ops []operator
	var prioAdd, allowMinus int
	var num string

	for I, sym := range expression { // Проходимся по всем символам выражения
		S := string(sym)
		T := tokenize(S)

		if T == 4 { // Если символ - скобка
			allowMinus = I + 1
			prioAdd += 5
		} else if T == 5 {
			prioAdd -= 5
		}

		if (T == 0 || T == 1 || T == 2) && I != allowMinus { // Если символ - оператор:
			op := newOperator(S, len(ops), T+prioAdd)
			ops = append(ops, op)

		} else if T == 3 || (T == 0 && I == allowMinus) { // Если символ - цифра:
			num += S

			if I == len(expression)-1 {
				numFloat64, err := strconv.ParseFloat(num, 64)
				if err != nil {
					return nil, nil, ErrIncorrectExpression
				}

				nums = append(nums, numFloat64)

			} else if t := tokenize(string(expression[I+1])); (t != 3) && t != T {
				numFloat64, err := strconv.ParseFloat(num, 64)
				if err != nil {
					return nil, nil, ErrIncorrectExpression
				}

				nums = append(nums, numFloat64)
				num = ""
			}

		} else if T == 6 { // Если символ - неизвестный:
			return nil, nil, ErrIncorrectExpression
		}
	}

	if (len(nums) != len(ops)+1) || (prioAdd != 0) {
		return nil, nil, ErrIncorrectExpression
	}

	return nums, ops, nil
}
