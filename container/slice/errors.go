package slice

import "errors"

var (
	ErrNegativeFixedLength    = errors.New("github.com/dywoq/dywoqlib/container/slice: negative fixed length")
	ErrFixedLengthOutOfBounds = errors.New("github.com/dywoq/dywoqlib/container/slice: fixed length is out of bounds")
	ErrOutOfBounds            = errors.New("github.com/dywoq/dywoqlib/container/slice: out of bounds")
)
