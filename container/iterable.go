package container

import "github.com/dywoq/dywoqlib/iterator"

type IterableSlice[T comparable] []T

func (it IterableSlice[T]) Iterating() *iterator.Combined[T] { return iterator.NewCombined(it) }
