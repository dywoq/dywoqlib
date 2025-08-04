package optional

import (
	"fmt"
)

type implementation[T any] struct {
	value   T
	present bool
}

func (o *implementation[T]) Present() bool {
	return o.present
}

func (o *implementation[T]) Get() (T, bool) {
	if !o.present {
		var zero T
		return zero, false
	}
	return o.value, true
}

func (o *implementation[T]) String() string {
	if o.present {
		return fmt.Sprintf("%v", o.value)
	}
	var zero T
	return fmt.Sprintf("%v", zero)
}

func (o *implementation[T]) Else(other T) T {
	if o.present {
		return o.value
	}
	return other
}

func (o *implementation[T]) Filter(filter func(T) bool) Maybe[T] {
	if o.present && filter(o.value) {
		return o
	}
	return None[T]()
}
