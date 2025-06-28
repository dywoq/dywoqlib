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

package iteration

// Default iterates over the slice s of type T, calling the provided function f for each element.
// The function f receives the index and the element as arguments, and should return a boolean indicating
// whether to continue iteration (true) or stop (false). The iteration stops early if f returns false.
// T must be a comparable type.
func Default[T comparable](s []T, f Type[T]) {
	for i, elem := range s {
		keepGoing := f(i, elem)
		if !keepGoing {
			break
		}
	}
}
