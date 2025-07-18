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

package slice

import "github.com/dywoq/dywoqlib/iterator"

type Fixed[T comparable] struct {
	err      error
	fixedLen int
	d        *Dynamic[T]
}

func NewFixed[T comparable](fixedLen int, elems ...T) *Fixed[T] {
	d := NewDynamic[T]()
	if d.Error() != nil {
		return &Fixed[T]{d.Error(), fixedLen, nil}
	}
	if fixedLen < 0 {
		return &Fixed[T]{ErrNegativeFixedLength, fixedLen, nil}
	}
	if fixedLen < len(elems) {
		return &Fixed[T]{ErrFixedLengthOutOfBounds, fixedLen, nil}
	}
	if len(elems) > fixedLen {
		return &Fixed[T]{ErrOutOfBounds, fixedLen, nil}
	}
	d.Grow(fixedLen)
	d.Append(elems...)
	return &Fixed[T]{nil, fixedLen, d}
}

func (f *Fixed[T]) Native() []T {
	return f.d.Native()
}

func (f *Fixed[T]) Error() error {
	return f.err
}

func (f *Fixed[T]) Length() int {
	return f.d.Length()
}

func (f *Fixed[T]) Iterating() *iterator.Combined[T] {
	return f.Iterating()
}

func (f *Fixed[T]) Append(elems ...T) []T {
	if ok := f.errorsOk(); !ok {
		return []T{}
	}
	appended := f.Append(elems...)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return []T{}
	}
	return appended
}

func (f *Fixed[T]) At(i int) T {
	if ok := f.errorsOk(); !ok {
		return f.zero()
	}
	got := f.d.At(i)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return got
}

func (f *Fixed[T]) Find(req T) T {
	if ok := f.errorsOk(); !ok {
		return f.zero()
	}
	found := f.d.Find(req)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return found
}

func (f *Fixed[T]) String() string {
	if ok := f.errorsOk(); !ok {
		return ""
	}
	formatted := f.d.String()
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return ""
	}
	return formatted
}

func (f *Fixed[T]) Set(elem T, i int) T {
	if ok := f.errorsOk(); !ok {
		return f.zero()
	}
	new := f.d.Set(elem, i)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return new
}

func (f *Fixed[T]) Delete(i int) T {
	if ok := f.errorsOk(); !ok {
		return f.zero()
	}
	deleted := f.d.Delete(i)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return deleted
}

func (f *Fixed[T]) Insert(i int, elem T) T {
	if ok := f.errorsOk(); !ok {
		return f.zero()
	}
	inserted := f.d.Insert(i, elem)
	if f.d.Error() != nil {
		f.err = f.d.Error()
		return f.zero()
	}
	return inserted
}

func (f *Fixed[T]) outOfBounds() bool {
	return len(f.d.s) > f.fixedLen
}

func (f *Fixed[T]) errorsOk() bool {
	if f.fixedLen < len(f.d.s) {
		f.err = ErrFixedLengthOutOfBounds
		return false
	}
	if f.err != nil {
		return false
	}
	if f.outOfBounds() {
		f.err = ErrOutOfBounds
		return false
	}
	return true
}

func (f *Fixed[T]) zero() T {
	var zero T
	return zero
}
