// Copyright 2025 dywoq
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slice

import (
	"slices"

	"github.com/dywoq/dywoqlib/iterator"
	"github.com/dywoq/dywoqlib/sliceutil"
)

// Fixed represents a slice with a predefined maximum size.
type Fixed[T comparable] struct {
	s         []T
	err       error
	fixedSize int
}

// NewFixed creates a new Fixed slice.
// It initializes the slice with given arguments and sets an error if the size constraints are violated.
func NewFixed[T comparable](size int, args ...T) *Fixed[T] {
	var err error
	if size < 0 {
		err = ErrNegativeFixedSize
	}
	if size < len(args) {
		err = ErrOverFixedSize
	}
	return &Fixed[T]{args, err, size}
}

// Length returns the current number of elements in the Fixed slice.
// It provides the actual count of stored elements.
func (f *Fixed[T]) Length() int {
	return len(f.s)
}

// Empty returns true if the fixed-size slice contains no elements.
func (f *Fixed[T]) Empty() bool {
	return len(f.s) == 0
}

// FixedSize returns the maximum capacity of the Fixed slice.
// It indicates the hard limit for the slice's size.
func (f *Fixed[T]) FixedSize() int {
	return f.fixedSize
}

// Err returns the first error encountered during operations.
// It allows checking the status of the Fixed slice.
func (f *Fixed[T]) Err() error {
	return f.err
}

// Begin returns an iterator pointing to the first element.
// It creates a new iterator starting from index 0.
func (f *Fixed[T]) Begin() *iterator.Iterator[T] {
	return iterator.New(-1, f.s)
}

// End returns an iterator pointing to the last element.
// It creates a new iterator for the final element.
func (f *Fixed[T]) End() *iterator.Iterator[T] {
	return iterator.New(len(f.s)-2, f.s)
}

// Find searches for a requested element within the Fixed slice.
// It returns the found element or a zero value if not found or an error occurs.
func (f *Fixed[T]) Find(reqElem T) T {
	if f.err != nil {
		var zero T
		return zero
	}
	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		var zero T
		return zero
	}
	m := sliceutil.Management[T]{}
	m.SetIterableType(f)
	foundElem := m.Find(reqElem)
	if m.Err() != nil {
		f.err = m.Err()
		var zero T
		return zero
	}
	return foundElem
}

// At returns the element at the specified index.
// It returns a zero value if an error occurs or the index is out of bounds.
func (f *Fixed[T]) At(i int) T {
	if f.err != nil {
		var zero T
		return zero
	}
	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		var zero T
		return zero
	}
	m := sliceutil.Management[T]{}
	m.SetIterableType(f)
	foundElem := m.At(i)
	if m.Err() != nil {
		f.err = m.Err()
		var zero T
		return zero
	}
	return foundElem
}

// String provides a string representation of the Fixed slice.
// It formats the slice's contents for display.
func (f *Fixed[T]) String() string {
	m := sliceutil.Management[T]{}
	m.SetIterableType(f)
	return m.Format()
}

// Front returns the first element of the Fixed slice.
// It returns a zero value if the slice is empty or an error occurs.
func (f *Fixed[T]) Front() T {
	if f.err != nil {
		var zero T
		return zero
	}
	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		var zero T
		return zero
	}
	if len(f.s) == 0 {
		f.err = ErrEmpty
		var zero T
		return zero
	}
	return f.At(0)
}

// Back returns the last element of the Fixed slice.
// It returns a zero value if the slice is empty or an error occurs.
func (f *Fixed[T]) Back() T {
	if f.err != nil {
		var zero T
		return zero
	}
	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		var zero T
		return zero
	}
	if len(f.s) == 0 {
		f.err = ErrEmpty
		var zero T
		return zero
	}
	return f.At(f.Length() - 1)
}

// AppendBack appends elements to the end of the Fixed slice.
// It returns the appended elements or a zero slice if an error occurs.
func (f *Fixed[T]) AppendBack(args ...T) []T {
	if f.err != nil {
		var zero []T
		return zero
	}
	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		var zero []T
		return zero
	}
	f.s = append(f.s, args...)
	return args
}

// Append adds elements to the end of the Fixed slice.
// It does not return any value and handles capacity checks internally.
func (f *Fixed[T]) Append(args ...T) {
	if f.err != nil {
		return
	}

	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		return
	}
	_ = f.AppendBack(args...)
}

// PopBack removes and returns the last element of the Fixed slice.
// It returns a zero value if the slice is empty or an error occurs.
func (f *Fixed[T]) PopBack() T {
	if f.err != nil {
		var zero T
		return zero
	}
	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		var zero T
		return zero
	}
	if len(f.s) == 0 {
		f.err = ErrEmpty
		var zero T
		return zero
	}
	i := len(f.s) - 1
	elem := f.s[i]
	f.s = slices.Delete(f.s, i, len(f.s))
	return elem
}

// Pop removes the last element of the Fixed slice.
// It handles potential errors and capacity issues.
func (f *Fixed[T]) Pop() {
	if f.err != nil {
		return
	}
	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		return
	}
	_ = f.PopBack()
}

// Erase clears all elements from the Fixed slice.
// It resets the slice to be empty while handling errors.
func (f *Fixed[T]) Erase() {
	if f.err != nil {
		return
	}
	if f.overFixedSize() {
		f.err = ErrOverFixedSize
		return
	}
	if len(f.s) == 0 {
		f.err = ErrEmpty
		return
	}
	f.s = []T{}
}

// overFixedSize checks if the current length of the slice exceeds its fixed size.
// It returns true if the slice is larger than its defined fixed size.
func (f *Fixed[T]) overFixedSize() bool {
	return len(f.s) > f.fixedSize
}
