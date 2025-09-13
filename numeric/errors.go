package numeric

import "errors"

var (
	ErrDivisionByZero  = errors.New("division by zero")
	ErrNegativeShift = errors.New("negative shift")
)
