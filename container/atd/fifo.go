package atd

import "github.com/dywoq/dywoqlib/container/slice"

type Fifo[T comparable] struct {
	err error
	d   *slice.Dynamic[T]
}

func NewFifo[T comparable]() *Fifo[T] {
	d := slice.NewDynamic[T]()
	if d.Error() != nil {
		return &Fifo[T]{d.Error(), nil}
	}
	return &Fifo[T]{nil, d}
}

func (f *Fifo[T]) Native() []T {
	return f.d.Native()
}

func (f *Fifo[T]) Error() error {
	return f.err
}

func (f *Fifo[T]) Empty() bool {
	return f.d.Length() == 0
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

func (f *Fifo[T]) String() string {
	if f.err != nil {
		return ""
	}
	return f.d.String()
}

func (f *Fifo[T]) zero() T {
	var zero T
	return zero
}
