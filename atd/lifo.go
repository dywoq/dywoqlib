package atd

import (
	"github.com/dywoq/dywoqlib/container/slice"
)

type Lifo[T comparable] struct {
	data *slice.Dynamic[T]
	err  error
}

func NewLifo[T comparable]() *Lifo[T] {
	return &Lifo[T]{slice.NewDynamic[T](), nil}
}

func (l *Lifo[T]) Err() error {
	return l.err
}

func (l *Lifo[T]) Length() int {
	return l.data.Length()
}

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
