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

package sliceutil

import (
	"fmt"
	"slices"
	"strings"

	"github.com/dywoq/dywoqlib/iterator"
)

// Format returns a formatted string representation of a slice,
// enclosed in brackets and with elements separated by commas.
// For example, a slice [1, 2, 3] is formatted as "[1, 2, 3]".
func Format[T comparable](s []T) string {
	if len(s) == 0 {
		return "[]"
	}
	var b strings.Builder
	b.WriteString("[")
	for i, elem := range s {
		fmt.Fprintf(&b, "%v", elem)
		if i != len(s)-1 {
			fmt.Fprintf(&b, ", ")
		}
	}
	b.WriteString("]")
	return b.String()
}

// Find searches for the requested element in a forward iterator.
// It returns the element and a nil error if found. If the element
// is not found, it returns the zero value of T and ErrElementNotFound.
// It also returns an error if the iterator encounters one during traversal.
func Find[T comparable](req T, it *iterator.Forward[T]) (T, error) {
	var val T
	for it.Next() {
		val = it.Value()
		if val == req {
			return val, nil
		}
	}
	if it.Error() != nil {
		return val, it.Error()
	}
	return val, ErrElementNotFound
}

// At returns the element at the given index. It returns ErrWrongIndex
// if the index is out of bounds.
func At[T comparable](i int, s []T) (T, error) {
	if i < 0 || i >= len(s) {
		var zero T
		return zero, ErrWrongIndex
	}
	return s[i], nil
}

// Set changes the value at the given index to elem and returns the old value.
// It returns ErrWrongIndex if the index is out of bounds.
func Set[T comparable](elem T, i int, s []T) (T, error) {
	if i < 0 || i >= len(s) {
		var zero T
		return zero, ErrWrongIndex
	}
	old := s[i]
	s[i] = elem
	return old, nil
}

// Delete removes the element at the given index and returns the removed value.
// It returns ErrWrongIndex if the index is out of bounds.
func Delete[T comparable](i int, s []T) (T, error) {
	if i < 0 || i >= len(s) {
		var zero T
		return zero, ErrWrongIndex
	}
	removed := s[i]
	s = slices.Delete(s, i, i+1)
	return removed, nil
}

// Insert places elem at the given index and returns the inserted value.
// It returns ErrWrongIndex if the index is out of bounds.
func Insert[T comparable](i int, s *[]T, elem T) (T, error) {
	if i < 0 || i >= len(*s) {
		var zero T
		return zero, ErrWrongIndex
	}
	*s = slices.Insert(*s, i, elem)
	return elem, nil
}

// Pop removes and returns the front element from s.
// If s is empty, it returns zero value.
func Pop[T comparable](s *[]T) T {
	slc := *s
	if len(slc) == 0 {
		var zero T
		return zero
	}
	lastIdx := len(slc) - 1
	poppedElem := slc[lastIdx]
	*s = slc[:lastIdx]
	return poppedElem
}

// Front returns the first element of s.
// If s is empty, it returns zero value.
func Front[T comparable](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	return s[0]
}

// Back returns the last element of s.
// If s is empty, it returns zero value.
func Back[T comparable](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	return s[len(s)-1]
}
