package container

import "errors"

var ErrElementNotFound = errors.New("github.com/dywoq/dywoqlib/container: element wasn't found")
var ErrSliceIsNil = errors.New("github.com/dywoq/dywoqlib/container: slice is nil")
var ErrIndexOutOfBounds = errors.New("github.com/dywoq/dywoqlib/container: index out of bounds")