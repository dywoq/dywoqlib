package sliceutil

import "github.com/dywoq/dywoqlib/iterator"

// Find searches for a specific element in a forward-iterating collection.
// It returns the found element and nil if successful, or the zero value of T and ErrNotFound if not found.
func Find[T comparable](req T, it iterator.Forward[T]) (T, error) {
	var val T
	for it.Next() {
		val = it.Value()
		if val == req {
			return val, nil
		}
	}
	if it.Error() != nil {
		return val, it.Error()
	}
	return val, ErrNotFound
}
