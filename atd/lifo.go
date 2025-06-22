package atd

import (
	"github.com/dywoq/dywoqlib/container/slice"
)

// Lifo is a generic last-in-first-out stack for comparable types.
type Lifo[T comparable] struct {
	data *slice.Dynamic[T]
	err  error
}

// NewLifo creates and returns a new Lifo.
func NewLifo[T comparable]() *Lifo[T] {
	return &Lifo[T]{slice.NewDynamic[T](), nil}
}

// Err returns the last error encountered by the Lifo.
func (l *Lifo[T]) Err() error {
	return l.err
}

// Length returns the number of elements in the Lifo.
func (l *Lifo[T]) Length() int {
	return l.data.Length()
}

// Empty reports whether the Lifo is empty.
// If an error occurs, it sets the error and returns false.
func (l *Lifo[T]) Empty() bool {
	if l.err != nil {
		return false
	}
	if l.data == nil {
		l.err = ErrNilData
		return false
	}
	empty := l.data.Empty()
	if l.data.Err() != nil {
		l.err = l.data.Err()
		return false
	}
	return empty
}

// Top returns the last element added to the Lifo without removing it.
// If an error occurs, it sets the error and returns the zero value.
func (l *Lifo[T]) Top() T {
	if l.err != nil {
		var zero T
		return zero
	}
	back := l.data.Back()
	if l.data.Err() != nil {
		l.err = l.data.Err()
		var zero T
		return zero
	}
	return back
}

// Append adds an element to the top of the Lifo.
// If an error occurs, it sets the error.
func (l *Lifo[T]) Append(elem T) {
	if l.err != nil {
		return
	}
	if l.data == nil {
		l.err = ErrNilData
		return
	}
	l.data.Append(elem)
	if l.data.Err() != nil {
		l.err = l.data.Err()
		return
	}
}

// Pop removes the last element from the Lifo.
// If an error occurs, it sets the error.
func (l *Lifo[T]) Pop() {
	if l.err != nil {
		return
	}
	if l.data == nil {
		l.err = ErrNilData
		return
	}
	l.data.Pop()
	if l.data.Err() != nil {
		l.err = l.data.Err()
		return
	}
}

// String returns a string representation of the Lifo.
func (l *Lifo[T]) String() string {
	return l.data.String()
}
