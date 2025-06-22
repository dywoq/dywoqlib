package atd

import "github.com/dywoq/dywoqlib/container/slice"

// Fifo is a generic first-in-first-out queue for comparable types.
type Fifo[T comparable] struct {
	data *slice.Dynamic[T]
	err  error
}

// NewFifo creates and returns a new Fifo.
func NewFifo[T comparable]() *Fifo[T] {
	return &Fifo[T]{slice.NewDynamic[T](), nil}
}

// Err returns the last error encountered by the Fifo.
func (f *Fifo[T]) Err() error {
	return f.err
}

// Length returns the number of elements in the Fifo.
func (f *Fifo[T]) Length() int {
	return f.data.Length()
}

// Empty reports whether the Fifo is empty.
// If an error occurs, it sets the error and returns false.
func (f *Fifo[T]) Empty() bool {
	if f.err != nil {
		return false
	}
	if f.data == nil {
		f.err = ErrNilData
		return false
	}
	empty := f.data.Empty()
	if f.data.Err() != nil {
		f.err = f.data.Err()
		return false
	}
	return empty
}

// Front returns the first element in the Fifo without removing it.
// If an error occurs, it sets the error and returns the zero value.
func (f *Fifo[T]) Front() T {
	if f.err != nil {
		var zero T
		return zero
	}
	front := f.data.Front()
	if f.data.Err() != nil {
		f.err = f.data.Err()
		var zero T
		return zero
	}
	return front
}

// Back returns the last element in the Fifo without removing it.
// If an error occurs, it sets the error and returns the zero value.
func (f *Fifo[T]) Back() T {
	if f.err != nil {
		var zero T
		return zero
	}
	back := f.data.Back()
	if f.data.Err() != nil {
		f.err = f.data.Err()
		var zero T
		return zero
	}
	return back
}

// Append adds one or more elements to the end of the Fifo.
// If an error occurs, it sets the error.
func (f *Fifo[T]) Append(elems ...T) {
	if f.err != nil {
		return
	}
	if f.data == nil {
		f.err = ErrNilData
		return
	}
	f.data.Append(elems...)
	if f.data.Err() != nil {
		f.err = f.data.Err()
		return
	}
}

// Pop removes the first element from the Fifo.
// If an error occurs, it sets the error.
func (f *Fifo[T]) Pop() {
	if f.err != nil {
		return
	}
	if f.data == nil {
		f.err = ErrNilData
		return
	}
	f.data.Pop()
	if f.data.Err() != nil {
		f.err = f.data.Err()
		return
	}
}
