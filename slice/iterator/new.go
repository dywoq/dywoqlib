package iterator

// New returns a new instance of the structure Iterator.
// The generic type must be given.
func New[T comparable](pos int, slice []T) Iterator[T] {
	return Iterator[T]{pos: pos, data: slice, err: nil}
}
