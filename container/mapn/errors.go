package mapn

import "errors"

// ErrEmpty indicates that the map is empty.
// This error occurs when an operation expects elements but none are present.
var ErrEmpty = errors.New("github.com/dywoq/dywoqlib/container/mapn: map is empty")

// ErrNotFound indicates that the requested key was not found in the map.
// This error occurs during lookup operations when the key does not exist.
var ErrNotFound = errors.New("github.com/dywoq/dywoqlib/container/mapn: not found")