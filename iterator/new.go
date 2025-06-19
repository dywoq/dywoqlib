package iterator

// New creates new instance of Iterator.
func New[T any](pos int, data []T) *Iterator[T] {
	return &Iterator[T]{pos, data, nil}
}
