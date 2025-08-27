package unique

import (
	"slices"

	"github.com/dywoq/dywoqlib/err"
	"github.com/dywoq/dywoqlib/iterator"
	"github.com/dywoq/dywoqlib/sliceutil"
)

type Slice[T comparable] struct {
	s   []T
	err err.Context
}

func NewSlice[T comparable](elems ...T) *Slice[T] {
	s := &Slice[T]{}
	result := []T{}
	for i, elem := range elems {
		// to prevent out of bounds runtime error
		if i == 0 {
			result = append(result, elem)
			continue
		}
		previous := elems[i-1]
		if previous != elem {
			result = append(result, elem)
		}
	}
	s.s = result
	s.err = err.NoneContext()
	return s
}

func (s Slice[T]) Grow(i int) {
	if !s.err.Nil() {
		return
	}
	if cap(s.s) < i {
		newSlice := make([]T, len(s.s), i)
		copy(newSlice, s.s)
		copy(s.s, newSlice)
	}
}

func (s Slice[T]) Native() []T {
	return s.s
}

func (s Slice[T]) Error() err.Context {
	return s.err
}

func (s Slice[T]) Length() int {
	return len(s.s)
}

func (s Slice[T]) Iterating() *iterator.Combined[T] {
	return iterator.NewCombined(s.s)
}

func (s Slice[T]) Append(elems ...T) []T {
	if !s.err.Nil() {
		return []T{}
	}
	appended := []T{}
	for _, elem := range elems {
		if slices.Contains(s.s, elem) {
			continue
		}
		s.s = append(s.s, elem)
		appended = append(appended, elem)
	}
	return appended
}

func (s Slice[T]) At(i int) T {
	if !s.err.Nil() {
		return s.zero()
	}
	found, err2 := sliceutil.At(i, s.s)
	if err2 != nil {
		s.err.SetError(err2)
		s.err.SetMore("source is \"unique.Slice[T].At(int) T\"")
		return s.zero()
	}
	return found
}

func (s Slice[T]) Find(req T) T {
	if !s.err.Nil() {
		return s.zero()
	}
	found, err2 := sliceutil.Find(req, s.Iterating().Forward())
	if err2 != nil {
		s.err.SetError(err2)
		s.err.SetMore("source is \"unique.Slice[T].Find(T) T\"")
		return s.zero()
	}
	return found
}

func (s Slice[T]) String() string {
	if !s.err.Nil() {
		return ""
	}
	formatted, err2 := sliceutil.Format(s.s)
	if err2 != nil {
		s.err.SetError(err2)
		s.err.SetMore("source is \"unique.Slice[T].String() string\"")
		return ""
	}
	return formatted
}

func (s Slice[T]) Set(elem T, i int) T {
	if !s.err.Nil() {
		return s.zero()
	}

	if slices.Contains(s.s, elem) {
		return s.zero()
	}

	new, err2 := sliceutil.Set(elem, i, s.s)
	if err2 != nil {
		s.err.SetError(err2)
		s.err.SetMore("source is \"unique.Slice[T].Set(T, int) T\"")
		return s.zero()
	}
	return new
}

func (s Slice[T]) Delete(i int) T {
	if !s.err.Nil() {
		return s.zero()
	}
	deleted, err2 := sliceutil.Delete(i, s.s)
	if err2 != nil {
		s.err.SetError(err2)
		s.err.SetMore("source is \"unique.Slice[T].Delete(int) T\"")
		return s.zero()
	}
	return deleted
}

func (s Slice[T]) Insert(i int, elem T) T {
	if !s.err.Nil() {
		return s.zero()
	}

	if slices.Contains(s.s, elem) {
		return s.zero()
	}

	inserted, err2 := sliceutil.Insert(i, &s.s, elem)
	if err2 != nil {
		s.err.SetError(err2)
		s.err.SetMore("source is \"unique.Slice[T].Insert(int, T) T\"")
		return s.zero()
	}
	return inserted
}

func (s Slice[T]) Front() T {
	if !s.err.Nil() {
		return s.zero()
	}
	got := s.At(0)
	if !s.err.Nil() {
		return s.zero()
	}
	return got
}

func (s Slice[T]) Back() T {
	if !s.err.Nil() {
		return s.zero()
	}
	got := s.At(len(s.s) - 1)
	if !s.err.Nil() {
		return s.zero()
	}
	return got
}

func (s Slice[T]) Pop() T {
	if !s.err.Nil() {
		return s.zero()
	}
	if len(s.s) == 0 {
		return s.zero()
	}
	lastIdx := len(s.s) - 1
	poppedElem := s.s[lastIdx]
	copy(s.s, s.s[:lastIdx])
	return poppedElem
}

func (s Slice[T]) zero() T {
	var zero T
	return zero
}
