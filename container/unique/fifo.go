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

	"github.com/dywoq/dywoqlib/err"
)

// Fifo is a generic first-in-first-out (FIFO) queue,
// with only unique elements, using Slice internally.
type Fifo[T comparable] struct {
	s   *Slice[T]
	err err.Context
}

// NewFifo creates and returns a new pointer to Fifo structure.
func NewFifo[T comparable]() *Fifo[T] {
	return &Fifo[T]{NewSlice[T](), err.NoneContext()}
}

// Native returns the underlying slice.
func (f *Fifo[T]) Native() []T {
	return f.s.Native()
}

// Error returns the possible encountered error context.
func (f *Fifo[T]) Error() err.Context {
	return f.err
}

// Empty checks whether the length of the underlying slice is 0.
func (f *Fifo[T]) Empty() bool {
	return f.s.Length() == 0
}

// Length returns the length of the underlying slice.
func (f *Fifo[T]) Length() int {
	return f.s.Length()
}

// Front returns the front element of the slice.
// If Fifo error or the internal Slice error is not nil, it returns the zero value and sets the error.
func (f *Fifo[T]) Front() T {
	if !f.err.Nil() {
		return f.zero()
	}
	res := f.s.Front()
	if !f.s.Error().Nil() {
		f.err.SetError(f.s.Error().Error())
		f.err.SetMore("source is \"unique.Fifo[T].Front() T\"")
		return f.zero()
	}
	return res
}

// Back returns the top element of the slice.
// If Fifo error or the internal Slice error is not nil, it returns the zero value and sets the error.
func (f *Fifo[T]) Back() T {
	if !f.err.Nil() {
		return f.zero()
	}
	res := f.s.Back()
	if !f.s.Error().Nil() {
		f.err.SetError(f.s.Error().Error())
		f.err.SetMore("source is \"unique.Fifo[T].Back() T\"")
		return f.zero()
	}
	return res
}

// Front appends the element to the slice, unless it already exists in the slice.
// If Fifo error or the internal Slice error is not nil, it returns the zero value and sets the error.
func (f *Fifo[T]) Append(elem T) T {
	if !f.err.Nil() {
		return f.zero()
	}

	if slices.Contains(f.s.Native(), elem) {
		return f.zero()
	}

	res := f.s.Append(elem)
	if !f.s.Error().Nil() {
		f.err.SetError(f.s.Error().Error())
		f.err.SetMore("source is \"unique.Fifo[T].Append(T) T\"")
		return f.zero()
	}
	return res[0]
}

// Pop removes the last element of the slice.
// If Fifo error or the internal Slice error is not nil, it returns the zero value and sets the error.
func (f *Fifo[T]) Pop() T {
	if !f.err.Nil() {
		return f.zero()
	}
	res := f.s.Pop()
	if !f.s.Error().Nil() {
		f.err.SetError(f.s.Error().Error())
		f.err.SetMore("source is \"unique.Fifo[T].Pop() T\"")
		return f.zero()
	}
	return res
}

// String returns the formatted presentation of slice.
// If Fifo error or the internal Slice error is not nil, it returns the empty string and sets the error.
func (f *Fifo[T]) String() string {
	if !f.err.Nil() {
		return ""
	}
	res := f.s.String()
	if !f.s.Error().Nil() {
		f.err.SetError(f.s.Error().Error())
		f.err.SetMore("source is \"unique.Fifo[T].String() string\"")
		return ""
	}
	return res
}

func (f *Fifo[T]) zero() T {
	var zero T
	return zero
}
