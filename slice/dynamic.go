package slice

const initialLength int = 256

type Dynamic[T comparable] struct {
	data []T
	err  error
}

func NewDynamic[T comparable]() *Dynamic[T] {
	return NewDynamicWithData(make([]T, initialLength))
}

func NewDynamicWithData[T comparable](data []T) *Dynamic[T] {
	return &Dynamic[T]{data: data, err: nil}
}
