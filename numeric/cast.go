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

package numeric

import (
	"fmt"
	"github.com/dywoq/dywoqlib/err"
	"github.com/dywoq/dywoqlib/numeric/constraints"
	"github.com/dywoq/dywoqlib/numeric/limits"
)

// Cast safely converts b to Base[To] integer. First, it checks if b fits the
// limits of integer To. If b doesn't, it returns zero number, and error.
// If b does, Cast converts b into Base[To].
func Cast[From, To constraints.Integral](b Base[From]) (Base[To], err.Context) {
	val := b.Get()
	minTo, maxTo := limits.Numeric[To]()

	if isSigned[From]() {
		if isSigned[To]() {
			if int64(val) < int64(minTo) || int64(val) > int64(maxTo) {
				return baseFactoryMethod[To](), err.NewContext(ErrOverflow, fmt.Sprintf("%d overflows %T", val, *new(To)))
			}
		} else {
			if val < 0 {
				return baseFactoryMethod[To](), err.NewContext(ErrOverflow, fmt.Sprintf("%d overflows %T", val, *new(To)))
			}
			if uint64(val) > uint64(maxTo) {
				return baseFactoryMethod[To](), err.NewContext(ErrOverflow, fmt.Sprintf("%d overflows %T", val, *new(To)))
			}
		}
	} else {
		if isSigned[To]() {
			if uint64(val) > uint64(maxTo) {
				return baseFactoryMethod[To](), err.NewContext(ErrOverflow, fmt.Sprintf("%d overflows %T", val, *new(To)))
			}
		} else {
			if uint64(val) > uint64(maxTo) {
				return baseFactoryMethod[To](), err.NewContext(ErrOverflow, fmt.Sprintf("%d overflows %T", val, *new(To)))
			}
		}
	}

	return baseFactoryMethod(To(val)), nil
}

func isSigned[T constraints.Integral]() bool {
	return T(0)-T(1) < T(0)
}
