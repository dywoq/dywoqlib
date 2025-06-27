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

import (
	"fmt"
	"strings"
)

// Map represents a generic map with error handling.
type Map[K comparable, V comparable] struct {
	data map[K]V
	err  error
}

// Err returns the first error encountered during map operations.
// It allows checking for any errors that have occurred.
func (m *Map[K, V]) Err() error {
	return m.err
}

// Keys returns a slice of all keys in the map.
// It returns a zero slice if the map is empty or an error occurred.
func (m *Map[K, V]) Keys() []K {
	if m.err != nil {
		var zero []K
		return zero
	}
	if len(m.data) == 0 {
		m.err = ErrEmpty
		var zero []K
		return zero
	}
	keys := []K{}
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice of all values in the map.
// It returns a zero slice if the map is empty or an error occurred.
func (m *Map[K, V]) Values() []V {
	if m.err != nil {
		var zero []V
		return zero
	}
	if len(m.data) == 0 {
		m.err = ErrEmpty
		var zero []V
		return zero
	}
	values := []V{}
	for _, v := range m.data {
		values = append(values, v)
	}
	return values
}

// At returns the value associated with the given key.
// It returns a zero value if the key is not found, the map is empty, or an error occurred.
func (m *Map[K, V]) At(key K) V {
	if m.err != nil {
		var zero V
		return zero
	}
	if len(m.data) == 0 {
		m.err = ErrEmpty
		var zero V
		return zero
	}
	foundValue, ok := m.data[key]
	if !ok {
		m.err = ErrNotFound
		var zero V
		return zero
	}
	return foundValue
}

// String provides a string representation of the map.
// It formats the map's key-value pairs for display.
func (m *Map[K, V]) String() string {
	var b strings.Builder
	b.WriteString("[")
	for k, v := range m.data {
		b.WriteString(fmt.Sprintf("\n  %v: %v", k, v))
	}
	b.WriteString("\n]")
	return b.String()
}

// Insert adds a new key-value pair to the Map. 
// If the Map has a non-nil error state, the operation is aborted.
// If the Map is empty, it sets the error to ErrEmpty and returns.
// If the key already exists in the Map, it sets the error to ErrKeyAlreadyExist and returns.
// Otherwise, it inserts the key-value pair into the Map.
func (m *Map[K, V]) Insert(reqKey K, reqValue V) {
	if m.err != nil {
		return
	}
	if len(m.data) == 0 {
		m.err = ErrEmpty
		return
	}

	for key := range m.data {
		if reqKey == key {
			m.err = ErrKeyAlreadyExist
			return
		}
	}
	m.data[reqKey] = reqValue
}
