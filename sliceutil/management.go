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

// sliceutil is primarily used by containers to prevent rewriting functionality every time
package sliceutil

import (
	"fmt"
	"strings"

	"github.com/dywoq/dywoqlib/iterator"
)

type Management[T comparable] struct {
	it  iterator.Iterable[T]
	err error
}

func (m *Management[T]) SetIterableType(it iterator.Iterable[T]) {
	m.it = it
}

func (m *Management[T]) Err() error {
	return m.err
}

func (m *Management[T]) Find(reqElem T) T {
	if m.err != nil {
		return m.Zero()
	}

	it := m.it.Begin()
	if it.Native() == nil {
		return m.Zero()
	}

	found := false
	var elem T

	if it.Value() == reqElem {
		found = true
		elem = it.Value()
	} else {
		for it.Next() {
			if it.Value() == reqElem {
				found = true
				elem = it.Value()
				break
			}
		}
	}

	if it.Err() != nil {
		m.err = it.Err()
		return m.Zero()
	}

	if !found {
		m.err = ErrNotFound
		return m.Zero()
	}

	m.err = nil
	return elem
}

func (m Management[T]) At(i int) T {
	if m.err != nil {
		return m.Zero()
	}

	it := m.it.Begin()
	n := it.Native()
	if n == nil {
		m.err = ErrSliceIsNil
		return m.Zero()
	}

	if i < 0 || i >= len(n) {
		m.err = ErrIndexOutOfBounds
		return m.Zero()
	}

	m.err = nil
	return n[i]
}

// used to implement fmt.Stringer interface
func (m Management[T]) Format() string {
	var b strings.Builder
	it := m.it.Begin()
	n := it.Native()

	b.WriteString("[")

	if len(n) > 0 {
		b.WriteString(fmt.Sprintf("%v", it.Value()))

		for it.Next() {
			b.WriteString(fmt.Sprintf(", %v", it.Value()))
		}
	}

	b.WriteString("]")
	return b.String()
}

func (m Management[T]) Zero() T {
	var zero T
	return zero
}
