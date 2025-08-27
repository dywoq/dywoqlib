package unique

import (
	"slices"

	"github.com/dywoq/dywoqlib/err"
)

// Lifo is a generic last-in-last-out (LIFO) queue,
// with only unique elements, using Slice internally. 
type Lifo[T comparable] struct {
	s   *Slice[T]
	err err.Context
}

// NewLifo creates and returns a new pointer to Lifo structure.
func NewLifo[T comparable]() *Lifo[T] {
	return &Lifo[T]{NewSlice[T](), err.NoneContext()}
}

// Native returns the underlying slice.
func (l *Lifo[T]) Native() []T {
	return l.s.Native()
}

// Error returns the possible encountered error context.
func (l *Lifo[T]) Error() err.Context {
	return l.s.err
}

// Length returns the length of the underlying slice.
func (l *Lifo[T]) Length() int {
	return l.s.Length()
}

// Empty checks whether the length of the underlying slice is 0.
func (l *Lifo[T]) Empty() bool {
	return l.s.Length() == 0
}

// Append appends the element to the slice, unless it already exists.
// If Lifo error or the internal Slice error is not nil, it returns the zero value and sets the error.
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

// Pop removes the last element to the slice.
// If Lifo error or the internal Slice error is not nil, it returns the zero value and sets the error.
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

// Top returns the top element of the slice.
// If Lifo error or the internal Slice error is not nil, it returns the zero value and sets the error.
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

// String returns the formatted presentation of slice.
// If Lifo error or the internal Slice error is not nil, it returns the empty string and sets the error.
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
