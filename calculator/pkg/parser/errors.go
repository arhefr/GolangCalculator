package parser

import "errors"

var (
	ErrIncorrectFloat64 error = errors.New("fail with convert string to float64")
	ErrDivisionByZero   error = errors.New("division by zero")

	ErrIncorrectExpression error = errors.New("unknown syphol or incorrect expression")
)
