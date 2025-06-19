package iterator

import "errors"

var ErrNoMoreElements = errors.New("github.com/dywoq/dywoqlib/container/iterator: no more elements")
var ErrOutOfBounds = errors.New("github.com/dywoq/dywoqlib/container/iterator: out of bounds")
var ErrInvalidPosition = errors.New("github.com/dywoq/dywoqlib/container/iterator: invalid position")