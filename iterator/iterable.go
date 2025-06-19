package iterator

type Iterable[T any] interface {
	Begin() *Iterator[T]
	End() *Iterator[T]
}
