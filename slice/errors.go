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

import "errors"

// ErrOverInitialLength occurs when a slice is over the initial length.
var ErrOverInitialLength = errors.New("github.com/dywoq/dywoqlib/slice: over the initial length")

// ErrNegativeInitialLength occurs when an initial length is negative.
var ErrNegativeInitialLength = errors.New("github.com/dywoq/dywoqlib/slice: initial length cannot be negative")

// ErrElementNotFound occurs when an element wasn't found in a slice.
var ErrElementNotFound = errors.New("github.com/dywoq/dywoqlib/slice: element wasn't found")

// ErrWrongIndex occurs when an element is not found at the specific index.
var ErrWrongIndex = errors.New("github.com/dywoq/dywoqlib/slice: the wrong index")
