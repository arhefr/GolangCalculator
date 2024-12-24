package parser

import (
	"fmt"
	"strings"
)

// Приводит математическое выражение в нужный вид для дальнейшей обработки
func fixExpression(expression string) string {
	for _, replace := range [][]string{{" ", ""}, {"--", "+"}, {"+-", "-"}} {
		expression = strings.ReplaceAll(expression, replace[0], replace[1])
	}

	return expression
}

// Приводит число в нужный вид
func FixNum(num float64) string {
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
