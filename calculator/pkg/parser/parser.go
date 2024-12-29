package parser

import (
	"strconv"
	"strings"
)

// Извлекает из математического выражения список чисел и операторов
func ParserExpression(expression string) ([]float64, []Operator, error) {
	var (
		nums []string
		ops  []Operator

		priorityOps int
		num         string
	)

	expression = convertExpression(expression)
	for i := 0; i < len(expression); i++ {
		sym := string(expression[i])

		switch sym {
		case "-", "+", ".", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			switch sym {
			case "+", "-":
				if num == "" {
					num += sym

					if i == len(expression)-1 {
						nums = append(nums, num)
						num = ""
					} else if i != 0 {
						ops = append(ops, NewOperator("+", len(ops), 0+priorityOps))
					}
				} else {
					nums = append(nums, num)
					num = sym
					ops = append(ops, NewOperator("+", len(ops), 0+priorityOps))
				}
			default:
				num += sym

				if i == len(expression)-1 {
					nums = append(nums, num)
					num = ""
				}
			}

		case "*", "/", "^", "(", ")":
			if num != "" {
				if num == "+" || num == "-" {
					ops = append(ops, NewOperator(num, len(ops), 0+priorityOps))
					num = ""
				} else {
					nums = append(nums, num)
					num = ""
				}
			}

			switch sym {
			case "(":
				if i != 0 {
					symPrev := string(expression[i-1])
					switch symPrev {
					case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ")":
						ops = append(ops, NewOperator("*", len(ops), 1+priorityOps))
					}
				}
				priorityOps += 5
			case ")":
				priorityOps -= 5
			case "*", "/":
				ops = append(ops, NewOperator(sym, len(ops), 1+priorityOps))
			case "^":
				ops = append(ops, NewOperator(sym, len(ops), 2+priorityOps))
			}
		default:
			return nil, nil, ErrIncorrectExpression
		}
	}

	numsFloat, err := ParseFloatSlice(nums)

	if err != nil || len(numsFloat)-1 != len(ops) || priorityOps != 0 {

		return nil, nil, ErrIncorrectExpression
	}
	SortOperators(ops)
	return numsFloat, ops, nil
}

// Приводит слайс типа string в тип float64
func ParseFloatSlice(numsString []string) ([]float64, error) {
	var numsFloat []float64

	for _, numStr := range numsString {
		numFloat, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return nil, ErrIncorrectFloat64
		}

		numsFloat = append(numsFloat, numFloat)
	}

	return numsFloat, nil
}

// Приводит математическое выражение в нужный вид для дальнейшей обработки
func convertExpression(expression string) string {
	for _, replace := range [][]string{{" ", ""}, {"--", "+"}, {"+-", "-"}, {"**", "^"}} {
		expression = strings.ReplaceAll(expression, replace[0], replace[1])
	}

	return expression
}
