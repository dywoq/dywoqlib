package mapnutil

import (
	"fmt"
	"strings"
)

// Format returns a string presentation of the map.
// If the function encounters any error, it returns it and empty string.
func Format[K, V comparable](m map[K]V) (string, error) {
	if len(m) == 0 {
		return "", nil
	}
	var b strings.Builder
	b.WriteString("{\n")
	for key, value := range m {
		_, err := fmt.Fprintf(&b, "  %v: %v\n", key, value)
		if err != nil {
			return "", err
		}
	}
	b.WriteString("}")
	return b.String(), nil
}
