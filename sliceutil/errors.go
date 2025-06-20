package sliceutil

import "errors"

var ErrNotFound = errors.New("github.com/dywoq/dywoqlib/sliceutil: element wasn't found")
var ErrSliceIsNil = errors.New("github.com/dywoq/dywoqlib/sliceutil: slice is nil")
var ErrIndexOutOfBounds = errors.New("github.com/dywoq/dywoqlib/sliceutil: index out of bounds")
