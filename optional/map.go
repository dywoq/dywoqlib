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

// Map applies a function to the m value if it's present and returns a new Maybe
// with the result. Otherwise, it returns an empty Maybe of the new type.
func Map[T, U any](m Maybe[T], f func(T) U) Maybe[U] {
	if m.Present() {
		v, _ := m.Get()
		res := f(v)
		return New(res)
	}
	return None[U]()
}
