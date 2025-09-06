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

import "github.com/dywoq/dywoqlib/lib/attribute"

// Combined is a generic iterator struct that wraps a slice of comparable elements and provides methods to create forward and reverse iterators.
// T must be comparable.
type Combined[T comparable] struct {
	s []T
}

// Forward returns a new forward iterator for the combined iterator's slice.
func (c Combined[T]) Forward() *Forward[T] { return NewForward(c.s) }

// Forward returns a new reverse iterator for the combined iterator's slice.
func (c Combined[T]) Reverse() *Reverse[T] { return NewReserve(c.s) }

// ReadonlyForward returns a new read-only forward iterator for the combined iterator's slice.
//
// DEPRECATED, MAY BE REMOVED IN THE FUTURE
func (c Combined[T]) ReadonlyForward() *ReadonlyForward[T] {
	attribute.Deprecated(nil)
	return NewReadonlyForward(c.s)
}

// ReadonlyReverse returns a new read-only forward iterator for the combined iterator's slice.
//
// DEPRECATED, MAY BE REMOVED IN THE FUTURE
func (c Combined[T]) ReadonlyReverse() *ReadonlyReverse[T] {
	attribute.Deprecated(nil)
	return NewReadonlyReverse(c.s)
}

// NewCombined creates a new Combined iterator instance with the provided slice.
// T must be comparable.
func NewCombined[T comparable](s []T) *Combined[T] {
	return &Combined[T]{s}
}
