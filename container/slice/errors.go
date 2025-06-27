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

// Package slice provides generic slice-based containers, including dynamic and fixed-size slices,
// with error handling and utility methods for element access and manipulation.
package slice

import "errors"

// ErrEmpty indicates that the slice is empty.
// This error occurs when an operation expects elements but none are present.
var ErrEmpty = errors.New("github.com/dywoq/dywoqlib/container/slice: slice is empty")

// ErrOverFixedSize indicates that an operation would exceed the fixed size of the slice.
// This error applies to Fixed slices when attempts are made to add too many elements.
var ErrOverFixedSize = errors.New("github.com/dywoq/dywoqlib/container/slice: over fixed size")

// ErrNegativeFixedSize indicates that a negative value was provided for the fixed size.
// This error occurs during the initialization of a Fixed slice with an invalid size.
var ErrNegativeFixedSize = errors.New("github.com/dywoq/dywoqlib/container/slice: negative fixed size")
