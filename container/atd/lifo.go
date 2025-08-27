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
	"github.com/dywoq/dywoqlib/sliceutil"
	"sync"
)

// Lifo represents a thread-safe generic Last-In-First-Out (LIFO) stack data structure.
// It stores elements of any comparable type T and provides stack operations.
type Lifo[T comparable] struct {
	err error
	s   []T
	mu  sync.Mutex
}

// NewLifo creates and returns a new instance of Lifo[T], a last-in-first-out (LIFO) stack.
// It initializes the underlying dynamic slice for storing elements of type T.
// If an error occurs during initialization, the returned Lifo will contain the error and a nil data slice.
func NewLifo[T comparable]() *Lifo[T] {
	return &Lifo[T]{nil, []T{}, sync.Mutex{}}
}

// Native returns the underlying slice of elements stored in the Lifo.
// This provides direct access to the internal data structure.
// Locks the mutex and unlocks after the completing.
func (l *Lifo[T]) Native() []T {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.s
}

// Error returns the last error encountered by the Lifo instance.
// If no error has occurred, it returns nil.
// Locks the mutex and unlocks after the completing.
func (l *Lifo[T]) Error() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.err
}

// Length returns the number of elements currently stored in the Lifo stack.
// Locks the mutex and unlocks after the completing.
func (l *Lifo[T]) Length() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return len(l.s)
}

// Empty returns true if the Lifo stack contains no elements.
// Locks the mutex and unlocks after the completing.
func (l *Lifo[T]) Empty() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return len(l.s) == 0
}

// Append adds the given element to the Lifo container.
// If the Lifo is in an error state, it returns the zero value of T.
// If an error occurs during the append operation, the error is stored in the Lifo
// and the zero value of T is returned.
// Otherwise, it returns the appended element.
// Locks the mutex and unlocks after the completing.
func (l *Lifo[T]) Append(elem T) T {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.err != nil {
		return l.zero()
	}
	l.s = append(l.s, elem)
	return elem
}

// Pop removes and returns the top element from the Lifo stack.
// If an error has previously occurred or an error occurs during the pop operation,
// it returns the zero value of type T and sets the error state.
// Locks the mutex and unlocks after the completing.
func (l *Lifo[T]) Pop() T {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.err != nil {
		return l.zero()
	}
	lastIdx := len(l.s) - 1
	poppedElem := l.s[lastIdx]
	l.s = l.s[:lastIdx]
	return poppedElem
}

// Top returns the element at the top of the LIFO stack without removing it.
// If the LIFO is in an error state or an error occurs while accessing the top element,
// it returns the zero value of type T.
// Locks the mutex and unlocks after the completing.
func (l *Lifo[T]) Top() T {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.err != nil {
		return l.zero()
	}
	if len(l.s) == 0 {
		return l.zero()
	}
	return l.s[len(l.s)-1]
}

// String returns the string representation of the Lifo stack.
// If the Lifo has an error, it returns an empty string.
// Locks the mutex and unlocks after the completing.
func (l *Lifo[T]) String() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.err != nil {
		return ""
	}
	res, err := sliceutil.Format(l.s)
	if err != nil {
		l.err = err
		return ""
	}
	return res
}

func (l *Lifo[T]) zero() T {
	var zero T
	return zero
}
