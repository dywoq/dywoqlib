package slice

import "github.com/dywoq/dywoqlib/iterator"

type Fixed[T comparable] struct {
	err      error
	fixedLen int
	d        *Dynamic[T]
}

func NewFixed[T comparable](fixedLen int, elems ...T) *Fixed[T] {
	d := NewDynamic[T]()
	if d.Error() != nil {
		return &Fixed[T]{d.Error(), fixedLen, nil}
	}
	if fixedLen < 0 {
		return &Fixed[T]{ErrNegativeFixedLength, fixedLen, nil}
	}
	if len(elems) < fixedLen {
		return &Fixed[T]{ErrFixedLengthOutOfBounds, fixedLen, nil}
	}
	d.Grow(fixedLen)
	d.Append(elems...)
	return &Fixed[T]{nil, fixedLen, d}
}

func (f *Fixed[T]) Native() []T {
	return f.d.Native()
}

func (f *Fixed[T]) Error() error {
	return f.err
}

func (f *Fixed[T]) Length() int {
	return f.d.Length()
}

func (f *Fixed[T]) Iterating() *iterator.Combined[T] {
	return f.Iterating()
}

func (f *Fixed[T]) Append(elems ...T) []T {
	if f.d.Length() < f.fixedLen {
		f.err = ErrFixedLengthOutOfBounds
		return []T{}
	}
	if f.err != nil {
		return []T{}
	}
	appended := f.Append(elems...)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return []T{}
	}
	return appended
}

func (f *Fixed[T]) At(i int) T {
	if f.d.Length() < f.fixedLen {
		f.err = ErrFixedLengthOutOfBounds
		return f.zero()
	}
	if f.err != nil {
		return f.zero()
	}
	got := f.d.At(i)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return got
}

func (f *Fixed[T]) Find(req T) T {
	if f.d.Length() < f.fixedLen {
		f.err = ErrFixedLengthOutOfBounds
		return f.zero()
	}
	if f.err != nil {
		return f.zero()
	}
	found := f.d.Find(req)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return found
}

func (f *Fixed[T]) String() string {
	if f.err != nil {
		return ""
	}
	formatted := f.d.String()
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return ""
	}
	return formatted
}

func (f *Fixed[T]) Set(elem T, i int) T {
	if f.d.Length() < f.fixedLen {
		f.err = ErrFixedLengthOutOfBounds
		return f.zero()
	}
	if f.err != nil {
		return f.zero()
	}
	new := f.d.Set(elem, i)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return new
}

func (f *Fixed[T]) Delete(i int) T {
	if f.d.Length() < f.fixedLen {
		f.err = ErrFixedLengthOutOfBounds
		return f.zero()
	}
	if f.err != nil {
		return f.zero()
	}
	deleted := f.d.Delete(i)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return deleted
}

func (f *Fixed[T]) Insert(i int, elem T) T {
	if f.d.Length() < f.fixedLen {
		f.err = ErrFixedLengthOutOfBounds
		return f.zero()
	}
	if f.err != nil {
		return f.zero()
	}
	inserted := f.d.Insert(i, elem)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return inserted
}

func (f *Fixed[T]) zero() T {
	var zero T
	return zero
}
