package pkg

// Токенизирует символ: 0 - другое | 1,2 - операторы | 3 - цифры | 4 - скобки | 5 - неизвестный символ
func tokenize(sym string) int {
	if sym == "+" || sym == "-" {
		return 0 // Операторы низкого порядка
	}
	if sym == "*" || sym == "/" {
		return 1 // Операторы высшего порядка
	}
	if sym == "^" {
		return 2
	}
	if sym == "0" || sym == "1" || sym == "2" || sym == "3" || sym == "4" || sym == "5" || sym == "6" || sym == "7" || sym == "8" || sym == "9" || sym == "." {
		return 3 // Цифры
	}

	// Скобки
	if sym == "(" {
		return 4
	} else if sym == ")" {
		return 5
	}

	return 6 // Неизвестный символ
}
