package slice

import "errors"

// ErrEmpty indicates that the slice is empty.
// This error occurs when an operation expects elements but none are present.
var ErrEmpty = errors.New("github.com/dywoq/dywoqlib/container/slice: slice is empty")

// ErrOverFixedSize indicates that an operation would exceed the fixed size of the slice.
// This error applies to Fixed slices when attempts are made to add too many elements.
var ErrOverFixedSize = errors.New("github.com/dywoq/dywoqlib/container/slice: over fixed size")

// ErrNegativeFixedSize indicates that a negative value was provided for the fixed size.
// This error occurs during the initialization of a Fixed slice with an invalid size.
var ErrNegativeFixedSize = errors.New("github.com/dywoq/dywoqlib/container/slice: negative fixed size")