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
	if !o.Present() {
		var zero T
		return zero, false
	}
	return o.value, true
}

func (o *implementation[T]) String() string {
	if o.Present() {
		return fmt.Sprintf("%v", o.value)
	}
	var zero T
	return fmt.Sprintf("%v", zero)
}
