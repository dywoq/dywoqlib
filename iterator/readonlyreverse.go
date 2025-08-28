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

package iterator

import "github.com/dywoq/dywoqlib/attribute"

//ReadonlyReverse is a generic iterator for traversing a slice of elements of type T in reverse order in a read-only manner.
//
// The underlying slice cannot be modified during iteration.
// T must be a comparable type.
//
// DEPRECATED, MAY BE REMOVED IN THE FUTURE
type ReadonlyReverse[T comparable] struct {
	data []T
	pos  int
	err  error
}

// New ReadonlyReverse creates a new  ReadonlyReverse iterator for the provided slice of type T.
// It creates a defensive copy of the input slice to ensure immutability.
// T must be a comparable type.
//
// DEPRECATED, MAY BE REMOVED IN THE FUTURE
func NewReadonlyReverse[T comparable](data []T) *ReadonlyReverse[T] {
	attribute.Deprecated(nil)
	dataCopy := make([]T, len(data))
	copy(dataCopy, data)
	return &ReadonlyReverse[T]{data: dataCopy, pos: len(data), err: nil}
}

// Error returns the error encountered during iteration, or nil if no error has occurred.
//
// DEPRECATED, MAY BE REMOVED IN THE FUTURE
func (r *ReadonlyReverse[T]) Error() error {
	attribute.Deprecated(nil)
	return r.err
}

// Position returns the current position of the iterator within the reversed collection.
// It indicates the index of the element that will be returned by the next call to Value().
//
// DEPRECATED, MAY BE REMOVED IN THE FUTURE
func (r *ReadonlyReverse[T]) Position() int {
	attribute.Deprecated(nil)
	return r.pos
}

// Value returns the current element in the  ReadonlyReverse iterator.
// If an error has occurred or the position is out of bounds, it sets the error to ErrOutOfBounds
// and returns the zero value of T.
//
// DEPRECATED, MAY BE REMOVED IN THE FUTURE
func (r *ReadonlyReverse[T]) Value() T {
	attribute.Deprecated(nil)
	if r.err != nil {
		return r.zero()
	}
	if !(r.pos >= 0 && r.pos < len(r.data)) {
		r.err = ErrOutOfBounds
		return r.zero()
	}
	return r.data[r.pos]
}

// Next moves the iterator to the previous element and returns true if the new position is valid.
// It decrements the current position and checks if it is still within the bounds of the collection.
//
// DEPRECATED, MAY BE REMOVED IN THE FUTURE
func (r *ReadonlyReverse[T]) Next() bool {
	attribute.Deprecated(nil)
	r.pos--
	return r.pos >= 0
}

// Reset sets the iterator position to the beginning of reverse iteration, and the error state to nil.
// If an error has occurred, Reset does nothing.
//
// DEPRECATED, MAY BE REMOVED IN THE FUTURE
func (r *ReadonlyReverse[T]) Reset() {
	attribute.Deprecated(nil)
	if r.err != nil {
		return
	}
	r.pos = len(r.data)
	r.err = nil
}

// Length returns the length of the slice.
// If an error has occurred, it returns 0.
//
// DEPRECATED, MAY BE REMOVED IN THE FUTURE
func (r *ReadonlyReverse[T]) Length() int {
	attribute.Deprecated(nil)
	if r.err != nil {
		return 0
	}
	return len(r.data)
}

func (r *ReadonlyReverse[T]) zero() T {
	var zero T
	return zero
}
