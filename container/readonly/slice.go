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

package readonly

import (
	"fmt"
	"sync"

	"github.com/dywoq/dywoqlib/err"
	"github.com/dywoq/dywoqlib/iterator"
	"github.com/dywoq/dywoqlib/sliceutil"
)

// Slice is a generic readonly wrapper around the standard Go slice.
type Slice[T comparable] struct {
	s   []T
	mu  sync.Mutex
	err err.Context
}

// NewSlice creates new a pointer to read-only container slice.
func NewSlice[T comparable](s ...T) *Slice[T] {
	return &Slice[T]{s, sync.Mutex{}, err.NoneContext()}
}

// Error returns the possibly encountered current error.
// If error doesn't present, the function returns err.NoneContext().
// The mutex locks and unlocks after the function completed.
func (s *Slice[T]) Error() err.Context {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.err
}

// Length returns the length of the underlying slice
// If error is present, it returns zero.
// The mutex locks and unlocks after the function completed.
func (s *Slice[T]) Length() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.err.Nil() {
		return 0
	}
	return len(s.s)
}

// At returns the element at i.
// If error is present, it returns zero value.
// The mutex locks and unlocks after the function completed.
func (s *Slice[T]) At(i int) T {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.err.Nil() {
		return s.zero()
	}
	res, err1 := sliceutil.At(i, s.s)
	if err1 != nil {
		s.err.SetError(err1)
		s.err.SetMore(s.errContextSource("At(int) T"))
		return s.zero()
	}
	return res
}

// Find finds req in the underlying slice
// and returns it if the finding was successful, otherwise,
// it updates the internal error state.
// If error is present, it returns zero value.
// The mutex locks and unlocks after the function completed.
func (s *Slice[T]) Find(req T) T {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.err.Nil() {
		return s.zero()
	}
	res, err1 := sliceutil.Find(req, iterator.NewForward(s.s))
	if err1 != nil {
		s.err.SetError(err1)
		s.err.SetMore(s.errContextSource("Find(T) T"))
		return s.zero()
	}
	return res
}

// String returns a string representation of the slice.
// It uses sliceutil.Format to format the underlying slice.
// The mutex locks and unlocks after the function completed.
func (s *Slice[T]) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.err.Nil() {
		return ""
	}
	return sliceutil.Format(s.s)
}

// Front returns the first element of the slice.
// It returns a zero value if the slice is empty or an error occurred.
// The mutex locks and unlocks after the function completed.
func (s *Slice[T]) Front() T {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.err.Nil() {
		return s.zero()
	}
	res := sliceutil.Front(s.s)
	return res
}

// Back returns the last element of the slice.
// It returns a zero value if the slice is empty or an error occurred.
// The mutex locks and unlocks after the function completed.
func (s *Slice[T]) Back() T {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.err.Nil() {
		return s.zero()
	}
	res := sliceutil.Back(s.s)
	return res
}

func (s *Slice[T]) zero() T {
	var zero T
	return zero
}

func (s *Slice[T]) errContextSource(method string) string {
	return fmt.Sprintf("source is readonly.Slice[T comparable].%s", method)
}
