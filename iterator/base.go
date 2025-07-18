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

// Base defines a generic iterator interface for traversing a collection of elements of type T.
type Base[T comparable] interface {
	// Error returns any error encountered during iteration.
	Error() error
	// Position returns the current position.
	Position() int
	// Value returns the current element of type T.
	Value() T
	// Next advances the iterator to the next element and returns true if there is a next element.
	Next() bool
	// Reset resets the iterator to its initial state.
	Reset()
	// Length returns the current length ofthe slice.
	Length() int
}
