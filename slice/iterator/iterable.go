package iterator

// Iterable interface for types that can provide iterators for traversal.
type Iterable[T comparable] interface {
	Begin() Iterator[T]
	End() Iterator[T]
}
