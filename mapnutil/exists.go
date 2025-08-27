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

package mapnutil

// Exists reports whether reqkey exists in map m.
// If map is empty, it returns false.
func Exists[K, V comparable](m map[K]V, reqkey K) bool {
	if len(m) == 0 {
		return false
	}
	for key := range m {
		if key == reqkey {
			return true
		}
	}
	return false
}
