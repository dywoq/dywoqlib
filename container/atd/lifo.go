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

import "github.com/dywoq/dywoqlib/container/slice"

// Lifo represents a generic Last-In-First-Out (LIFO) stack data structure.
// It stores elements of any comparable type T and provides stack operations.
type Lifo[T comparable] struct {
	err error
	d   *slice.Dynamic[T]
}

// NewLifo creates and returns a new instance of Lifo[T], a last-in-first-out (LIFO) stack.
// It initializes the underlying dynamic slice for storing elements of type T.
// If an error occurs during initialization, the returned Lifo will contain the error and a nil data slice.
func NewLifo[T comparable]() *Lifo[T] {
	d := slice.NewDynamic[T]()
	if d.Error() != nil {
		return &Lifo[T]{d.Error(), nil}
	}
	return &Lifo[T]{nil, d}
}

// Native returns the underlying slice of elements stored in the Lifo.
// This provides direct access to the internal data structure.
func (l *Lifo[T]) Native() []T {
	return l.d.Native()
}

// Error returns the last error encountered by the Lifo instance.
// If no error has occurred, it returns nil.
func (l *Lifo[T]) Error() error {
	return l.err
}

// Length returns the number of elements currently stored in the Lifo stack.
func (l *Lifo[T]) Length() int {
	return l.d.Length()
}

// Empty returns true if the Lifo stack contains no elements.
func (l *Lifo[T]) Empty() bool {
	return l.d.Length() == 0
}

// Append adds the given element to the Lifo container.
// If the Lifo is in an error state, it returns the zero value of T.
// If an error occurs during the append operation, the error is stored in the Lifo
// and the zero value of T is returned.
// Otherwise, it returns the appended element.
func (l *Lifo[T]) Append(elem T) T {
	if l.err != nil {
		return l.zero()
	}
	res := l.d.Append(elem)
	if l.d.Error() != nil {
		l.err = l.d.Error()
		return l.zero()
	}
	return res[0]
}

// Pop removes and returns the top element from the Lifo stack.
// If an error has previously occurred or an error occurs during the pop operation,
// it returns the zero value of type T and sets the error state.
func (l *Lifo[T]) Pop() T {
	if l.err != nil {
		return l.zero()
	}
	res := l.d.Pop()
	if l.d.Error() != nil {
		l.err = l.d.Error()
		return l.zero()
	}
	return res
}

// Top returns the element at the top of the LIFO stack without removing it.
// If the LIFO is in an error state or an error occurs while accessing the top element,
// it returns the zero value of type T.
func (l *Lifo[T]) Top() T {
	if l.err != nil {
		return l.zero()
	}
	res := l.d.Back()
	if l.d.Error() != nil {
		l.err = l.d.Error()
		return l.zero()
	}
	return res
}

// String returns the string representation of the Lifo stack.
// If the Lifo has an error, it returns an empty string.
func (l *Lifo[T]) String() string {
	if l.err != nil {
		return ""
	}
	return l.d.String()
}

func (l *Lifo[T]) zero() T {
	var zero T
	return zero
}
