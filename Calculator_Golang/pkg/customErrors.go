package pkg

import "errors"

var (
	ErrDivisionByZero      error = errors.New("error | with division by zero")
	ErrIncorrectExpression error = errors.New("error | incorrect expression like a '2+2*(' or '2++2'")
)
