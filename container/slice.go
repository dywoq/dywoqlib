package container

import (
	"github.com/dywoq/dywoqlib/iterator"
	"github.com/dywoq/dywoqlib/sliceutil"
)

type IterableSlice[T comparable] []T
type FormattableSlice[T comparable] []T

func (f FormattableSlice[T]) Format() (string, error)        { return sliceutil.Format(f) }
func (it IterableSlice[T]) Iterating() *iterator.Combined[T] { return iterator.NewCombined(it) }
