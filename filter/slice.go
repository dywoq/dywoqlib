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

package filter

// Slice returns a filtered slice of elements that satisfy pred.
// Returns an empty slice if len(s) is 0.
func Slice[S any](s []S, pred func(S) bool) []S {
	if len(s) == 0 {
		return []S{}
	}
	result := make([]S, 0, len(s))
	for _, elem := range s {
		if pred(elem) {
			result = append(result, elem)
		}
	}
	return result
}

// SliceNot returns a filtered slice of elements that don't satisfy pred.
// Returns an empty slice if len(s) is 0.
func SliceNot[S any](s []S, pred func(S) bool) []S {
	return Slice(s, func(s S) bool { return !pred(s) })
}
