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

// ErrOffTheInitialLength indicates that the actual size is more than the initial length.
// This is a critical error and it's recommended to generate panic if the condition is violated.
var ErrOffTheInitialLength = formatError("the actual size of data is out of the initial length")

// ErrNegativeInitialLength means the initial length can't be negative.
// Since this is a critical error, it's recommended to generate panic if the condition is violated.
var ErrNegativeInitialLength = formatError("the initial length can't be negative")

// ErrEmptySlice is a error which can be returned if a slice is empty.
var ErrEmptySlice = formatError("cannot access element from an empty fixed-length slice")

// ErrSliceFull is returned when the slice is full.
var ErrSliceFull = formatError("slice is full")

// ErrInvalidIndex indicates that the index out of bounds.
var ErrInvalidIndex = formatError("index out of bounds")
