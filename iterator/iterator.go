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

// Iterator represents a generic iterator that can traverse over a slice of any type.
// It keeps track of the current position, the data slice, and any error encountered.
type Iterator[T any] struct {
	pos  int
	data []T
	err  error
}

// Native returns the internal slice.
func (i *Iterator[T]) Native() []T {
	return i.data
}

// Err returns the first error that was encountered by the iterator.
func (i *Iterator[T]) Err() error {
	return i.err
}

// Position returns the current position of the iterator.
// It returns -1 if an error occurred or the position is invalid.
func (i *Iterator[T]) Position() int {
	if i.err != nil || !i.isValidPosition(i.pos) {
		return -1
	}
	return i.pos
}

// Next advances the iterator to the next position.
// It returns true if the iterator successfully moved to the next valid position, false otherwise.
func (i *Iterator[T]) Next() bool {
	if i.err != nil {
		return false
	}
	i.pos++
	return i.isValidPosition(i.pos)
}

// Value returns the value at the current position of the iterator.
// It returns the zero value of type T and sets an error if the position is invalid or an error already exists.
func (i *Iterator[T]) Value() T {
	if i.err != nil || !i.isValidPosition(i.pos) {
		return i.zeroValueAndSetError()
	}
	return i.data[i.pos]
}

func (i *Iterator[T]) isValidPosition(pos int) bool {
	return pos >= 0 && pos < len(i.data)
}

func (i *Iterator[T]) zeroValueAndSetError() T {
	var zero T
	if i.err == nil {
		i.err = ErrInvalidPosition
	}
	return zero
}
