package container

import "errors"

// Fixed indicates a fixed-length slice wrapper.
// Fixed is a generic struct that wraps a Go slice (`[]T`) and associates it
// with a predetermined `initialLength`. This `initialLength` acts as a cap
// for the maximum allowed size of the underlying `data` slice.
type Fixed[T any] struct {
	initialLength int
	data          []T
}

// NewFixed returns a new instance of Fixed.
// Returns an error if initialLength is negative or if the actual length of
// data is greater than the initialLength.
func NewFixed[T any](initialLength int, data []T) (*Fixed[T], error) {
	if initialLength < 0 {
		return nil, ErrNegativeInitialLength
	}
	if len(data) > initialLength {
		return nil, ErrOffTheInitialLength
	}
	return &Fixed[T]{initialLength, data}, nil
}

// IsOverCapacity checks if the fixed-length slice has exceeded its initial length (capacity).
// This indicates an invalid state where the slice has grown beyond its intended maximum size.
func (f Fixed[T]) IsOverCapacity() bool {
	return f.ActualLength() > f.InitialLength()
}

// IsFull checks if the fixed-length slice has reached its initial length.
func (f Fixed[T]) IsFull() bool {
	return f.ActualLength() == f.InitialLength()
}

// InitialLength returns the initial length (capacity) of the fixed-length slice.
func (f Fixed[T]) InitialLength() int {
	return f.initialLength
}

// ActualLength returns the current number of elements in the fixed-length slice.
func (f Fixed[T]) ActualLength() int {
	return len(f.data)
}

// Empty checks if the fixed-length slice is empty.
func (f Fixed[T]) Empty() bool {
	return f.ActualLength() == 0
}

// Front returns the first element of the slice.
// Returns an error if the slice is empty.
func (f Fixed[T]) Front() (T, error) {
	if f.Empty() {
		var zero T
		return zero, ErrEmptyFixedSlice
	}

	if f.IsOverCapacity() {
		var zero T
		return zero, ErrOffTheInitialLength
	}
	return f.data[0], nil
}

// Back returns the last element of the slice.
// Returns an error if the slice is empty.
func (f Fixed[T]) Back() (T, error) {
	if f.Empty() {
		var zero T
		return zero, ErrEmptyFixedSlice
	}
	if f.IsOverCapacity() {
		var zero T
		return zero, ErrOffTheInitialLength
	}
	return f.data[f.ActualLength()-1], nil
}

// Add appends an element to the fixed-length slice.
// Returns an error if the slice is already full.
func (f *Fixed[T]) Add(val T) error {
	if f.IsFull() {
		return ErrFixedSliceFull
	}
	f.data = append(f.data, val)
	return nil
}

// Get returns the element at the specified index.
// Returns an error if the index is out of bounds.
func (f Fixed[T]) Get(index int) (T, error) {
	if index < 0 || index >= f.ActualLength() {
		var zero T
		return zero, ErrInvalidIndex
	}
	return f.data[index], nil
}

// Set sets the element at the specified index.
// Returns an error if the index is out of bounds.
func (f *Fixed[T]) Set(index int, val T) error {
	if index < 0 || index >= f.ActualLength() {
		return ErrInvalidIndex
	}
	f.data[index] = val
	return nil
}

// Native returns the current slice data.
func (f *Fixed[T]) Native() []T {
	return f.data
}

// Clear resets the fixed-length slice by setting its data to an empty slice.
// The initial length remains unchanged.
func (f *Fixed[T]) Clear() {
	f.data = []T{}
}

// RemoveAt removes the element at the specified index from the fixed-length slice.
// Returns an error if the index is out of bounds or if the slice is empty.
func (f *Fixed[T]) RemoveAt(index int) error {
	if f.Empty() {
		return ErrEmptyFixedSlice
	}
	if index < 0 || index >= f.ActualLength() {
		return ErrInvalidIndex
	}
	f.data = append(f.data[:index], f.data[index+1:]...)
	return nil
}

// Pop removes and returns the last element of the fixed-length slice.
// Returns an error if the slice is empty.
func (f *Fixed[T]) Pop() (T, error) {
	if f.Empty() {
		var zero T
		return zero, ErrEmptyFixedSlice
	}
	lastIndex := f.ActualLength() - 1
	val := f.data[lastIndex]
	f.data = f.data[:lastIndex]
	return val, nil
}

// Slice returns a sub-slice of the fixed-length slice within the specified range [start, end).
// Returns an error if the start or end indices are out of bounds or if start is greater than end.
func (f Fixed[T]) Slice(start, end int) ([]T, error) {
	if start < 0 || start > f.ActualLength() || end < 0 || end > f.ActualLength() {
		return nil, ErrInvalidIndex
	}
	if start > end {
		return nil, errors.New("start index cannot be greater than end index")
	}
	return f.data[start:end], nil
}

// Fill fills the fixed-length slice with a given value until it reaches its initialLength.
// If the slice is already full, it returns an error.
func (f *Fixed[T]) Fill(val T) error {
	if f.IsFull() {
		return ErrFixedSliceFull
	}
	for f.ActualLength() < f.InitialLength() {
		f.data = append(f.data, val)
	}
	return nil
}

// Filter returns a new Fixed slice containing only the elements for which the provided
// predicate function returns true. The new slice's initialLength will be that of the original.
func (f Fixed[T]) Filter(predicate func(T) bool) (*Fixed[T], error) {
	var filteredData []T
	for _, val := range f.data {
		if predicate(val) {
			filteredData = append(filteredData, val)
		}
	}
	return NewFixed(f.InitialLength(), filteredData)
}

// Replace replaces all occurrences of an old value with a new value in the fixed-length slice.
// Returns the number of replacements made.
// Note: This requires the type T to be comparable. For non-comparable types,
// a custom comparison function would be needed.
func (f *Fixed[T]) Replace(oldVal, newVal T) int {
	count := 0
	for i, val := range f.data {
		if any(val) == any(oldVal) {
			f.data[i] = newVal
			count++
		}
	}
	return count
}

// CurrentCapacity returns the total allocated capacity of the underlying slice.
// This is different from InitialLength which is your user-defined "fixed" limit.
func (f Fixed[T]) CurrentCapacity() int {
	return cap(f.data)
}
