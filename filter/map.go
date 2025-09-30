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

// Map returns a filtered map of keys that satisfy pred.
// Returns an empty map if len(m) is 0.
func Map[K comparable, V any](m map[K]V, pred func(K, V) bool) map[K]V {
	if len(m) == 0 {
		return map[K]V{}
	}
	result := map[K]V{}
	for k, v := range m {
		if pred(k, v) {
			result[k] = v
		}
	}
	return result
}

// Map returns a filtered map of keys that don't satisfy pred.
// Returns an empty map if len(m) is 0.
func MapNot[K comparable, V any](m map[K]V, pred func(K, V) bool) map[K]V {
	return Map(m, func(k K, v V) bool { return !pred(k, v) })
}

// MapKeys returns a filtered map with the keys that satisfy pred.
// Returns an empty map if len(m) is 0.
func MapKeys[K comparable, V any](m map[K]V, pred func(K) bool) map[K]V {
	return Map(m, func(k K, v V) bool { return pred(k) })
}

// MapValues returns a filtered map with the values, and their keys, that satisfy pred.
// Returns an empty map if len(m) is 0.
func MapValues[K comparable, V any](m map[K]V, pred func(V) bool) map[K]V {
	return Map(m, func(k K, v V) bool { return pred(v) })
}
