package atd

import "github.com/dywoq/dywoqlib/container/slice"

type Fifo[T comparable] struct {
	data *slice.Dynamic[T]
	err  error
}

func NewFifo[T comparable]() *Fifo[T] {
	return &Fifo[T]{slice.NewDynamic[T](), nil}
}

func (f *Fifo[T]) Err() error {
	return f.err
}

func (f *Fifo[T]) Length() int {
	return f.data.Length()
}

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
