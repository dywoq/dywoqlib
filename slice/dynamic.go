package slice

import (
	"slices"

	"github.com/dywoq/dywoqlib/slice/iterator"
)

const initialCapacity int = 256

// Dynamic is a generic dynamic slice that provides common slice operations similar to those found in other programming languages' list or vector implementations.
// It wraps a standard Go slice and includes an internal error tracker for operation failures.
type Dynamic[T comparable] struct {
	data []T
	err  error
}

// NewDynamic creates a new, empty Dynamic slice with an initial predefined capacity (256 elements).
// This pre-allocation helps to reduce reallocations when appending elements.
func NewDynamic[T comparable]() *Dynamic[T] {
	return NewDynamicWithData(make([]T, 0, initialCapacity))
}

// NewDynamicWithData creates a new Dynamic slice, initializing it with the provided slice of data.
// The capacity of the new Dynamic slice will be at least the length of the input data.
func NewDynamicWithData[T comparable](data []T) *Dynamic[T] {
	return &Dynamic[T]{data: data, err: nil}
}

// Err returns the last error that occurred during an operation on the Dynamic slice.
// If no error has occurred, it returns nil. This allows for checking the state of the slice after operations.
func (d *Dynamic[T]) Err() error {
	return d.err
}

// ActualLength returns the current number of elements stored in the Dynamic slice.
// This is equivalent to `len()` on the underlying Go slice.
func (d *Dynamic[T]) ActualLength() int {
	return len(d.data)
}

// Empty returns true if the Dynamic slice contains no elements (i.e., its actual length is zero).
func (d *Dynamic[T]) Empty() bool {
	return d.ActualLength() == 0
}

// Begin returns an Iterator pointing to the first element of the Dynamic slice.
// If an error exists on the Dynamic slice, an empty Iterator is returned.
func (d *Dynamic[T]) Begin() iterator.Iterator[T] {
	if d.err != nil {
		return iterator.Iterator[T]{}
	}
	return iterator.New(0, d.data)
}

// End returns an Iterator pointing to the last element of the Dynamic slice.
// If an error exists on the Dynamic slice, an empty Iterator is returned.
func (d *Dynamic[T]) End() iterator.Iterator[T] {
	if d.err != nil {
		return iterator.Iterator[T]{}
	}
	return iterator.New(d.ActualLength()-1, d.data)
}

// At returns the element at the specified zero-based index.
// If the index is out of bounds (less than 0 or greater than or equal to the slice's length),
// an ErrWrongIndex error is set, and a zero value of type T is returned.
func (d *Dynamic[T]) At(index int) T {
	if index < 0 || index >= len(d.data) {
		d.err = ErrWrongIndex
		var zero T
		return zero
	}
	return d.data[index]
}

// Front returns the first element of the Dynamic slice.
// If an error exists on the Dynamic slice, a zero value of type T is returned.
// This is equivalent to calling At(0).
func (d *Dynamic[T]) Front() T {
	if d.err != nil {
		var zero T
		return zero
	}
	return d.At(0)
}

// Back returns the last element of the Dynamic slice.
// If an error exists on the Dynamic slice, a zero value of type T is returned.
// This is equivalent to calling At(d.ActualLength() - 1).
func (d *Dynamic[T]) Back() T {
	if d.err != nil {
		var zero T
		return zero
	}
	return d.At(len(d.data) - 1)
}

// String returns a string representation of the Dynamic slice's contents,
// typically formatted as a list of elements.
func (d *Dynamic[T]) String() string {
	m := management[T]{}
	return m.formatSlice(d.data)
}

// Append adds one or more elements to the end of the Dynamic slice.
// The slice's capacity will grow as needed to accommodate new elements.
// This operation is skipped if an error already exists on the Dynamic slice.
func (d *Dynamic[T]) Append(elements ...T) {
	if d.err != nil {
		return
	}
	d.data = append(d.data, elements...)
}

// Remove deletes the element at the specified zero-based index from the Dynamic slice.
// If the index is out of bounds, an ErrElementNotFound error is set.
// This operation modifies the underlying slice in-place.
func (d *Dynamic[T]) Remove(index int) {
	if d.err != nil {
		return
	}
	if index < 0 || index >= len(d.data) {
		d.err = ErrElementNotFound
		return
	}
	d.data = slices.Delete(d.data, index, index+1)
}

// Contains checks if the Dynamic slice contains the given element.
// It returns true if the element is found, false otherwise.
// If the element is not found, an ErrElementNotFound error is set on the slice.
// This operation is skipped if an error already exists on the Dynamic slice.
func (d *Dynamic[T]) Contains(t T) bool {
	if d.err != nil {
		return false
	}
	if !slices.Contains(d.data, t) {
		d.err = ErrElementNotFound
		return false
	}
	return true
}

// Clear removes all elements from the Dynamic slice, effectively making it empty.
// This sets each element to its zero value of type T, but retains the underlying capacity.
// This operation is skipped if an error already exists on the Dynamic slice.
func (d *Dynamic[T]) Clear() {
	if d.err != nil {
		return
	}

	for i := range d.data {
		d.data[i] = *new(T)
	}
}
