package container

import (
	"errors"
	"strings"
)

func formatError(message string) error {
	strs := []string{"[github.com/dywoq/dywoqlib/container] ", message}
	formattedMessage := strings.Join(strs, "")
	result := errors.New(formattedMessage)
	return result
}

var ErrOffTheInitialLength = formatError("the actual size of data is out of the initial length")

var ErrNegativeInitialLength = formatError("the initial length can't be negative")

var ErrEmptyFixedSlice = formatError("cannot access element from an empty fixed-length slice")

var ErrFixedSliceFull = formatError("fixed-length slice is full")

var ErrInvalidIndex = formatError("index out of bounds")
