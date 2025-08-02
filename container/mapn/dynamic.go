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
	"maps"
	"strings"
)

// Dynamic is a generic container that wraps a map with keys of type K and values of type V.
// It also includes an error field to track any errors associated with the map operations.
// K and V must be comparable types.
type Dynamic[K, V comparable] struct {
	err error
	m   map[K]V
}

// NewDynamic creates and returns a new Dynamic instance wrapping the provided map m.
// K and V are the key and value types of the map, which must be comparable.
// The returned Dynamic can be used to perform dynamic operations on the underlying map.
func NewDynamic[K, V comparable](m map[K]V) *Dynamic[K, V] {
	return &Dynamic[K, V]{nil, m}
}

// Length returns the number of key-value pairs currently stored in the Dynamic map.
func (d *Dynamic[K, V]) Length() int {
	return len(d.m)
}

// Error returns the last error encountered by the Dynamic container, or nil if no error has occurred.
func (d *Dynamic[K, V]) Error() error {
	return d.err
}

// Grow increases the capacity of the underlying map to accommodate at least i elements.
// If the map is empty, it initializes the map with the specified capacity.
// If the map already contains elements, it creates a new map with a capacity equal to
// the greater of the current length or i, and copies all existing elements into it.
// If an error is present in the Dynamic instance, the method returns immediately.
func (d *Dynamic[K, V]) Grow(i int) {
	if d.err != nil {
		return
	}
	if len(d.m) == 0 {
		d.m = make(map[K]V, i)
		return
	}

	newMap := make(map[K]V, max(len(d.m), i))
	maps.Copy(newMap, d.m)
	d.m = newMap
}

// Exists checks if the specified key exists in the dynamic map.
// It returns true if the key is present, and false otherwise.
// If the Dynamic instance has an error state (d.err != nil), it always returns false.
func (d *Dynamic[K, V]) Exists(reqkey K) (exists bool) {
	if d.err != nil {
		return false
	}
	_, exists = d.m[reqkey]
	return
}

// Add attempts to add a key-value pair (reqkey, reqvalue) to the Dynamic map.
// If the key does not already exist, it inserts the pair and returns the key and value.
// If the key already exists, it sets an error (ErrKeyAlreadyExist) and returns zero values.
// If the Dynamic instance has a pre-existing error, the operation is skipped and zero values are returned.
// Returns the inserted key and value on success, or zero values on failure.
func (d *Dynamic[K, V]) Add(reqkey K, reqvalue V) (k K, v V) {
	if d.err != nil {
		return
	}
	if !d.Exists(reqkey) {
		d.m[reqkey] = reqvalue
		k = reqkey
		v = reqvalue
		return
	}
	d.err = ErrKeyAlreadyExist
	return
}

// Set attempts to update the value associated with the given key in the Dynamic map.
// If the key exists, it sets the value to reqvalue and returns the key and value.
// If the key does not exist, it sets an internal error (ErrKeyNotFound) and returns zero values.
// If a previous error exists in the Dynamic instance, the method returns immediately with zero values.
func (d *Dynamic[K, V]) Set(reqkey K, reqvalue V) (k K, v V) {
	if d.err != nil {
		return
	}
	if d.Exists(reqkey) {
		d.m[reqkey] = reqvalue
		k = reqkey
		v = reqvalue
		return
	}
	d.err = ErrKeyNotFound
	return
}

// Keys returns a slice containing all the keys present in the Dynamic map.
// If an internal error has occurred, it returns an empty slice.
// The method uses a fixed-size slice to collect the keys efficiently.
// If an error occurs while appending keys, it sets the internal error state and returns an empty slice.
func (d *Dynamic[K, V]) Keys() []K {
	if d.err != nil {
		return []K{}
	}
	keys := make([]K, 0, len(d.m))
	for k := range d.m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns a slice containing all the values stored in the Dynamic map.
// If an internal error has occurred, it returns an empty slice.
// The method collects the values using a fixed-size slice and handles any errors
// that may arise during the collection process. If an error occurs while appending
// values, it sets the internal error state and returns an empty slice.
func (d *Dynamic[K, V]) Values() []V {
	if d.err != nil {
		return []V{}
	}
	values := make([]V, 0, len(d.m))
	for _, v := range d.m {
		values = append(values, v)
	}
	return values
}

// Delete removes the entry with the specified key from the Dynamic map.
// If the key exists, it is deleted and the key is returned.
// If the key does not exist, the error field is set to ErrKeyNotFound and the zero value of K is returned.
// If the Dynamic instance already has an error, the method returns immediately without performing any operation.
func (d *Dynamic[K, V]) Delete(reqkey K) (k K) {
	if d.err != nil {
		return
	}
	if d.Exists(reqkey) {
		delete(d.m, reqkey)
		k = reqkey
		return
	}
	d.err = ErrKeyNotFound
	return
}

// Get retrieves the value associated with the given key from the Dynamic map.
// If the key exists, it returns the key and its corresponding value.
// If the key does not exist, it sets the internal error to ErrKeyNotFound and returns zero values for K and V.
// If there is a pre-existing error in the Dynamic instance, it returns zero values for K and V immediately.
func (d *Dynamic[K, V]) Get(reqkey K) (k K, v V) {
	if d.err != nil {
		return
	}
	if d.Exists(reqkey) {
		k = reqkey
		v = d.m[reqkey]
		return
	}
	d.err = ErrKeyNotFound
	return
}

// String returns a formatted string representation of the Dynamic map.
// If an error has previously occurred or the map is empty, it returns an empty string.
// The output is a multi-line string with each key-value pair on a new line, enclosed in braces.
// If formatting fails, it sets the error in the Dynamic struct and returns an empty string.
func (d *Dynamic[K, V]) String() string {
	if d.err != nil {
		return ""
	}
	if len(d.m) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("{\n")
	for key, value := range d.m {
		_, err := fmt.Fprintf(&b, "  %v: %v\n", key, value)
		if err != nil {
			d.err = err
			return ""
		}
	}
	b.WriteString("}")
	return b.String()
}

func (d *Dynamic[K, V]) Native() map[K]V {
	return d.m
}
