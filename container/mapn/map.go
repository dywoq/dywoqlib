package mapn

import (
	"fmt"
	"strings"
)

// Map represents a generic map with error handling.
type Map[K comparable, V comparable] struct {
	data map[K]V
	err  error
}

// Err returns the first error encountered during map operations.
// It allows checking for any errors that have occurred.
func (m *Map[K, V]) Err() error {
	return m.err
}

// Keys returns a slice of all keys in the map.
// It returns a zero slice if the map is empty or an error occurred.
func (m *Map[K, V]) Keys() []K {
	if m.err != nil {
		var zero []K
		return zero
	}
	if len(m.data) == 0 {
		m.err = ErrEmpty
		var zero []K
		return zero
	}
	keys := []K{}
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice of all values in the map.
// It returns a zero slice if the map is empty or an error occurred.
func (m *Map[K, V]) Values() []V {
	if m.err != nil {
		var zero []V
		return zero
	}
	if len(m.data) == 0 {
		m.err = ErrEmpty
		var zero []V
		return zero
	}
	values := []V{}
	for _, v := range m.data {
		values = append(values, v)
	}
	return values
}

// At returns the value associated with the given key.
// It returns a zero value if the key is not found, the map is empty, or an error occurred.
func (m *Map[K, V]) At(key K) V {
	if m.err != nil {
		var zero V
		return zero
	}
	if len(m.data) == 0 {
		m.err = ErrEmpty
		var zero V
		return zero
	}
	foundValue, ok := m.data[key]
	if !ok {
		m.err = ErrNotFound
		var zero V
		return zero
	}
	return foundValue
}

// String provides a string representation of the map.
// It formats the map's key-value pairs for display.
func (m *Map[K, V]) String() string {
	var b strings.Builder
	b.WriteString("[")
	for k, v := range m.data {
		b.WriteString(fmt.Sprintf("\n  %v: %v", k, v))
	}
	b.WriteString("\n]")
	return b.String()
}