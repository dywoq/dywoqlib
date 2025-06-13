// Copyright 2025 dywoq
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slice

import (
	"slices"

	"github.com/dywoq/dywoqlib/slice/iterator"
)

// Fixed represents a slice with a predefined initial length, allowing checks for overflow and other states.
// It wraps a standard Go slice and includes an internal error tracker for operation failures.
type Fixed[T comparable] struct {
	data          []T
	initialLength int
	err           error
}

// NewFixed creates a new Fixed slice instance with a specified initial length and optional initial data.
// It returns an error if the initial data length exceeds the initialLength or if initialLength is negative.
func NewFixed[T comparable](initialLength int, data []T) *Fixed[T] {
	if len(data) > initialLength {
		return &Fixed[T]{err: ErrOverInitialLength}
	}
	if initialLength < 0 {
		return &Fixed[T]{err: ErrNegativeInitialLength}
	}
	return &Fixed[T]{data, initialLength, nil}
}

// Err returns the current error state of the Fixed slice.
// If no error has occurred, it returns nil. This allows for checking the state of the slice after operations..
func (f *Fixed[T]) Err() error {
	return f.err
}

// InitialLength returns the maximum allowed length (predefined capacity) of the Fixed slice.
// This is the length set during its creation.
func (f *Fixed[T]) InitialLength() int {
	return f.initialLength
}

// ActualLength returns the current number of elements in the Fixed slice.
// This is equivalent to `len()` on the underlying Go slice.
func (f *Fixed[T]) ActualLength() int {
	return len(f.data)
}

// Empty checks if the Fixed slice currently contains no elements (i.e., its actual length is zero).
func (f *Fixed[T]) Empty() bool {
	return f.ActualLength() == 0
}

// OverInitialLength checks if the Fixed slice's current actual length exceeds its initial (maximum allowed) length.
// It returns true if the slice has grown beyond its defined capacity.
func (f *Fixed[T]) OverInitialLength() bool {
	return f.ActualLength() > f.InitialLength()
}

// Negative checks if the Fixed slice's actual length is negative, which indicates an invalid state.
// This method serves as a safeguard against unexpected internal slice corruption.
func (f *Fixed[T]) Negative() bool {
	return f.ActualLength() < 0
}

// Begin returns an Iterator positioned at the first element of the Fixed slice.
// It first updates the internal error state based on current slice properties.
// If an error exists on the Fixed slice, an empty Iterator is returned.
func (f *Fixed[T]) Begin() iterator.Iterator[T] {
	f.updateErrorState()
	if f.err != nil {
		return iterator.Iterator[T]{}
	}
	return iterator.New(0, f.data)
}

// End returns an Iterator positioned at the last element of the Fixed slice.
// It first updates the internal error state based on current slice properties.
// If an error exists on the Fixed slice, an empty Iterator is returned.
func (f *Fixed[T]) End() iterator.Iterator[T] {
	f.updateErrorState()
	if f.err != nil {
		return iterator.Iterator[T]{}
	}
	return iterator.New(f.ActualLength()-1, f.data)
}

// At returns the element at the specified zero-based index within the Fixed slice.
// It first updates the internal error state based on current slice properties.
// If the index is out of bounds (less than 0 or greater than or equal to the slice's length),
// an ErrWrongIndex error is set, and a zero value of type T is returned.
func (f *Fixed[T]) At(index int) T {
	f.updateErrorState()
	if index < 0 || index >= len(f.data) {
		f.err = ErrWrongIndex
		var zero T
		return zero
	}
	return f.data[index]
}

// Front returns the first element of the Fixed slice.
// It first updates the internal error state based on current slice properties.
// If an error exists on the Fixed slice, a zero value of type T is returned.
// This is equivalent to calling At(0).
func (f *Fixed[T]) Front() T {
	f.updateErrorState()
	if f.err != nil {
		var zero T
		return zero
	}
	return f.At(0)
}

// Back returns the last element of the Fixed slice.
// It first updates the internal error state based on current slice properties.
// If an error exists on the Fixed slice, a zero value of type T is returned.
// This is equivalent to calling At(f.ActualLength() - 1).
func (f *Fixed[T]) Back() T {
	f.updateErrorState()
	if f.err != nil {
		var zero T
		return zero
	}
	return f.At(len(f.data) - 1)
}

// String returns a string representation of the Fixed slice's underlying data,
// typically formatted as a list of elements.
func (f *Fixed[T]) String() string {
	m := management[T]{}
	return m.formatSlice(f.data)
}

// Set updates the element at the specified zero-based index with a new value.
// It first updates the internal error state based on current slice properties.
// This operation is skipped if an error already exists on the Fixed slice.
// No bounds checking is performed within this method; ensure the index is valid externally.
func (f *Fixed[T]) Set(index int, value T) {
	f.updateErrorState()
	if f.err != nil {
		return
	}
	f.data[index] = value
}

// Append adds new elements to the end of the Fixed slice.
// It first updates the internal error state based on current slice properties.
// If adding elements would cause the slice to exceed its `InitialLength`, the operation might proceed
// but the `OverInitialLength` error will be set on the slice.
// This operation is skipped if an error already exists on the Fixed slice.
func (f *Fixed[T]) Append(elements ...T) {
	f.updateErrorState()
	if f.err != nil {
		return
	}
	f.data = append(f.data, elements...)
}

// Remove deletes the element at the specified zero-based index from the Fixed slice.
// It first updates the internal error state based on current slice properties.
// This operation modifies the underlying slice in-place.
// If the index is out of bounds, an error will be set by a preceding `updateErrorState` call or subsequent `At` call.
// This operation is skipped if an error already exists on the Fixed slice.
func (f *Fixed[T]) Remove(index int) {
	f.updateErrorState()
	if f.err != nil {
		return
	}
	f.data = slices.Delete(f.data, index, index+1)
}

// Contains checks if the Fixed slice contains a specific element.
// It first updates the internal error state based on current slice properties.
// It returns true if the element is found, false otherwise.
// This operation is skipped if an error already exists on the Fixed slice.
func (f *Fixed[T]) Contains(t T) bool {
	f.updateErrorState()
	if f.err != nil {
		return false
	}
	return slices.Contains(f.data, t)
}

// Clear sets all elements in the Fixed slice to their zero value, effectively emptying it
// while retaining the underlying slice's capacity.
// It first updates the internal error state based on current slice properties.
// This operation is skipped if an error already exists on the Fixed slice.
func (f *Fixed[T]) Clear() {
	f.updateErrorState()
	if f.err != nil {
		return
	}

	for i := range f.data {
		f.data[i] = *new(T)
	}
}

func (f *Fixed[T]) updateErrorState() {
	errs := make(map[bool]error, 3)
	errs[f.Negative()] = ErrNegativeInitialLength
	errs[f.OverInitialLength()] = ErrOverInitialLength

	for condition, err := range errs {
		if condition {
			f.err = err
			break
		}
	}
}
