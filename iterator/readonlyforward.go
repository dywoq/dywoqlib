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

// ReadonlyForward is a generic iterator for traversing a slice of elements of type T in a read-only manner.
// The underlying slice cannot be modified during iteration.
// T must be a comparable type.
type ReadonlyForward[T comparable] struct {
	data []T
	pos  int
	err  error
}

// NewReadonlyForward creates a new ReadonlyForward iterator for the provided slice of type T.
// It creates a defensive copy of the input slice to ensure immutability.
// T must be a comparable type.
func NewReadonlyForward[T comparable](data []T) *ReadonlyForward[T] {
	dataCopy := make([]T, len(data))
	copy(dataCopy, data)
	return &ReadonlyForward[T]{data: dataCopy, pos: -1, err: nil}
}

// Error returns the error encountered during iteration, or nil if no error has occurred.
func (f *ReadonlyForward[T]) Error() error {
	return f.err
}

// Position returns the current position of the iterator.
func (f *ReadonlyForward[T]) Position() int {
	return f.pos
}

// Value returns the current element of the ReadonlyForward iterator.
// If an error has occurred or the position is out of bounds, it sets the error to ErrOutOfBounds
// and returns the zero value of T.
func (f *ReadonlyForward[T]) Value() T {
	if f.err != nil {
		return f.zero()
	}
	if !(f.pos >= 0 && f.pos < len(f.data)) {
		f.err = ErrOutOfBounds
		return f.zero()
	}
	return f.data[f.pos]
}

// Next advances the iterator to the next element and returns true if there are more elements.
func (f *ReadonlyForward[T]) Next() bool {
	f.pos++
	return f.pos < len(f.data)
}

// Reset sets the iterator position to the beginning.
// If an error has occurred, Reset does nothing.
func (f *ReadonlyForward[T]) Reset() {
	if f.err != nil {
		return
	}
	f.pos = -1
}

// Length returns the length of the slice.
// If an error has occurred, it returns 0.
func (f *ReadonlyForward[T]) Length() int {
	if f.err != nil {
		return 0
	}
	return len(f.data)
}

// zero returns the zero value of type T.
func (f *ReadonlyForward[T]) zero() T {
	var zero T
	return zero
}
