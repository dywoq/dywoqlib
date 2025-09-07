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

var (
	ErrKeyAlreadyExist        = errors.New("github.com/dywoq/dywoqlib/container/mapn: key already exists")
	ErrKeyNotFound            = errors.New("github.com/dywoq/dywoqlib/container/mapn: key not found")
	ErrNegativeFixedLength    = errors.New("github.com/dywoq/dywoqlib/container/mapn: negative fixed length")
	ErrFixedLengthOutOfBounds = errors.New("github.com/dywoq/dywoqlib/container/mapn: fixed length out of bounds")
	ErrOutOfBounds            = errors.New("github.com/dywoq/dywoqlib/container/mapn: out of bounds")
)
