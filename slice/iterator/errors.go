package iterator

import "errors"

// ErrOutOfBounds indicates an attempt to access an element outside the valid
// range of the iterator's data. This error occurs when trying to retrieve a value
// at an invalid position.
var ErrOutOfBounds = errors.New("slice.iterator.Iterator: not within bounds after Next()")

// ErrNoMoreElements signifies that the iterator has reached the end of its
// data collection and no further elements are available for iteration.
var ErrNoMoreElements = errors.New("slice.iterator.Iterator: no more elements to iterate")
