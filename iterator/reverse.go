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

// Reverse is a generic iterator that traverses a slice of comparable elements in reverse order.
type Reverse[T comparable] struct {
	data []T
	pos  int
	err  error
}

// NewReserve creates a new Reverse iterator for the provided slice of elements.
// T must be a comparable type.
func NewReserve[T comparable](data []T) *Reverse[T] {
	return &Reverse[T]{data, len(data), nil}
}

// Error returns the error encountered during iteration, if any.
// If no error has occurred, it returns nil.
func (r *Reverse[T]) Error() error {
	return r.err
}

// Position returns the current position of the iterator within the reversed collection.
// It indicates the index of the element that will be returned by the next call to Next().
func (r *Reverse[T]) Position() int {
	return r.pos
}

// Value returns the current element in the Reverse iterator.
// If an error has occurred or the current position is out of bounds,
// it sets the error to ErrOutOfBounds (if applicable) and returns the zero value of T.
func (r *Reverse[T]) Value() T {
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
func (r *Reverse[T]) Next() bool {
	r.pos--
	return r.pos >= 0
}

// Reset sets the iterator position to the beginning, allowing iteration to start over, 
// and sets the error state to nil.
// If an error has occurred (r.err is not nil), Reset does nothing.
func (r *Reverse[T]) Reset() {
	if r.err != nil {
		return
	}
	r.pos = 0
	r.err = nil
}

// Length returns the current length ofthe slice.
// If an error has occured (f.err is not nil), it returns 0.
func (r *Reverse[T]) Length() int {
	if r.err != nil {
		return 0
	}
	return len(r.data)
}

func (r *Reverse[T]) zero() T {
	var zero T
	return zero
}
