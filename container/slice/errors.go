package slice

import "errors"

var ErrEmpty = errors.New("github.com/dywoq/dywoqlib/container/slice: slice is empty")
var ErrOverFixedSize = errors.New("github.com/dywoq/dywoqlib/container/slice: over fixed size")
var ErrNegativeFixedSize = errors.New("github.com/dywoq/dywoqlib/container/slice: negative fixed size")
