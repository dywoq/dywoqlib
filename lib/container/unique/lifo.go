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

	"github.com/dywoq/dywoqlib/lib/err"
	"github.com/dywoq/dywoqlib/lib/sliceutil"
)

// Lifo is a thread-safe generic last-in-last-out (LIFO) queue,
// with only unique elements, using Slice internally.
type Lifo[T comparable] struct {
	s   []T
	err err.Context
	mu  sync.Mutex
}

// NewLifo creates and returns a new pointer to Lifo structure.
func NewLifo[T comparable]() *Lifo[T] {
	return &Lifo[T]{[]T{}, err.NoneContext(), sync.Mutex{}}
}

// Native returns the underlying slice.
// Locks the mutex and unlocks after the completing.
func (l *Lifo[T]) Native() []T {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.s
}

// Error returns the possible encountered error context.
// Locks the mutex and unlocks after the completing.
func (l *Lifo[T]) Error() err.Context {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.err
}

// Length returns the length of the underlying slice.
// Locks the mutex and unlocks after the completing.
func (l *Lifo[T]) Length() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return len(l.s)
}

// Empty checks whether the length of the underlying slice is 0.
// Locks the mutex and unlocks after the completing.
func (l *Lifo[T]) Empty() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return len(l.s) == 0
}

// Append appends the element to the slice, unless it already exists.
// If Lifo error or the internal Slice error is not nil, it returns the zero value and sets the error.
// Locks the mutex and unlocks after the completing.
func (l *Lifo[T]) Append(elem T) T {
	l.mu.Lock()
	defer l.mu.Unlock()
	if !l.err.Nil() {
		return l.zero()
	}
	if slices.Contains(l.s, elem) {
		return l.zero()
	}
	l.s = append(l.s, elem)
	return elem
}

// Pop removes the last element to the slice.
// If Lifo error or the internal Slice error is not nil, it returns the zero value and sets the error.
// Locks the mutex and unlocks after the completing.
func (l *Lifo[T]) Pop() T {
	l.mu.Lock()
	defer l.mu.Unlock()
	if !l.err.Nil() {
		return l.zero()
	}
	res := sliceutil.Pop(&l.s)
	return res
}

// Top returns the top element of the slice.
// If Lifo error or the internal Slice error is not nil, it returns the zero value and sets the error.
// Locks the mutex and unlocks after the completing.
func (l *Lifo[T]) Top() T {
	l.mu.Lock()
	defer l.mu.Unlock()
	if !l.err.Nil() {
		return l.zero()
	}
	res := sliceutil.Back(l.s)
	return res
}

// String returns the formatted presentation of slice.
// If Lifo error or the internal Slice error is not nil, it returns the empty string and sets the error.
// Locks the mutex and unlocks after the completing.
func (l *Lifo[T]) String() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	if !l.err.Nil() {
		return ""
	}
	res, err := sliceutil.Format(l.s)
	if err != nil {
		l.err.SetError(err)
		l.err.SetMore("source is \"unique.Lifo[T].String() string\"")
		return ""
	}
	return res
}

func (l *Lifo[T]) zero() T {
	var zero T
	return zero
}
