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

package polymorph

import (
	"reflect"
	"sync"
)

var methodCache sync.Map

type typeMethod struct {
	typ  reflect.Type
	name string
}

// HasMethod checks if T has a method with the given name.
func HasMethod[T any](name string) bool {
	key := typeMethod{
		typ:  TypeOfGeneric[T](),
		name: name,
	}
	if found, ok := methodCache.Load(key); ok {
		return found.(bool)
	}
	_, found := TypeOfGeneric[T]().MethodByName(name)
	methodCache.Store(key, found)
	return found
}

// HasMethod checks if struct S has a method with the given name.
// Returns false if S is not a structure.
func HasField[S any](name string) bool {
	tType := TypeOfGeneric[S]()
	if tType.Kind() != reflect.Struct {
		return false
	}
	_, found := tType.FieldByName(name)
	return found
}
