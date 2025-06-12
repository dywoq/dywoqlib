package slice

import (
	"fmt"
	"strings"
)

type management[T any] struct{}

func (m *management[T]) formatSlice(slice []T) string {
	var b strings.Builder
	b.WriteString("[")
	for i, elem := range slice {
		b.WriteString(fmt.Sprintf("%v", elem))
		if i < len(slice)-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString("]")
	return b.String()
}
