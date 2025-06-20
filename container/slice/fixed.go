package slice

import (
	"slices"

	"github.com/dywoq/dywoqlib/iterator"
	"github.com/dywoq/dywoqlib/sliceutil"
)

type Fixed[T comparable] struct {
	s         []T
	err       error
	fixedSize int
}

func NewFixed[T comparable](size int, args ...T) *Fixed[T] {
	var err error
	if size < 0 {
		err = ErrNegativeFixedSize
	}
	if size < len(args) {
		err = ErrOverFixedSize
	}
	return &Fixed[T]{args, err, size}
}

func (f *Fixed[T]) Length() int {
	return len(f.s)
}

func (f *Fixed[T]) FixedSize() int {
	return f.fixedSize
}

func (f *Fixed[T]) Err() error {
	return f.err
}

func (f *Fixed[T]) Begin() *iterator.Iterator[T] {
	return iterator.New(0, f.s)
}

func (f *Fixed[T]) End() *iterator.Iterator[T] {
	return iterator.New(len(f.s)-1, f.s)
}

func (f *Fixed[T]) Find(reqElem T) T {
	if f.err != nil {
		var zero T
		return zero
	}
	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		var zero T
		return zero
	}
	m := sliceutil.Management[T]{}
	m.SetIterableType(f)
	foundElem := m.Find(reqElem)
	if m.Err() != nil {
		f.err = m.Err()
		var zero T
		return zero
	}
	return foundElem
}

func (f *Fixed[T]) At(i int) T {
	if f.err != nil {
		var zero T
		return zero
	}
	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		var zero T
		return zero
	}
	m := sliceutil.Management[T]{}
	m.SetIterableType(f)
	foundElem := m.At(i)
	if m.Err() != nil {
		f.err = m.Err()
		var zero T
		return zero
	}
	return foundElem
}

func (f *Fixed[T]) String() string {
	m := sliceutil.Management[T]{}
	m.SetIterableType(f)
	return m.Format()
}

func (f *Fixed[T]) Front() T {
	if f.err != nil {
		var zero T
		return zero
	}
	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		var zero T
		return zero
	}
	if (len(f.s)) == 0 {
		f.err = ErrEmpty
		var zero T
		return zero
	}
	return f.At(0)
}

func (f *Fixed[T]) Back() T {
	if f.err != nil {
		var zero T
		return zero
	}
	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		var zero T
		return zero
	}
	if (len(f.s)) == 0 {
		f.err = ErrEmpty
		var zero T
		return zero
	}
	return f.At(f.Length() - 1)
}

func (f *Fixed[T]) AppendBack(args ...T) []T {
	if f.err != nil {
		var zero []T
		return zero
	}
	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		var zero []T
		return zero
	}
	if (len(f.s)) == 0 {
		f.err = ErrEmpty
		var zero []T
		return zero
	}
	f.s = append(f.s, args...)
	return args
}

func (f *Fixed[T]) Append(args ...T) {
	if f.err != nil {
		return
	}

	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		return
	}
	_ = f.AppendBack(args...)
}

func (f *Fixed[T]) PopBack() T {
	if f.err != nil {
		var zero T
		return zero
	}
	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		var zero T
		return zero
	}
	if (len(f.s)) == 0 {
		f.err = ErrEmpty
		var zero T
		return zero
	}
	i := len(f.s) - 1
	elem := f.s[i]
	f.s = slices.Delete(f.s, i, len(f.s))
	return elem
}

func (f *Fixed[T]) Pop() {
	if f.err != nil {
		return
	}
	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		return
	}
	_ = f.PopBack()
}

func (f *Fixed[T]) Erase() {
	if f.err != nil {
		return
	}
	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		return
	}
	if (len(f.s)) == 0 {
		f.err = ErrEmpty
		return
	}
	f.s = []T{}
}

func (f *Fixed[T]) overFixedSize() bool {
	return len(f.s) > f.fixedSize
}
