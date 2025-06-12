package iterator

// Iterator provides a generic way to traverse over a collection of elements.
type Iterator[T any] struct {
	pos  int
	data []T
	err  error
}

// HasNext checks if there are more elements to iterate over.
func (it *Iterator[T]) HasNext() bool {
	return it.pos < len(it.data)
}

// Next advances the iterator to the next element and returns true if successful.
// It returns false if there are no more elements, setting the internal error.
func (it *Iterator[T]) Next() bool {
	if !(it.pos < len(it.data)) {
		it.err = ErrNoMoreElements
		return false
	}
	it.pos++
	return true
}

// Value returns the element at the current position.
// This should be called after Next() returns true. If called at an invalid
// position, it sets an error and returns the zero value of T.
func (it *Iterator[T]) Value() T {
	if !(it.pos > 0 && it.pos <= len(it.data)) {
		it.err = ErrOutOfBounds
		var zero T
		return zero
	}
	return it.data[it.pos-1] // pos is already incremented, so use pos-1
}

// Reset resets the iterator to the beginning, allowing re-iteration.
func (it *Iterator[T]) Reset() {
	it.pos = 0
}

// Err returns the first non-nil error encountered by the iterator.
func (it *Iterator[T]) Err() error {
	return it.err
}
