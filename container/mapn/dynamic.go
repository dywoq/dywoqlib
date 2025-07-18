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

	"github.com/dywoq/dywoqlib/container/slice"
)

type Dynamic[K, V comparable] struct {
	err error
	m   map[K]V
}

func NewDynamic[K, V comparable](m map[K]V) *Dynamic[K, V] {
	return &Dynamic[K, V]{nil, m}
}

func (d *Dynamic[K, V]) Length() int {
	return len(d.m)
}

func (d *Dynamic[K, V]) Error() error {
	return d.err
}

func (d *Dynamic[K, V]) Grow(i int) {
	if d.err != nil {
		return
	}
	if len(d.m) == 0 {
		d.m = make(map[K]V, i)
		return
	}

	newMap := make(map[K]V, max(len(d.m), i))
	for k, v := range d.m {
		newMap[k] = v
	}
	d.m = newMap
}

func (d *Dynamic[K, V]) Exists(reqkey K) (exists bool) {
	if d.err != nil {
		return false
	}
	_, exists = d.m[reqkey]
	return
}

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

func (d *Dynamic[K, V]) Keys() []K {
	if d.err != nil {
		return []K{}
	}
	keys := slice.NewFixed[K](len(d.m))
	// we don't use keys.Grow() here because we don't need to preallocate the slice -
	// slice.NewFixed already does that
	// or because there's no method Grow() in slice.Fixed xD
	for key := range d.m {
		keys.Append(key)
	}
	if keys.Error() != nil {
		d.err = keys.Error()
		return []K{}
	}
	return keys.Native()
}

func (d *Dynamic[K, V]) Values() []V {
	if d.err != nil {
		return []V{}
	}
	values := slice.NewFixed[V](len(d.m))
	// there's same situation as described above
	for _, value := range d.m {
		values.Append(value)
	}
	if values.Error() != nil {
		d.err = values.Error()
		return []V{}
	}
	return values.Native()
}

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
