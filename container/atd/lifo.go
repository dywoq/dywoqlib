package atd

import "github.com/dywoq/dywoqlib/container/slice"

type Lifo[T comparable] struct {
	err error
	d   *slice.Dynamic[T]
}

func NewLifo[T comparable]() *Lifo[T] {
	d := slice.NewDynamic[T]()
	if d.Error() != nil {
		return &Lifo[T]{d.Error(), nil}
	}
	return &Lifo[T]{nil, d}
}

func (l *Lifo[T]) Native() []T {
	return l.d.Native()
}

func (l *Lifo[T]) Error() error {
	return l.err
}

func (l *Lifo[T]) Length() int {
	return l.d.Length()
}

func (l *Lifo[T]) Empty() bool {
	return l.d.Length() == 0
}

func (l *Lifo[T]) Append(elem T) T {
	if l.err != nil {
		return l.zero()
	}
	res := l.d.Append(elem)
	if l.d.Error() != nil {
		l.err = l.d.Error()
		return l.zero()
	}
	return res[0]
}

func (l *Lifo[T]) Pop() T {
	if l.err != nil {
		return l.zero()
	}
	res := l.d.Pop()
	if l.d.Error() != nil {
		l.err = l.d.Error()
		return l.zero()
	}
	return res
}

func (l *Lifo[T]) Top() T {
	if l.err != nil {
		return l.zero()
	}
	res := l.d.Back()
	if l.d.Error() != nil {
		l.err = l.d.Error()
		return l.zero()
	}
	return res
}

func (l *Lifo[T]) String() string {
	if l.err != nil {
		return ""
	}
	return l.d.String()
}

func (l *Lifo[T]) zero() T {
	var zero T
	return zero
}
