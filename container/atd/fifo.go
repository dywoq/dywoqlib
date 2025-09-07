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

package atd

import (
	"sync"

	"github.com/dywoq/dywoqlib/sliceutil"
)

// Fifo represents a thread-safe generic first-in-first-out (FIFO) queue.
// It stores elements of any comparable type T and provides queue operations.
type Fifo[T comparable] struct {
	err error
	s   []T
	mu  sync.Mutex
}

// NewFifo creates and returns a new instance of Fifo for elements of type T.
func NewFifo[T comparable]() *Fifo[T] {
	return &Fifo[T]{nil, []T{}, sync.Mutex{}}
}

// Native returns the underlying slice of elements stored in the Fifo.
// This provides direct access to the internal data representation.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) Native() []T {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.s
}

// Error returns the last error encountered by the Fifo instance.
// If no error has occurred, it returns nil.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) Error() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.err
}

// Empty returns true if the FIFO queue contains no elements, otherwise false.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) Empty() bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.s) == 0
}

// Length returns the number of elements currently stored in the Fifo queue.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) Length() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.s)
}

// Front returns the element at the front of the FIFO queue without removing it.
// If an internal error has occurred, it returns the zero value of T.
// If an error is encountered when accessing the underlying data structure, it sets the error state and returns the zero value of T.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) Front() T {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.err != nil {
		return f.zero()
	}
	if len(f.s) == 0 {
		return f.zero()
	}
	res := f.s[0]
	return res
}

// Back returns the last element in the FIFO queue without removing it.
// If an error has previously occurred or an error is encountered while accessing
// the underlying data structure, Back returns the zero value of type T.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) Back() T {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.err != nil {
		return f.zero()
	}
	if len(f.s) == 0 {
		return f.zero()
	}
	res := f.s[len(f.s)-1]
	return res
}

// Append adds the specified element to the FIFO queue.
// If the FIFO is in an error state, it returns the zero value of T.
// If an error occurs during the append operation, the error is stored in the FIFO
// and the zero value of T is returned. Otherwise, it returns the appended element.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) Append(elem T) T {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.err != nil {
		return f.zero()
	}
	f.s = append(f.s, elem)
	return elem
}

// Pop removes and returns the front element from the FIFO queue.
// If an internal error has occurred previously, or if an error occurs during the pop operation,
// it returns the zero value of type T and sets the error state accordingly.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) Pop() T {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.err != nil {
		return f.zero()
	}
	lastIdx := len(f.s) - 1
	poppedElem := f.s[lastIdx]
	f.s = f.s[:lastIdx]
	return poppedElem
}

// String returns the string representation of the Fifo.
// If there is an error present in the Fifo, it returns an empty string.
// Locks the mutex and unlocks after the completing.
func (f *Fifo[T]) String() string {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.err != nil {
		return ""
	}
	formatted, err := sliceutil.Format(f.s)
	if err != nil {
		f.err = err
		return ""
	}
	return formatted
}

func (f *Fifo[T]) zero() T {
	var zero T
	return zero
}
