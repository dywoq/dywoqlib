package mapn

import (
	"fmt"
	"strings"
)

type Map[K comparable, V comparable] struct {
	data map[K]V
	err  error
}

func (m *Map[K, V]) Err() error {
	return m.err
}

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
	for k, _ := range m.data {
		keys = append(keys, k)
	}
	return keys
}

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

func (m *Map[K, V]) String() string {
	var b strings.Builder
	b.WriteString("[")
	for k, v := range m.data {
		b.WriteString(fmt.Sprintf("\n  %v: %v", k, v))
	}
	b.WriteString("\n]")
	return b.String()
}
