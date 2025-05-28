package container

import "errors"

var ErrOffTheInitialLength = errors.New("the actual size of data is out of the initial length")
