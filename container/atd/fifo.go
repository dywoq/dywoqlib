package atd

import "github.com/dywoq/dywoqlib/container/slice"

type Fifo[T comparable] struct {
	err error
	d   *slice.Dynamic[T]
}

func (f *Fifo[T]) Error() error {
	return f.err
}

func (f *Fifo[T]) Length() int {
	return f.d.Length()
}

func (f *Fifo[T]) Front() T {
	if f.err != nil {
		return f.zero()
	}
	res := f.d.Front()
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return res
}

func (f *Fifo[T]) Back() T {
	if f.err != nil {
		return f.zero()
	}
	res := f.d.Back()
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return res
}

func (f *Fifo[T]) Append(elem T) T {
	if f.err != nil {
		return f.zero()
	}
	res := f.d.Append(elem)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return res[0]
}

func (f *Fifo[T]) Pop() T {
	if f.err != nil {
		return f.zero()
	}
	res := f.d.Pop()
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return res
}

func (f *Fifo[T]) zero() T {
	var zero T
	return zero
}
