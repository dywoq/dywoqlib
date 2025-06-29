package sliceutil

import (
	"fmt"
	"strings"
)

// Format converts a slice into a string representation.
// If the slice is empty, Format does nothing and returns "[]".
func Format[T any](s []T) string {
	if len(s) == 0 {
		return "[]"
	}
	var b strings.Builder
	b.WriteString("[")
	for i, elem := range s {
		b.WriteString(fmt.Sprintf("%v", elem))
		if i != len(s)-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString("]")
	return b.String()
}
