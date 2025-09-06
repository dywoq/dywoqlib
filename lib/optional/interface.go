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

	"github.com/dywoq/dywoqlib/lib/err"
)

// Maybe is an interface representing a optional value.
type Maybe[T any] interface {
	fmt.Stringer
	// Present checks if the optional value is present.
	Present() bool
	// Get returns the value and a boolean indicating its presence.
	// The boolean always match what Present() returns.
	Get() (T, bool)
	// Else returns the value if it's present, otherwise it returns a default value.
	Else(T) T
	// Filter returns the Maybe if it's present and the value satisfies
	// the provided predicate function. Otherwise, it returns an empty Maybe.
	Filter(func(T) bool) Maybe[T]
	// Unwrap returns the value if it's present, otherwise it panics.
	Unwrap() T
	// Or returns the value if it's present, otherwise it returns the result of the
	// provided function.
	Or(func() T) T
	// Error returns the associated error context if the value is not present.
	Error() err.Context
}

// New retruns a new Maybe with a value of a generic parameter T.
func New[T any](val T) Maybe[T] {
	return &implementation[T]{val, true, err.NewContext(nil, "")}
}

// None creates a new Maybe with no value,
// but a generic parameter T must be still present.
func None[T any]() Maybe[T] {
	return &implementation[T]{present: false, e: err.NewContext(nil, "")}
}

// None creates a new Maybe with no value,
// but a generic parameter T must be still present.
// Error context can be provided unlike None.
func NoneContext[T any](e err.Context) Maybe[T] {
	return &implementation[T]{present: false, e: e}
}
