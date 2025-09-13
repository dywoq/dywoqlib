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

// Forward is a generic iterator for traversing a slice of elements of type T.
// It maintains the current position within the slice and any error encountered during iteration.
// T must be a comparable type.
type Forward[T comparable] struct {
	data []T
	pos  int
	err  error
}

// NewForward creates and returns a new Forward iterator for the provided slice of type T.
// T must be a comparable type.
func NewForward[T comparable](data []T) *Forward[T] {
	return &Forward[T]{data, -1, nil}
}

// Error returns the error encountered during iteration, or nil if no error has occurred.
func (f *Forward[T]) Error() error {
	return f.err
}

// Position returns the current position of the iterator.
// It indicates the index of the element that the iterator is currently pointing to.
func (f *Forward[T]) Position() int {
	return f.pos
}

// Value returns the current element of the Forward iterator.
// If an error has occurred or the current position is out of bounds,
// it sets the error to ErrOutOfBounds (if applicable) and returns the zero value of T.
func (f *Forward[T]) Value() T {
	if f.err != nil {
		return f.zero()
	}
	if !(f.pos >= 0 && f.pos < len(f.data)) {
		f.err = ErrOutOfBounds
		return f.zero()
	}
	return f.data[f.pos]
}

// ValuePtr returns a pointer to the current element of the Forward iterator.
// If an error has occurred or the current position is out of bounds,
// it sets the error to ErrOutOfBounds (if applicable) and returns nil.
func (f *Forward[T]) ValuePtr() *T {
	if f.err != nil {
		return nil
	}
	if !(f.pos >= 0 && f.pos < len(f.data)) {
		f.err = ErrOutOfBounds
		return nil
	}
	return &f.data[f.pos]
}

// Next advances the iterator to the next element and returns true if there are more elements to iterate over.
// It increments the current position and checks if it is still within the bounds of the data slice.
func (f *Forward[T]) Next() bool {
	f.pos++
	return f.pos < len(f.data)
}

// Reset sets the iterator position to the beginning, and the error state to nil.
// If an error has occurred (f.err is not nil), Reset does nothing.
func (f *Forward[T]) Reset() {
	if f.err != nil {
		return
	}
	f.pos = 0
	f.err = nil
}

// Length returns the current length ofthe slice.
// If an error has occured (f.err is not nil), it returns 0.
func (f *Forward[T]) Length() int {
	if f.err != nil {
		return 0
	}
	return len(f.data)
}

func (f *Forward[T]) zero() T {
	var zero T
	return zero
}
