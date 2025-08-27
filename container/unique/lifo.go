package unique

import (
	"slices"

	"github.com/dywoq/dywoqlib/err"
)

type Lifo[T comparable] struct {
	s   *Slice[T]
	err err.Context
}

func NewLifo[T comparable]() *Lifo[T] {
	return &Lifo[T]{NewSlice[T](), err.NoneContext()}
}

func (l *Lifo[T]) Native() []T {
	return l.s.Native()
}

func (l *Lifo[T]) Error() err.Context {
	return l.s.err
}

func (l *Lifo[T]) Length() int {
	return l.s.Length()
}

func (l *Lifo[T]) Empty() bool {
	return l.s.Length() == 0
}

func (l *Lifo[T]) Append(elem T) T {
	if !l.err.Nil() {
		return l.zero()
	}

	if slices.Contains(l.s.Native(), elem) {
		return l.zero()
	}

	res := l.s.Append(elem)
	if !l.s.Error().Nil() {
		l.err.SetError(l.s.Error().Error())
		l.err.SetMore("source is \"unique.Lifo[T].Append(T) T\"")
		return l.zero()
	}
	return res[0]
}

func (l *Lifo[T]) Pop() T {
	if !l.err.Nil() {
		return l.zero()
	}
	res := l.s.Pop()
	if !l.s.Error().Nil() {
		l.err.SetError(l.s.Error().Error())
		l.err.SetMore("source is \"unique.Lifo[T].Pop() T\"")
		return l.zero()
	}
	return res
}

func (l *Lifo[T]) Top() T {
	if !l.err.Nil() {
		return l.zero()
	}
	res := l.s.Back()
	if !l.s.Error().Nil() {
		l.err.SetError(l.s.Error().Error())
		l.err.SetMore("source is \"unique.Lifo[T].Top() T\"")
		return l.zero()
	}
	return res
}

func (l *Lifo[T]) String() string {
	if !l.err.Nil() {
		return ""
	}
	res := l.s.String()
	if !l.s.Error().Nil() {
		l.err.SetError(l.s.Error().Error())
		l.err.SetMore("source is \"unique.Lifo[T].String() string\"")
		return ""
	}
	return res
}

func (l *Lifo[T]) zero() T {
	var zero T
	return zero
}
