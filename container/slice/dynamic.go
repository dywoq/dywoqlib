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

package slice

import (
	"sync"

	"github.com/dywoq/dywoqlib/iterator"
	"github.com/dywoq/dywoqlib/sliceutil"
)

// Dynamic provides a thread-safe generic, error-aware wrapper around a Go slice.
type Dynamic[T comparable] struct {
	err error
	s   []T
	mu  sync.Mutex
}

// NewDynamic creates a new Dynamic slice instance with initial elements.
// It initializes the Dynamic struct with the provided elements and no error.
func NewDynamic[T comparable](elems ...T) *Dynamic[T] {
	return &Dynamic[T]{nil, elems, sync.Mutex{}}
}

// Grow increases the capacity of the underlying slice to at least i.
// If the current capacity is less than i, a new slice is allocated and elements are copied.
// Locks the mutex and unlocks after the completing.
func (d *Dynamic[T]) Grow(i int) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.err != nil {
		return
	}
	if cap(d.s) < i {
		newSlice := make([]T, len(d.s), i)
		copy(newSlice, d.s)
		d.s = newSlice
	}
}

// Native returns the underlying Go slice.
// This allows direct interaction with the standard slice if needed.
// Locks the mutex and unlocks after the completing.
func (d *Dynamic[T]) Native() []T {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.s
}

// Error returns the first error encountered during operations.
// Subsequent operations will be no-ops if an error is present.
// Locks the mutex and unlocks after the completing.
func (d *Dynamic[T]) Error() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.err
}

// Length returns the number of elements in the dynamic slice.
// It provides the current length of the wrapped slice.
// Locks the mutex and unlocks after the completing.
func (d *Dynamic[T]) Length() int {
	d.mu.Lock()
	defer d.mu.Unlock()
	return len(d.s)
}

// Iterating returns a Combined iterator for the slice.
// This allows for flexible iteration over the elements.
// Locks the mutex and unlocks after the completing.
func (d *Dynamic[T]) Iterating() *iterator.Combined[T] {
	d.mu.Lock()
	defer d.mu.Unlock()
	return iterator.NewCombined(d.s)
}

// Append adds new elements to the end of the slice.
// It modifies the underlying slice and returns the appended elements.
// Locks the mutex and unlocks after the completing.
func (d *Dynamic[T]) Append(elems ...T) []T {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.err != nil {
		return []T{}
	}
	d.s = append(d.s, elems...)
	return elems
}

// At returns the element at the specified index.
// It updates the internal error if the index is out of bounds.
// Locks the mutex and unlocks after the completing.
func (d *Dynamic[T]) At(i int) T {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.err != nil {
		return d.zero()
	}
	found, err := sliceutil.At(i, d.s)
	if err != nil {
		d.err = err
		return d.zero()
	}
	return found
}

// Find searches for the first occurrence of a requested element.
// It returns the found element or a zero value if not found or an error occurs.
// Locks the mutex and unlocks after the completing.
func (d *Dynamic[T]) Find(req T) T {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.err != nil {
		return d.zero()
	}
	found, err := sliceutil.Find(req, iterator.NewForward(d.s))
	if err != nil {
		return d.zero()
	}
	return found
}

// String returns a string representation of the slice.
// It uses sliceutil.Format to format the underlying slice.
// Locks the mutex and unlocks after the completing.
func (d *Dynamic[T]) String() string {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.err != nil {
		return ""
	}
	return sliceutil.Format(d.s)
}

// Set updates the element at a given index.
// It returns the updated element or a zero value on error.
// Locks the mutex and unlocks after the completing.
func (d *Dynamic[T]) Set(elem T, i int) T {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.err != nil {
		return d.zero()
	}
	new, err := sliceutil.Set(elem, i, d.s)
	if err != nil {
		d.err = err
		return d.zero()
	}
	return new
}

// Delete removes the element at the specified index.
// It returns the deleted element or a zero value on error.
// Locks the mutex and unlocks after the completing.
func (d *Dynamic[T]) Delete(i int) T {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.err != nil {
		return d.zero()
	}
	deleted, err := sliceutil.Delete(i, d.s)
	if err != nil {
		d.err = err
		return d.zero()
	}
	return deleted
}

// Insert adds an element at a specific index.
// It returns the inserted element or a zero value on error.
// Locks the mutex and unlocks after the completing.
func (d *Dynamic[T]) Insert(i int, elem T) T {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.err != nil {
		return d.zero()
	}
	inserted, err := sliceutil.Insert(i, &d.s, elem)
	if err != nil {
		d.err = err
		return d.zero()
	}
	return inserted
}

// Front returns the first element of the slice.
// It returns a zero value if the slice is empty or an error occurred.
// Locks the mutex and unlocks after the completing.
func (d *Dynamic[T]) Front() T {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.err != nil || len(d.s) == 0 {
		return d.zero()
	}
	return d.s[0]
}

// Back returns the last element of the slice.
// It returns a zero value if the slice is empty or an error occurred.
// Locks the mutex and unlocks after the completing.
func (d *Dynamic[T]) Back() T {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.err != nil || len(d.s) == 0 {
		return d.zero()
	}
	return d.s[len(d.s)-1]
}

// Pop removes and returns the last element of the slice.
// It returns a zero value if the slice is empty or an error occurred.
// Locks the mutex and unlocks after the completing.
func (d *Dynamic[T]) Pop() T {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.err != nil {
		return d.zero()
	}
	if len(d.s) == 0 {
		return d.zero()
	}
	lastIdx := len(d.s) - 1
	poppedElem := d.s[lastIdx]
	d.s = d.s[:lastIdx]
	return poppedElem
}

func (d *Dynamic[T]) zero() T {
	var zero T
	return zero
}
