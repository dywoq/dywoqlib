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

import "sync"

// Fixed is a generic, fixed-length map container that wraps a Dynamic map.
// It enforces a fixed length and can store an error state.
// K and V must be comparable.
type Fixed[K, V comparable] struct {
	err      error
	fixedLen int
	m        *Dynamic[K, V]
	mu       sync.Mutex
}

// NewFixed creates a new Fixed map container with a specified fixed length and initial map values.
// It returns a pointer to a Fixed[K, V] instance. If the provided fixedLen is negative, less than
// the length of the initial map, or if there is an error initializing the underlying dynamic map,
// the returned Fixed will contain the appropriate error. The function ensures that the resulting
// container has a capacity of at least fixedLen and is initialized with the contents of m.
func NewFixed[K, V comparable](fixedLen int, m map[K]V) *Fixed[K, V] {
	d := NewDynamic(map[K]V{})
	if d.Error() != nil {
		return &Fixed[K, V]{d.Error(), fixedLen, nil, sync.Mutex{}}
	}
	if fixedLen < 0 {
		return &Fixed[K, V]{ErrNegativeFixedLength, fixedLen, nil, sync.Mutex{}}
	}
	if fixedLen < len(m) {
		return &Fixed[K, V]{ErrFixedLengthOutOfBounds, fixedLen, nil, sync.Mutex{}}
	}
	if len(m) > fixedLen {
		return &Fixed[K, V]{ErrOutOfBounds, fixedLen, nil, sync.Mutex{}}
	}
	d.Grow(fixedLen)
	for key, value := range m {
		d.Add(key, value)
	}
	return &Fixed[K, V]{nil, fixedLen, d, sync.Mutex{}}
}

// Length returns the number of key-value pairs currently stored in the Fixed map.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[K, V]) Length() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.m.Length()
}

// Error returns the error associated with the Fixed container, if any.
// It implements the error interface for the Fixed type.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[K, V]) Error() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.err
}

// Exists checks whether the specified key exists in the Fixed map.
// It first verifies the internal error state; if an error is present, it returns false.
// Otherwise, it delegates the existence check to the underlying map.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[K, V]) Exists(reqkey K) bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return false
	}
	return f.m.Exists(reqkey)
}

// Add inserts the specified key-value pair (reqkey, reqvalue) into the Fixed map.
// It returns the resulting key and value after insertion. If an error occurs during
// the operation, the error is stored internally and the zero values for K and V are returned.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[K, V]) Add(reqkey K, reqvalue V) (k K, v V) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return
	}
	res1, res2 := f.m.Add(reqkey, reqvalue)
	if f.m.Error() != nil {
		f.err = f.m.Error()
	}
	k = res1
	v = res2
	return
}

// Set inserts or updates the value associated with the given key in the Fixed map.
// If there are any existing errors in the Fixed instance, the operation is skipped and zero values are returned.
// The method returns the key and value as stored in the underlying map.
// If an error occurs during the operation, it is recorded in the Fixed instance's error field.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[K, V]) Set(reqkey K, reqvalue V) (k K, v V) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return
	}
	res1, res2 := f.m.Set(reqkey, reqvalue)
	if f.m.Error() != nil {
		f.err = f.m.Error()
	}
	k = res1
	v = res2
	return
}

// Keys returns a slice containing all the keys present in the Fixed map.
// If there are any existing errors in the Fixed instance or in the underlying map,
// it returns an empty slice and sets the error state accordingly.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[K, V]) Keys() []K {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return []K{}
	}
	keys := f.m.Keys()
	if f.m.Error() != nil {
		f.err = f.m.Error()
		return []K{}
	}
	return keys
}

// Values returns a slice containing all the values stored in the Fixed map.
// If there are any errors detected by errorsOk or if the underlying map has an error,
// it returns an empty slice and sets the error state accordingly.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[K, V]) Values() []V {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return []V{}
	}
	values := f.m.Values()
	if f.m.Error() != nil {
		f.err = f.m.Error()
		return []V{}
	}
	return values
}

// Delete removes the entry with the specified key (reqkey) from the Fixed map.
// If there are any existing errors in the Fixed instance, the operation is aborted.
// After attempting deletion, if an error occurs in the underlying map, it is stored in the Fixed instance's error field.
// The method returns the key that was deleted, or the zero value of K if the operation was not successful.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[K, V]) Delete(reqkey K) (k K) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return
	}
	res1 := f.m.Delete(reqkey)
	if f.m.Error() != nil {
		f.err = f.m.Error()
		return
	}
	k = res1
	return
}

// Get retrieves the value associated with the provided key from the Fixed map.
// It returns the key and its corresponding value. If an error occurs during the
// retrieval process or if the Fixed map is in an error state, the returned key
// and value will be their zero values. Any error encountered is stored in the
// Fixed struct's err field.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[K, V]) Get(reqkey K) (k K, v V) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return
	}
	res1, res2 := f.m.Get(reqkey)
	if f.m.Error() != nil {
		f.err = f.m.Error()
		return
	}
	k = res1
	v = res2
	return
}

// String returns the string representation of the Fixed map.
// If there are any errors detected by errorsOk or from the underlying map's Error method,
// it sets the error field and returns an empty string.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[K, V]) String() string {
	f.mu.Lock()
	defer f.mu.Unlock()
	if ok := f.errorsOk(); !ok {
		return ""
	}
	res := f.m.String()
	if f.m.Error() != nil {
		f.err = f.m.Error()
		return ""
	}
	return res
}

// Native returns the underlying map.
// Locks the mutex and unlocks after the completing.
func (f *Fixed[K, V]) Native() map[K]V {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.m.Native()
}

func (f *Fixed[K, V]) outOfBounds() bool {
	return f.m.Length() > f.fixedLen
}

func (f *Fixed[K, V]) errorsOk() bool {
	if f.err != nil {
		return false
	}
	if f.m.Error() != nil {
		f.err = f.m.Error()
		return false
	}
	if f.outOfBounds() {
		f.err = ErrOutOfBounds
		return false
	}
	return true
}
