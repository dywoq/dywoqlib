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
	"github.com/dywoq/dywoqlib/sliceutil"
)

// Fifo is a thread-safe generic first-in-first-out (FIFO) queue,
// with only unique elements, using Slice internally.
type Fifo[T comparable] struct {
	s   []T
	err err.Context
	mu  sync.Mutex
}

// NewFifo creates and returns a new pointer to Fifo structure.
func NewFifo[T comparable]() *Fifo[T] {
	return &Fifo[T]{[]T{}, err.NoneContext(), sync.Mutex{}}
}

// Native returns the underlying slice.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) Native() []T {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.s
}

// Error returns the possible encountered error context.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) Error() err.Context {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.err
}

// Empty checks whether the length of the underlying slice is 0.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) Empty() bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.s) == 0
}

// Length returns the length of the underlying slice.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) Length() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.s)
}

// Front returns the front element of the slice.
// If Fifo error or the internal Slice error is not nil, it returns the zero value and sets the error.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) Front() T {
	f.mu.Lock()
	defer f.mu.Unlock()
	if !f.err.Nil() {
		return f.zero()
	}
	res := sliceutil.Front(f.s)
	return res
}

// Back returns the top element of the slice.
// If Fifo error or the internal Slice error is not nil, it returns the zero value and sets the error.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) Back() T {
	f.mu.Lock()
	defer f.mu.Unlock()
	if !f.err.Nil() {
		return f.zero()
	}
	res := sliceutil.Back(f.s)
	return res
}

// Front appends the element to the slice, unless it already exists in the slice.
// If Fifo error or the internal Slice error is not nil, it returns the zero value and sets the error.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) Append(elem T) T {
	f.mu.Lock()
	defer f.mu.Unlock()
	if !f.err.Nil() {
		return f.zero()
	}
	if slices.Contains(f.s, elem) {
		return f.zero()
	}
	f.s = append(f.s, elem)
	return elem
}

// Pop removes the last element of the slice.
// If Fifo error or the internal Slice error is not nil, it returns the zero value and sets the error.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) Pop() T {
	f.mu.Lock()
	defer f.mu.Unlock()
	if !f.err.Nil() {
		return f.zero()
	}
	res := sliceutil.Pop(&f.s)
	return res
}

// String returns the formatted presentation of slice.
// If Fifo error or the internal Slice error is not nil, it returns the empty string and sets the error.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) String() string {
	f.mu.Lock()
	defer f.mu.Unlock()
	if !f.err.Nil() {
		return ""
	}
	res, err := sliceutil.Format(f.s)
	if err != nil {
		f.err.SetError(err)
		f.err.SetMore("source is \"unique.Fifo[T].String() string\"")
		return ""
	}
	return res
}

func (f *Fifo[T]) zero() T {
	var zero T
	return zero
}
