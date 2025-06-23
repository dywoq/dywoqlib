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

package iterator

import "errors"

// ErrNoMoreElements is returned when an iterator has no more elements to return.
var ErrNoMoreElements = errors.New("github.com/dywoq/dywoqlib/container/iterator: no more elements")

// ErrOutOfBounds is returned when an iterator is accessed out of its valid range.
var ErrOutOfBounds = errors.New("github.com/dywoq/dywoqlib/container/iterator: out of bounds")

// ErrInvalidPosition is returned when an iterator is in an invalid position.
var ErrInvalidPosition = errors.New("github.com/dywoq/dywoqlib/container/iterator: invalid position")
