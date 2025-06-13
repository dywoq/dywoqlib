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

const initialCapacity int = 256

type Dynamic[T comparable] struct {
	data []T
	err  error
}

func NewDynamic[T comparable]() *Dynamic[T] {
	return NewDynamicWithData(make([]T, 0, initialCapacity))
}

func NewDynamicWithData[T comparable](data []T) *Dynamic[T] {
	return &Dynamic[T]{data: data, err: nil}
}

func (d *Dynamic[T]) Err() error {
	return d.err
}

func (d *Dynamic[T]) ActualLength() int {
	return len(d.data)
}

func (d *Dynamic[T]) Empty() bool {
	return d.ActualLength() == 0
}

func (d *Dynamic[T]) Begin() iterator.Iterator[T] {
	if d.err != nil {
		return iterator.Iterator[T]{}
	}
	return iterator.New(0, d.data)
}

func (d *Dynamic[T]) End() iterator.Iterator[T] {
	if d.err != nil {
		return iterator.Iterator[T]{}
	}
	return iterator.New(d.ActualLength()-1, d.data)
}

func (d *Dynamic[T]) At(index int) T {
	if index < 0 && index > len(d.data) {
		d.err = ErrWrongIndex
		var zero T
		return zero
	}
	return d.data[index]
}

func (d *Dynamic[T]) Front() T {
	if d.err != nil {
		var zero T
		return zero
	}
	return d.At(0)
}

func (d *Dynamic[T]) Back() T {
	if d.err != nil {
		var zero T
		return zero
	}
	return d.At(len(d.data) - 1)
}

func (d *Dynamic[T]) String() string {
	m := management[T]{}
	return m.formatSlice(d.data)
}

func (d *Dynamic[T]) Append(elements ...T) {
	if d.err != nil {
		return
	}
	d.data = append(d.data, elements...)
}

func (d *Dynamic[T]) Remove(index int) {
	if d.err != nil {
		return
	}
	if index < 0 && index > len(d.data) {
		d.err = ErrElementNotFound
		return
	}
	d.data = slices.Delete(d.data, index, index+1)
}

func (d *Dynamic[T]) Contains(t T) bool {
	if d.err != nil {
		return false
	}
	if !slices.Contains(d.data, t) {
		d.err = ErrElementNotFound
		return false
	}
	return true
}

func (d *Dynamic[T]) Clear() {
	if d.err != nil {
		return
	}

	for i := range d.data {
		d.data[i] = *new(T)
	}
}
