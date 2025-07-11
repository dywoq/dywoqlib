package iterator

type Combined[T comparable] struct {
	s []T
}

func (c Combined[T]) Forward() *Forward[T] { return NewForward(c.s) }
func (c Combined[T]) Reverse() *Reverse[T] { return NewReserve(c.s) }

func NewCombined[T comparable](s []T) *Combined[T] {
	return &Combined[T]{s}
}
