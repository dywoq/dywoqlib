package slice

import (
	"github.com/dywoq/dywoqlib/iterator"
	"github.com/dywoq/dywoqlib/sliceutil"
)

type Dynamic[T comparable] struct {
	err error
	s   []T
}

func NewDynamic[T comparable](elems ...T) *Dynamic[T] {
	return &Dynamic[T]{nil, elems}
}

func (d *Dynamic[T]) Grow(i int) {
	if d.err != nil {
		return
	}
	if cap(d.s) < i {
		newSlice := make([]T, len(d.s), i)
		copy(newSlice, d.s)
		d.s = newSlice
	}
}

func (d *Dynamic[T]) Native() []T {
	return d.s
}

func (d *Dynamic[T]) Error() error {
	return d.err
}

func (d *Dynamic[T]) Length() int {
	return len(d.s)
}

func (d *Dynamic[T]) Iterating() *iterator.Combined[T] {
	return iterator.NewCombined(d.s)
}

func (d *Dynamic[T]) Append(elems ...T) []T {
	if d.err != nil {
		return []T{}
	}
	d.s = append(d.s, elems...)
	return elems
}

func (d *Dynamic[T]) At(i int) T {
	if d.err != nil {
		return d.zero()
	}
	found, err := sliceutil.At(i, d.s)
	if err != nil {
		d.err = err
		return d.zero()
	}
	return found
}

func (d *Dynamic[T]) Find(req T) T {
	if d.err != nil {
		return d.zero()
	}
	found, err := sliceutil.Find(req, d.Iterating().Forward())
	if err != nil {
		return d.zero()
	}
	return found
}

func (d *Dynamic[T]) String() string {
	if d.err != nil {
		return ""
	}
	formatted, err := sliceutil.Format(d.s)
	if err != nil {
		d.err = err
		return ""
	}
	return formatted
}

func (d *Dynamic[T]) Set(elem T, i int) T {
	if d.err != nil {
		return d.zero()
	}
	new, err := sliceutil.Set(elem, i, d.s)
	if err != nil {
		d.err = err
		return d.zero()
	}
	return new
}

func (d *Dynamic[T]) Delete(i int) T {
	if d.err != nil {
		return d.zero()
	}
	deleted, err := sliceutil.Delete(i, d.s)
	if err != nil {
		d.err = err
		return d.zero()
	}
	return deleted
}

func (d *Dynamic[T]) Insert(i int, elem T) T {
	if d.err != nil {
		return d.zero()
	}
	inserted, err := sliceutil.Insert(i, elem, d.s)
	if err != nil {
		d.err = err
		return d.zero()
	}
	return inserted
}

func (d *Dynamic[T]) zero() T {
	var zero T
	return zero
}
