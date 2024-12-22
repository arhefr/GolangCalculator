package pkg

import "strconv"

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
			prioAdd += 2
		} else if T == 5 {
			prioAdd -= 2
		}

		if (T == 1 || T == 2) && I != allowMinus { // Если символ - оператор:
			op := newOperator(S, len(ops), T+prioAdd)
			ops = append(ops, op)

		} else if T == 3 || (T == 1 && I == allowMinus) { // Если символ - цифра:
			num += S

			if I == len(expression)-1 {
				numFloat64, err := strconv.ParseFloat(num, 64)
				if err != nil {
					return nil, nil, ErrIncorrectExpression
				}

				nums = append(nums, numFloat64)

			} else if t := tokenize(string(expression[I+1])); (t != 0 && t != 3) && t != T {
				numFloat64, err := strconv.ParseFloat(num, 64)
				if err != nil {
					return nil, nil, ErrIncorrectExpression
				}

				nums = append(nums, numFloat64)
				num = ""
			}

		} else if T == 0 { // Если символ - неизвестный:
			return nil, nil, ErrIncorrectExpression
		}
	}

	if (len(nums) != len(ops)+1) || (prioAdd != 0) {
		return nil, nil, ErrIncorrectExpression
	}

	return nums, ops, nil
}
