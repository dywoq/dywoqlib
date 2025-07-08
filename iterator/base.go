package iterator

// Base defines a generic iterator interface for traversing a collection of elements of type T.
type Base[T comparable] interface {
	// Error returns any error encountered during iteration.
	Error() error
	// Position returns the current position.
	Position() int
	// Value returns the current element of type T.
	Value() T
	// Next advances the iterator to the next element and returns true if there is a next element.
	Next() bool
	// Reset resets the iterator to its initial state.
	Reset()
	// Length returns the current length ofthe slice.
	Length() int
}
