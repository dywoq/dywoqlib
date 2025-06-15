package iterator

import "errors"

var (
	ErrInvalidPosition = errors.New("github.com/dywoq/dywoqlib/container/iterator: invalid position")
	ErrMapKeyNotFound = errors.New("iterator: map key not found")
)
