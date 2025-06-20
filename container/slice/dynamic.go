package slice

import (
	"slices"

	"github.com/dywoq/dywoqlib/iterator"
	"github.com/dywoq/dywoqlib/sliceutil"
)

// Dynamic represents a dynamic slice with error handling.
type Dynamic[T comparable] struct {
	s   []T
	err error
}

// NewDynamic creates a new Dynamic slice.
// It initializes the slice with the provided data.
func NewDynamic[T comparable](data ...T) *Dynamic[T] {
	return &Dynamic[T]{data, nil}
}

// Length returns the number of elements in the Dynamic slice.
// It provides the current size of the underlying slice.
func (d *Dynamic[T]) Length() int {
	return len(d.s)
}

// Err returns the first error that occurred during operations on the Dynamic slice.
// It allows checking for any errors encountered.
func (d *Dynamic[T]) Err() error {
	return d.err
}

// Begin returns an iterator pointing to the first element.
// It sets an internal error if iterator creation fails.
func (d *Dynamic[T]) Begin() *iterator.Iterator[T] {
	it := iterator.New(0, d.s)
	if it.Err() != nil {
		d.err = it.Err()
		var zero *iterator.Iterator[T]
		return zero
	}
	return it
}

// End returns an iterator pointing to the last element.
// It sets an internal error if iterator creation fails.
func (d *Dynamic[T]) End() *iterator.Iterator[T] {
	it := iterator.New(len(d.s)-1, d.s)
	if it.Err() != nil {
		d.err = it.Err()
		var zero *iterator.Iterator[T]
		return zero
	}
	return it
}

// Find searches for a requested element in the Dynamic slice.
// It returns the found element or a zero value if not found or an error occurs.
func (d *Dynamic[T]) Find(reqElem T) T {
	if d.err != nil {
		var zero T
		return zero
	}
	m := sliceutil.Management[T]{}
	m.SetIterableType(d)
	foundElem := m.Find(reqElem)
	if m.Err() != nil {
		d.err = m.Err()
		var zero T
		return zero
	}
	return foundElem
}

// At returns the element at the specified index.
// It returns a zero value if an error occurred or the index is out of bounds.
func (d *Dynamic[T]) At(i int) T {
	if d.err != nil {
		var zero T
		return zero
	}
	m := sliceutil.Management[T]{}
	m.SetIterableType(d)
	foundElem := m.At(i)
	if m.Err() != nil {
		d.err = m.Err()
		var zero T
		return zero
	}
	return foundElem
}

// String provides a string representation of the Dynamic slice.
// It formats the slice using the internal management utility.
func (d *Dynamic[T]) String() string {
	m := sliceutil.Management[T]{}
	m.SetIterableType(d)
	return m.Format()
}

// Front returns the first element of the Dynamic slice.
// It returns a zero value if the slice is empty or an error occurs.
func (d *Dynamic[T]) Front() T {
	if d.err != nil {
		var zero T
		return zero
	}
	if len(d.s) == 0 {
		d.err = ErrEmpty
		var zero T
		return zero
	}
	return d.At(0)
}

// Back returns the last element of the Dynamic slice.
// It returns a zero value if the slice is empty or an error occurs.
func (d *Dynamic[T]) Back() T {
	if d.err != nil {
		var zero T
		return zero
	}
	if len(d.s) == 0 {
		d.err = ErrEmpty
		var zero T
		return zero
	}
	return d.At(d.Length() - 1)
}

// AppendBack appends elements to the end of the Dynamic slice.
// It returns the appended elements or a zero slice if an error occurs.
func (d *Dynamic[T]) AppendBack(args ...T) []T {
	if d.err != nil {
		var zero []T
		return zero
	}
	if len(d.s) == 0 {
		d.err = ErrEmpty
		var zero []T
		return zero
	}
	d.s = append(d.s, args...)
	return args
}

// Append adds elements to the end of the Dynamic slice.
// It does not return any value, but internally calls AppendBack.
func (d *Dynamic[T]) Append(args ...T) {
	if d.err != nil {
		return
	}
	_ = d.AppendBack(args...)
}

// PopBack removes and returns the last element of the Dynamic slice.
// It returns a zero value if the slice is empty or an error occurs.
func (d *Dynamic[T]) PopBack() T {
	if d.err != nil {
		var zero T
		return zero
	}
	if len(d.s) == 0 {
		d.err = ErrEmpty
		var zero T
		return zero
	}
	i := len(d.s) - 1
	elem := d.s[i]
	d.s = slices.Delete(d.s, i, len(d.s))
	return elem
}

// Pop removes the last element of the Dynamic slice.
// It does not return any value, but internally calls PopBack.
func (d *Dynamic[T]) Pop() {
	if d.err != nil {
		return
	}
	_ = d.PopBack()
}

// Erase clears all elements from the Dynamic slice.
// It sets an internal error if the slice is already empty.
func (d *Dynamic[T]) Erase() {
	if d.err != nil {
		return
	}
	if len(d.s) == 0 {
		d.err = ErrEmpty
		return
	}
	d.s = []T{}
}
