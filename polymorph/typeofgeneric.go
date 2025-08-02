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

// Package polymorph provides generic reflection utilities for Go types,
// including type information, kind, comparability, method and field checks, and interface implementation.
package polymorph

import (
	"reflect"
	"sync"
)

var typeCache sync.Map

// TypeOfGeneric returns the reflect.Type of the generic type parameter T.
func TypeOfGeneric[T any]() reflect.Type {
	var zero T
	t := reflect.TypeOf(&zero).Elem()
	cacheKey := t.String()

	if cachedType, ok := typeCache.Load(cacheKey); ok {
		return cachedType.(reflect.Type)
	}

	typ := reflect.TypeOf(zero)
	typeCache.Store(cacheKey, typ)
	return typ
}
