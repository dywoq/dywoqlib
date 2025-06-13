package slice

import "errors"

// ErrOverInitialLength occurs when a slice is over the initial length.
var ErrOverInitialLength = errors.New("github.com/dywoq/dywoqlib/slice: over the initial length")

// ErrNegativeInitialLength occurs when an initial length is negative.
var ErrNegativeInitialLength = errors.New("github.com/dywoq/dywoqlib/slice: initial length cannot be negative")

// ErrElementNotFound occurs when an element wasn't found in a slice.
var ErrElementNotFound = errors.New("github.com/dywoq/dywoqlib/slice: element wasn't found")

// ErrWrongIndex occurs when an element is not found at the specific index.
var ErrWrongIndex = errors.New("github.com/dywoq/dywoqlib/slice: the wrong index")
