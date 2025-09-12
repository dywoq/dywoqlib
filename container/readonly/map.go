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

package readonly

import (
	"sync"

	"github.com/dywoq/dywoqlib/err"
	"github.com/dywoq/dywoqlib/mapnutil"
)

// Map is a generic thread-safe and readonly map container,
// with K as the key type and V as the value type.
type Map[K, V comparable] struct {
	m   map[K]V
	mu  sync.Mutex
	err err.Context
}

// Length returns the length of the underlying map.
// The method returns zero if error is present.
// The mutex locks and unlock after the completion of the function.
func (m *Map[K, V]) Length() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !m.err.Nil() {
		return 0
	}
	return len(m.m)
}

// Error returns the possibly encountered error,
// otherwise it returns err.NoneContext.
// The mutex locks and unlock after the completion of the function.
func (m *Map[K, V]) Error() err.Context {
	return m.err
}

// Exists reports whether reqkey exists.
// If error is present, it returns false.
// The mutex locks and unlock after the completion of the function.
func (m *Map[K, V]) Exists(reqkey K) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !m.err.Nil() {
		return false
	}
	_, exists := m.m[reqkey]
	return exists
}

// Keys returns a slice of keys of the underlying map.
// If error is present, it returns empty slice.
// The mutex locks and unlock after the completion of the function.
func (m *Map[K, V]) Keys() []K {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !m.err.Nil() {
		return []K{}
	}
	keys := []K{}
	for key := range m.m {
		keys = append(keys, key)
	}
	return keys
}

// Values returns a slice of values of the underlying map.
// If error is present, it returns empty slice.
// The mutex locks and unlock after the completion of the function.
func (m *Map[K, V]) Values() []V {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !m.err.Nil() {
		return []V{}
	}
	values := []V{}
	for _, value := range m.m {
		values = append(values, value)
	}
	return values
}

// Get gets reqkey and returns reqkey and its value from the underlying map.
// If the method didn't find reqkey, it returns zero values.
// If error is present, it returns zero values.
// The mutex locks and unlock after the completion of the function.
func (m *Map[K, V]) Get(reqkey K) (key K, val V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !m.err.Nil() {
		return
	}
	if _, exists := m.m[reqkey]; exists {
		key = reqkey
		val = m.m[reqkey]
	}
	return
}

// String returns a string representation of the underlying map.
// It returns an empty string if error is present.
// The mutex locks and unlock after the completion of the function.
func (m *Map[K, V]) String() string {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !m.err.Nil() {
		return ""
	}
	return mapnutil.Format(m.m)
}

// NewMap creates new a pointer to read-only container map.
func NewMap[K, V comparable](s map[K]V) *Map[K, V] {
	return &Map[K, V]{s, sync.Mutex{}, err.NoneContext()}
}
