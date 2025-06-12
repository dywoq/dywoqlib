package iterator

// New returns a new instance of the structure Iterator.
// The generic type must be given.
func New[T any](slice []T) Iterator[T] {
	return Iterator[T]{pos: 0, data: slice, err: nil}
}
