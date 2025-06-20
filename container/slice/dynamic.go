package slice

import (
	"slices"

	"github.com/dywoq/dywoqlib/iterator"
	"github.com/dywoq/dywoqlib/sliceutil"
)

type Dynamic[T comparable] struct {
	s   []T
	err error
}

func NewDynamic[T comparable](args ...T) *Dynamic[T] {
	return &Dynamic[T]{args, nil}
}

func (d *Dynamic[T]) Length() int {
	return len(d.s)
}

func (d *Dynamic[T]) Err() error {
	return d.err
}

func (d *Dynamic[T]) Begin() *iterator.Iterator[T] {
	return iterator.New(0, d.s)
}

func (d *Dynamic[T]) End() *iterator.Iterator[T] {
	return iterator.New(len(d.s)-1, d.s)
}

func (d *Dynamic[T]) Find(reqElem T) T {
	if d.err != nil {
		var zero T
		return zero
	}
	m := sliceutil.Management[T]{}
	m.SetIterableType(d)
	foundElem := m.Find(reqElem)
	if m.Err() != nil {
		d.err = m.Err()
		var zero T
		return zero
	}
	return foundElem
}

func (d *Dynamic[T]) At(i int) T {
	if d.err != nil {
		var zero T
		return zero
	}
	m := sliceutil.Management[T]{}
	m.SetIterableType(d)
	foundElem := m.At(i)
	if m.Err() != nil {
		d.err = m.Err()
		var zero T
		return zero
	}
	return foundElem
}

func (d *Dynamic[T]) String() string {
	m := sliceutil.Management[T]{}
	m.SetIterableType(d)
	return m.Format()
}

func (d *Dynamic[T]) Front() T {
	if d.err != nil {
		var zero T
		return zero
	}
	if (len(d.s)) == 0 {
		d.err = ErrSliceIsEmpty
		var zero T
		return zero
	}
	return d.At(0)
}

func (d *Dynamic[T]) Back() T {
	if d.err != nil {
		var zero T
		return zero
	}
	if (len(d.s)) == 0 {
		d.err = ErrSliceIsEmpty
		var zero T
		return zero
	}
	return d.At(d.Length() - 1)
}

func (d *Dynamic[T]) AppendBack(args ...T) []T {
	if d.err != nil {
		var zero []T
		return zero
	}
	if (len(d.s)) == 0 {
		d.err = ErrSliceIsEmpty
		var zero []T
		return zero
	}
	d.s = append(d.s, args...)
	return args
}

func (d *Dynamic[T]) Append(args ...T) {
	if d.err != nil {
		return
	}
	_ = d.AppendBack(args...)
}

func (d *Dynamic[T]) PopBack() T {
	if d.err != nil {
		var zero T
		return zero
	}
	if (len(d.s)) == 0 {
		d.err = ErrSliceIsEmpty
		var zero T
		return zero
	}
	i := len(d.s) - 1
	elem := d.s[i]
	d.s = slices.Delete(d.s, i, len(d.s))
	return elem
}

func (d *Dynamic[T]) Pop() {
	if d.err != nil {
		return
	}
	_ = d.PopBack()
}

func (d *Dynamic[T]) Erase() {
	if d.err != nil {
		return
	}
	if (len(d.s)) == 0 {
		d.err = ErrSliceIsEmpty
		return
	}
	d.s = []T{}
}
