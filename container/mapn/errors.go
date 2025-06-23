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

package mapn

import "errors"

// ErrEmpty indicates that the map is empty.
// This error occurs when an operation expects elements but none are present.
var ErrEmpty = errors.New("github.com/dywoq/dywoqlib/container/mapn: map is empty")

// ErrNotFound indicates that the requested key was not found in the map.
// This error occurs during lookup operations when the key does not exist.
var ErrNotFound = errors.New("github.com/dywoq/dywoqlib/container/mapn: not found")