package unique

import (
	"slices"

	"github.com/dywoq/dywoqlib/err"
)

type Fifo[T comparable] struct {
	s   *Slice[T]
	err err.Context
}

func NewFifo[T comparable]() *Fifo[T] {
	return &Fifo[T]{NewSlice[T](), err.NoneContext()}
}

func (f *Fifo[T]) Native() []T {
	return f.s.Native()
}

func (f *Fifo[T]) Error() err.Context {
	return f.err
}

func (f *Fifo[T]) Empty() bool {
	return f.s.Length() == 0
}

func (f *Fifo[T]) Length() int {
	return f.s.Length()
}

func (f *Fifo[T]) Front() T {
	if !f.err.Nil() {
		return f.zero()
	}
	res := f.s.Front()
	if !f.s.Error().Nil() {
		f.err.SetError(f.s.Error().Error())
		f.err.SetMore("source is \"unique.Fifo[T].Front() T\"")
		return f.zero()
	}
	return res
}

func (f *Fifo[T]) Back() T {
	if !f.err.Nil() {
		return f.zero()
	}
	res := f.s.Back()
	if !f.s.Error().Nil() {
		f.err.SetError(f.s.Error().Error())
		f.err.SetMore("source is \"unique.Fifo[T].Back() T\"")
		return f.zero()
	}
	return res
}

func (f *Fifo[T]) Append(elem T) T {
	if !f.err.Nil() {
		return f.zero()
	}

	if slices.Contains(f.s.Native(), elem) {
		return f.zero()
	}

	res := f.s.Append(elem)
	if !f.s.Error().Nil() {
		f.err.SetError(f.s.Error().Error())
		f.err.SetMore("source is \"unique.Fifo[T].Append(T) T\"")
		return f.zero()
	}
	return res[0]
}

func (f *Fifo[T]) Pop() T {
	if !f.err.Nil() {
		return f.zero()
	}
	res := f.s.Pop()
	if !f.s.Error().Nil() {
		f.err.SetError(f.s.Error().Error())
		f.err.SetMore("source is \"unique.Fifo[T].Pop() T\"")
		return f.zero()
	}
	return res
}

func (f *Fifo[T]) String() string {
	if !f.err.Nil() {
		return ""
	}
	res := f.s.String()
	if !f.s.Error().Nil() {
		f.err.SetError(f.s.Error().Error())
		f.err.SetMore("source is \"unique.Fifo[T].String() string\"")
		return ""
	}
	return res
}

func (f *Fifo[T]) zero() T {
	var zero T
	return zero
}
