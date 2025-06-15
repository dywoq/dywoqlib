package iterator

import "sort"

// Map is a generic iterator for Go maps with comparable keys and values.
// It iterates over the map's keys in a sorted order to provide deterministic iteration.
type Map[K comparable, V comparable] struct {
	data map[K]V
	keys []K
	pos  int
	err  error
}

// NewMap creates a new Map iterator for the given data.
// It initializes the iterator by extracting and sorting all keys from the map.
// The iterator's position is set to -1, indicating it's before the first element.
func NewMap[K comparable, V comparable](data map[K]V) *Map[K, V] {
	if data == nil {
		data = make(map[K]V)
	}

	keys := make([]K, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}

	if sorter, ok := any(keys).(interface {
		Len() int
		Less(i, j int) bool
		Swap(i, j int)
	}); ok {
		sort.Sort(sorter)
	}

	return &Map[K, V]{
		data: data,
		keys: keys,
		pos:  -1,
	}
}

// Err returns the first error encountered during iteration, if any.
// This allows consumers to check for errors after the iteration is complete.
func (m *Map[K, V]) Err() error {
	return m.err
}

// Position returns the current index of the iterator within the sorted keys.
// It returns -1 if the iterator is exhausted or an error has occurred.
func (m *Map[K, V]) Position() int {
	if m.err != nil || !m.isValidPosition(m.pos) {
		return -1
	}
	return m.pos
}

// Next advances the iterator to the next element in the sorted keys.
// It returns true if there is a next element, or false if the iterator is exhausted or an error occurred.
func (m *Map[K, V]) Next() bool {
	if m.err != nil {
		return false
	}

	m.pos++
	return m.isValidPosition(m.pos)
}

// Key returns the key at the current iterator position.
// If the iterator is not at a valid position, it returns the zero value of type K
// and sets an ErrInvalidPosition.
func (m *Map[K, V]) Key() K {
	if m.err != nil || !m.isValidPosition(m.pos) {
		m.zeroValueAndSetError()
		var zero K
		return zero
	}
	return m.keys[m.pos]
}

// Value returns the value associated with the current key.
// If the iterator is not at a valid position or the key is no longer in the map,
// it returns the zero value of type V and sets an appropriate error.
func (m *Map[K, V]) Value() V {
	if m.err != nil || !m.isValidPosition(m.pos) {
		m.zeroValueAndSetError()
		var zero V
		return zero
	}

	currentKey := m.keys[m.pos]
	val, ok := m.data[currentKey]
	if !ok {
		if m.err == nil {
			m.err = ErrMapKeyNotFound
		}
		var zero V
		return zero
	}
	return val
}

func (m *Map[K, V]) isValidPosition(pos int) bool {
	return pos >= 0 && pos < len(m.keys)
}

func (m *Map[K, V]) zeroValueAndSetError() {
	if m.err == nil {
		m.err = ErrInvalidPosition
	}
}
