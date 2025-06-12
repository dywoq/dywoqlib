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
	"slices"

	"github.com/dywoq/dywoqlib/slice/iterator"
)

type Fixed[T any] struct {
	data          []T
	initialLength int
	err           error
}

func NewFixed[T any](initialLength int) *Fixed[T] {
	if initialLength < 0 {
		return &Fixed[T]{err: ErrNegativeInitialLength}
	}
	return &Fixed[T]{make([]T, initialLength), initialLength, nil}
}

func NewFixedWithData[T any](initialLength int, data []T) *Fixed[T] {
	if len(data) > initialLength {
		return &Fixed[T]{err: ErrOverInitialLength}
	}
	if initialLength < 0 {
		return &Fixed[T]{err: ErrNegativeInitialLength}
	}
	return &Fixed[T]{data, initialLength, nil}
}

func (f *Fixed[T]) Err() error {
	return f.err
}

func (f *Fixed[T]) InitialLength() int {
	return f.initialLength
}

func (f *Fixed[T]) ActualLength() int {
	return len(f.data)
}

func (f *Fixed[T]) Empty() bool {
	return f.ActualLength() == 0
}

func (f *Fixed[T]) OverInitialLength() bool {
	return f.ActualLength() > f.InitialLength()
}

func (f *Fixed[T]) Negative() bool {
	return f.ActualLength() < 0
}

func (f *Fixed[T]) Begin() iterator.Iterator[T] {
	f.updateErrorState()
	if f.err != nil {
		return iterator.Iterator[T]{}
	}
	return iterator.New(0, f.data)
}

func (f *Fixed[T]) End() iterator.Iterator[T] {
	f.updateErrorState()
	if f.err != nil {
		return iterator.Iterator[T]{}
	}
	return iterator.New(f.ActualLength()-1, f.data)
}

func (f *Fixed[T]) At(index int) T {
	f.updateErrorState()
	return f.data[index]
}

func (f *Fixed[T]) Front() T {
	f.updateErrorState()
	if f.err != nil {
		var zero T
		return zero
	}
	return f.At(0)
}

func (f *Fixed[T]) Back() T {
	f.updateErrorState()
	if f.err != nil {
		var zero T
		return zero
	}
	return f.At(len(f.data) - 1)
}

func (f *Fixed[T]) String() string {
	m := management[T]{}
	return m.formatSlice(f.data)
}

func (f *Fixed[T]) Set(index int, value T) {
	f.updateErrorState()
	if f.err != nil {
		return
	}
	f.data[index] = value
}

func (f *Fixed[T]) Append(elements ...T) {
	f.updateErrorState()
	if f.err != nil {
		return
	}
	f.data = append(f.data, elements...)
}

func (f *Fixed[T]) Remove(index int) {
	f.updateErrorState()
	if f.err != nil {
		return
	}
	f.data = slices.Delete(f.data, index, index+1)
}

func (f *Fixed[T]) updateErrorState() {
	errs := make(map[bool]error, 3)
	errs[f.Empty()] = ErrNoElements
	errs[f.Negative()] = ErrNegativeInitialLength
	errs[f.OverInitialLength()] = ErrOverInitialLength

	for condition, err := range errs {
		if condition {
			f.err = err
			break
		}
	}
}
