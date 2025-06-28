// Package stringn provides advanced string types with dynamic and fixed-length constraints,
// along with utility methods for manipulation and error handling.
package stringn

import "errors"

// ErrEmpty is returned when an operation is attempted on an empty string.
var ErrEmpty = errors.New("github.com/dywoq/dywoqlib/stringn: string is empty")

// ErrOutOfFixedLength is returned when an operation violates the fixed length constraint.
var ErrOutOfFixedLength = errors.New("github.com/dywoq/dywoqlib/stringn: out of fixed length")

// ErrNegativeFixedLength is returned when a negative value is provided for a fixed-length string operation.
var ErrNegativeFixedLength = errors.New("github.com/dywoq/dywoqlib/stringn: negative fixed-length")
