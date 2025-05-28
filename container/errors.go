package container

import "errors"

var ErrOffTheInitialLength = errors.New("the actual size of data is out of the initial length")

var ErrNegativeInitialLength = errors.New("the initial length can't be negative")

var ErrEmptyFixedSlice = errors.New("cannot access element from an empty fixed-length slice")

var ErrFixedSliceFull = errors.New("fixed-length slice is full")

var ErrInvalidIndex = errors.New("index out of bounds")
