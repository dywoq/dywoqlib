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
	"sync"

	"github.com/dywoq/dywoqlib/iterator"
)

// Fixed provides a generic, error-aware wrapper around a Go slice with a fixed maximum length.
// It uses a Dynamic slice internally and enforces the fixed length constraint.
type Fixed[T comparable] struct {
	err      error
	fixedLen int
	d        *Dynamic[T]
	mu       sync.Mutex
}

// NewFixed creates a new Fixed slice instance with a specified fixed length and initial elements.
// It returns an error if the fixed length is invalid or initial elements exceed the fixed length.
func NewFixed[T comparable](fixedLen int, elems ...T) *Fixed[T] {
	d := NewDynamic[T]()
	if d.Error() != nil {
		return &Fixed[T]{d.Error(), fixedLen, nil, sync.Mutex{}}
	}
	if fixedLen < 0 {
		return &Fixed[T]{ErrNegativeFixedLength, fixedLen, nil, sync.Mutex{}}
	}
	if fixedLen < len(elems) {
		return &Fixed[T]{ErrFixedLengthOutOfBounds, fixedLen, nil, sync.Mutex{}}
	}
	if len(elems) > fixedLen {
		return &Fixed[T]{ErrOutOfBounds, fixedLen, nil, sync.Mutex{}}
	}
	d.Grow(fixedLen)
	d.Append(elems...)
	return &Fixed[T]{nil, fixedLen, d, sync.Mutex{}}
}

// Native returns the underlying Go slice from the internal Dynamic slice.
// This allows direct access to the raw slice data.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[T]) Native() []T {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.d.Native()
}

// Error returns the first error encountered during operations on the Fixed slice.
// It reflects errors from the Fixed slice itself or its underlying Dynamic slice.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[T]) Error() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.err
}

// Length returns the current number of elements in the Fixed slice.
// This is the logical length, not the fixed capacity.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[T]) Length() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.d.Length()
}

// Iterating returns a Combined iterator for the elements in the Fixed slice.
// This allows for standard iteration patterns over the slice's contents.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[T]) Iterating() *iterator.Combined[T] {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.d.Iterating()
}

// Append adds new elements to the Fixed slice, if capacity allows.
// It returns the appended elements or an empty slice if an error occurs or capacity is exceeded.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[T]) Append(elems ...T) []T {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return []T{}
	}
	appended := f.d.Append(elems...)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return []T{}
	}
	return appended
}

// At returns the element at the specified index.
// It updates the internal error if the index is out of bounds.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[T]) At(i int) T {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return f.zero()
	}
	got := f.d.At(i)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return got
}

// Find searches for the first occurrence of a requested element.
// It returns the found element or a zero value if not found or an error occurs.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[T]) Find(req T) T {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return f.zero()
	}
	found := f.d.Find(req)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return found
}

// String returns a string representation of the Fixed slice.
// It delegates to the underlying Dynamic slice's String method.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[T]) String() string {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return ""
	}
	formatted := f.d.String()
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return ""
	}
	return formatted
}

// Set updates the element at a given index within the fixed bounds.
// It returns the updated element or a zero value on error.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[T]) Set(elem T, i int) T {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return f.zero()
	}
	new := f.d.Set(elem, i)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return new
}

// Delete removes the element at the specified index.
// It returns the deleted element or a zero value on error.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[T]) Delete(i int) T {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return f.zero()
	}
	deleted := f.d.Delete(i)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return deleted
}

// Insert adds an element at a specific index, if it doesn't exceed the fixed length.
// It returns the inserted element or a zero value on error or if capacity is exceeded.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[T]) Insert(i int, elem T) T {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return f.zero()
	}
	inserted := f.d.Insert(i, elem)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return inserted
}

// Front returns the first element of the Fixed slice.
// It returns a zero value if the slice is empty or an error occurred.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[T]) Front() T {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return f.zero()
	}
	got := f.d.Front()
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return got
}

// Back returns the last element of the Fixed slice.
// It returns a zero value if the slice is empty or an error occurred.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[T]) Back() T {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return f.zero()
	}
	got := f.d.Back()
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return got
}

// Pop removes and returns the last element of the Fixed slice.
// It returns a zero value if the slice is empty or an error occurred.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[T]) Pop() T {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return f.zero()
	}
	got := f.d.Pop()
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return got
}

func (f *Fixed[T]) outOfBounds() bool {
	return len(f.d.s) > f.fixedLen
}

func (f *Fixed[T]) errorsOk() bool {
	if f.fixedLen < len(f.d.s) {
		f.err = ErrFixedLengthOutOfBounds
		return false
	}
	if f.err != nil {
		return false
	}
	if f.outOfBounds() {
		f.err = ErrOutOfBounds
		return false
	}
	return true
}

func (f *Fixed[T]) zero() T {
	var zero T
	return zero
}
