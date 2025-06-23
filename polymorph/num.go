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

import "reflect"

// NumMethods returns the number of the T methods.
// Inside the function, it uses reflect.Type.NumMethod().
// See https://pkg.go.dev/reflect#Value.NumMethod for more.
func NumMethods[T any]() int {
	return TypeOfGeneric[T]().NumMethod()
}

// NumFields returns the number of the T fields.
// Returns zero if S is not structure,
// as it's required by https://pkg.go.dev/reflect#Value.NumField.
func NumFields[S any]() int {
	if KindOf[S]() != reflect.Struct {
		return 0
	}
	return TypeOfGeneric[S]().NumField()
}
