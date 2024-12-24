package pkg

/* Возвращает токен символа:
   · 0: "+-",
   · 1: "/*",
   · 2: "^",
   · 3: "0123456789.",
   · 4: "(",
   · 5: ")",
   · 6: "?". */
func Tokenize(sym string) int {
	switch sym {

	case "+", "-":
		return 0
	case "*", "/":
		return 1
	case "^":
		return 2

	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ".":
		return 3

	case "(":
		return 4
	case ")":
		return 5

	default:
		return 6

	}
}
