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

// ErrOutOfBounds indicates an attempt to access an element outside the valid
// range of the iterator's data. This error occurs when trying to retrieve a value
// at an invalid position.
var ErrOutOfBounds = errors.New("slice.iterator.Iterator: not within bounds after Next()")

// ErrNoMoreElements signifies that the iterator has reached the end of its
// data collection and no further elements are available for iteration.
var ErrNoMoreElements = errors.New("slice.iterator.Iterator: no more elements to iterate")
