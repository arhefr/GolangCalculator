package parser

import (
	"calculator/internal/pkg"
	"strconv"
)

// Извлекает из математического выражения список чисел и операторов
func ParserExpression(expression string) ([]float64, []pkg.Operator, error) {
	var (
		nums []float64
		ops  []pkg.Operator

		priority   int
		allowMinus int
		num        string
	)

	expression = fixExpression(expression)
	for I, sym := range expression {
		sym := string(sym)
		token := pkg.Tokenize(sym)

		// Если оператор:
		if (token == 0 || token == 1 || token == 2) && I != allowMinus {
			allowMinus = I + 1
			ops = append(ops, pkg.NewOperator(sym, len(ops), token+priority))

			// Если число:
		} else if token == 3 || (token == 0 && I == allowMinus) {
			num += sym

			if I == len(expression)-1 {
				numFloat64, err := strconv.ParseFloat(num, 64)
				if err != nil {
					return nil, nil, pkg.ErrIncorrectFloat64
				}

				nums = append(nums, numFloat64)

			} else {
				nextToken := pkg.Tokenize(string(expression[I+1]))
				if ((nextToken != 3) && nextToken != token) || (nextToken == 4) {
					numFloat64, err := strconv.ParseFloat(num, 64)
					if err != nil {
						return nil, nil, pkg.ErrIncorrectFloat64
					}

					nums = append(nums, numFloat64)
					num = ""

				}
			}
			// "(" -> приоритет операторов увеличивается
		} else if token == 4 {
			allowMinus = I + 1
			if I != 0 {
				prevToken := pkg.Tokenize(string(expression[I-1]))
				if prevToken == 3 || prevToken == 5 {
					ops = append(ops, pkg.NewOperator("*", len(ops), 1+priority))
				}
			}
			priority += 5
			// ")" -> приоритет операторов уменьшается
		} else if token == 5 {
			priority -= 5

		} else if token == 6 {
			return nil, nil, pkg.ErrIncorrectExpression
		}
	}

	// Если количество чисел НЕ соответсвует количеству операторов ИЛИ Если скобки НЕ закрылись
	if (len(nums) != len(ops)+1) || (priority != 0) {
		return nil, nil, pkg.ErrIncorrectExpression
	}

	pkg.SortOperators(ops)
	return nums, ops, nil
}
