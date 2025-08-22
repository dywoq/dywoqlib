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

package optional

import (
	"fmt"
)

type implementation[T any] struct {
	value   T
	present bool
}

func (o *implementation[T]) Present() bool {
	return o.present
}

func (o *implementation[T]) Get() (T, bool) {
	if !o.present {
		var zero T
		return zero, false
	}
	return o.value, true
}

func (o *implementation[T]) String() string {
	if o.present {
		return fmt.Sprintf("%v", o.value)
	}
	var zero T
	return fmt.Sprintf("%v", zero)
}

func (o *implementation[T]) Else(other T) T {
	if o.present {
		return o.value
	}
	return other
}

func (o *implementation[T]) Filter(filter func(T) bool) Maybe[T] {
	if o.present && filter(o.value) {
		return o
	}
	return None[T]()
}

func (o *implementation[T]) Unwrap() T {
	if !o.present {
		panic(ErrNotPresent)
	}
	return o.value
}

func (o *implementation[T]) Or(f func() T) T {
	if o.present {
		return o.value
	}
	return f()
}
