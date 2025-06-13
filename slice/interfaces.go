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

// ErrorChecker interface for types that can report an error state.
type ErrorChecker interface {
	Err() error
}

// LengthChecker interface for types that provide information about their length and capacity.
type LengthChecker interface {
	ActualLength() int
	Empty() bool
}

// OverInitialLengthChecker interface for types that can check if a slice is over the initial length.
type OverInitialLengthChecker interface {
	OverInitialLength() bool
}

// NegativeLengthChecker interface for types that can check if an initial length is negative.
type NegativeLengthChecker interface {
	Negative() bool
}

// InitialLengthGetter interface for types that provide information about their initial length.
type InitialLengthGetter interface {
	InitialLength() int
}

// ElementReader interface for types whose elements can be read.
type ElementReader[T comparable] interface {
	At(index int) T
	Front() T
	Back() T
	Contains(t T) bool
}

// ElementModifier interface for types whose elements can be modified, appended, or removed.
type ElementModifier[T comparable] interface {
	Set(index int, value T)
	Append(elements ...T)
	Remove(index int)
	Clear()
}
