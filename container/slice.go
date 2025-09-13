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

package container

import (
	"github.com/dywoq/dywoqlib/iterator"
	"github.com/dywoq/dywoqlib/sliceutil"
)

// IterableSlice is a slice of comparable generic parameter T
// for iteration over standard Go slice.
type IterableSlice[T comparable] []T

// FormattableSlice is a slice of comparable generic parameter T,
// for more readable formatting of standard Go slice.
type FormattableSlice[T comparable] []T

// GrowableContainer is a slice of comparable generic parameter T
// for safe pre-allocation of standard Go slice, implementing allocation.Sizer interface.
type GrowableSlice[T comparable] []T

// Format returns the formatted string of FormattableSlice[T].
func (f FormattableSlice[T]) String() string { return sliceutil.Format(f) }

// Iterating returns a pointer to iterator.Combined[T] structure.
// It uses a factory method iterator.NewCombined() internally.
func (it IterableSlice[T]) Iterating() *iterator.Combined[T] { return iterator.NewCombined(it) }

// Grow increases the capacity of the slice to i.
// If the current capacity is less than i, a new slice is allocated and elements are copied.
// It will panic if i is negative.
func (g *GrowableSlice[T]) Grow(i int) {
	if i < 0 {
		panic("github.com/dywoq/dywoqlib/container: capacity cannot be negative")
	}

	if cap(*g) < i {
		newSlice := make([]T, len(*g), i)
		copy(newSlice, *g)
		*g = newSlice
	}
}
