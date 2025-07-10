package sliceutil

import "errors"

var (
	ErrElementNotFound = errors.New("github.com/dywoq/dywoqlib/container/traits: element wasn't found")
	ErrWrongIndex      = errors.New("github.com/dywoq/dywoqlib/container/traits: wrong index")
)
