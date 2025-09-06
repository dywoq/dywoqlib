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

package stringn

import "errors"

var (
	ErrIndexOutOfBounds       = errors.New("github.com/dywoq/dywoqlib/lib/stringn: index out of bounds")
	ErrRuneNotFound           = errors.New("github.com/dywoq/dywoqlib/lib/stringn: rune not found ")
	ErrInvalidIndexForRemoval = errors.New("github.com/dywoq/dywoqlib/lib/stringn: invalid index for removal")
)
