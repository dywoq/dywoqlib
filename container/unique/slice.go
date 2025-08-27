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

package unique

import (
	"slices"
	"sync"

	"github.com/dywoq/dywoqlib/err"
	"github.com/dywoq/dywoqlib/iterator"
	"github.com/dywoq/dywoqlib/sliceutil"
)

// Slice represents a dynamic slice structure,
// ensuring all elements are unique and there are no duplicates.
type Slice[T comparable] struct {
	s   []T
	err err.Context
	mu  sync.Mutex
}

// NewSlice creates and returns a new pointer to Slice.
// First it checks for any duplicates in elems.
// If previous element is same as the current one, the current one is skipped,
// otherwise it's appended to slice. After the iteration completed,
// it returns a pointer to Slice.
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
	s.mu = sync.Mutex{}
	return s
}

// Grow pre-allocates the underlying slice, unless there are no encountered errors.
// If capacity of the slice is lower than i, it creates a new slice with the initial capacity i,
// and copies the new slice to the underlying one.
// Locks the mutex and unlocks after the completing.
func (s *Slice[T]) Grow(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.err.Nil() {
		return
	}
	if cap(s.s) < i {
		newSlice := make([]T, len(s.s), i)
		copy(newSlice, s.s)
		copy(s.s, newSlice)
	}
}

// Native returns the underlying slice.
// Locks the mutex and unlocks after the completing.
func (s *Slice[T]) Native() []T {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.s
}

// Error returns the possible encountered error.
// Locks the mutex and unlocks after the completing.
func (s *Slice[T]) Error() err.Context {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.err
}

// Length returns the length of the underlying slice.
// Locks the mutex and unlocks after the completing.
func (s *Slice[T]) Length() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.s)
}

// Iterating returns a pointer to iterator.Combined structure.
// Locks the mutex and unlocks after the completing.
func (s *Slice[T]) Iterating() *iterator.Combined[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	return iterator.NewCombined(s.s)
}

// Append appends elems to the underlying slice, unless elems has elements that duplicate
// already existing elements in the underlying slice.
// If the error state is not nil, it returns the zero value and skips the operation.
// If any error has encountered during the operation, the function sets the error to the internal error state.
// Returns the appended elements.
// Locks the mutex and unlocks after the completing.
func (s *Slice[T]) Append(elems ...T) []T {
	s.mu.Lock()
	defer s.mu.Unlock()
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

// At returns the element at i, if i is not out of bounds of the underlying slice.
// If the error state is not nil, it returns the zero value and skips the operation.
// If any error has encountered during the operation, the function sets the error to the internal error state.
// Returns the element at i.
// Locks the mutex and unlocks after the completing.
func (s *Slice[T]) At(i int) T {
	s.mu.Lock()
	defer s.mu.Unlock()
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

// Find finds req in the underlying slice.
// If the error state is not nil, it returns the zero value and skips the operation.
// If any error has encountered during the operation, the function sets the error to the internal error state.
// Returns the found element.
// Locks the mutex and unlocks after the completing.
func (s *Slice[T]) Find(req T) T {
	s.mu.Lock()
	defer s.mu.Unlock()
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

// String returns the formatted underlying slice.
// If the error state is not nil, it returns the zero value and skips the operation.
// If any error has encountered during the operation, the function sets the error to the internal error state.
// Locks the mutex and unlocks after the completing.
func (s *Slice[T]) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()
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

// Set sets elem at i in the underlying slice, if elem doesn't exist in the underlying slice.
// If the error state is not nil, it returns the zero value and skips the operation.
// If any error has encountered during the operation, the function sets the error to the internal error state.
// Returns the set element.
// Locks the mutex and unlocks after the completing.
func (s *Slice[T]) Set(elem T, i int) T {
	s.mu.Lock()
	defer s.mu.Unlock()
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

// Delete deletes the element at i in the underlying slice.
// If the error state is not nil, it returns the zero value and skips the operation.
// If any error has encountered during the operation, the function sets the error to the internal error state.
// Returns the deleted element.
// Locks the mutex and unlocks after the completing.
func (s *Slice[T]) Delete(i int) T {
	s.mu.Lock()
	defer s.mu.Unlock()
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

// Insert inserts elem at i in the underlying slice, if it doesn't exist in the slice already.
// If the error state is not nil, it returns the zero value and skips the operation.
// If any error has encountered during the operation, the function sets the error to the internal error state.
// Returns the inserted element.
// Locks the mutex and unlocks after the completing.
func (s *Slice[T]) Insert(i int, elem T) T {
	s.mu.Lock()
	defer s.mu.Unlock()
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

// Front returns the front element in the underlying slice.
// If the error state is not nil, it returns the zero value and skips the operation.
// If any error has encountered during the operation, the function sets the error to the internal error state.
// Locks the mutex and unlocks after the completing.
func (s *Slice[T]) Front() T {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.err.Nil() || len(s.s) == 0 {
		return s.zero()
	}
	return s.s[0]
}

// Back returns the back element in the underlying slice.
// If the error state is not nil, it returns the zero value and skips the operation.
// If any error has encountered during the operation, the function sets the error to the internal error state.
// Locks the mutex and unlocks after the completing.
func (s *Slice[T]) Back() T {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.err.Nil() || len(s.s) == 0 {
		return s.zero()
	}
	return s.s[len(s.s)-1]
}

// Pop removes the last element in the underlying slice.
// If the error state is not nil, it returns the zero value and skips the operation.
// If any error has encountered during the operation, the function sets the error to the internal error state.
// Returns the popped element.
// Locks the mutex and unlocks after the completing.
func (s *Slice[T]) Pop() T {
	s.mu.Lock()
	defer s.mu.Unlock()
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

func (s *Slice[T]) zero() T {
	var zero T
	return zero
}
