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

package polymorphic

import "reflect"

// Implements checks if I is implemented by S.
// Also, the function can return false if I is not an interface.
// The function doesn't panics.
func Implements[I any, S any]() bool {
	if !IsKind[I](reflect.Interface) {
		return false
	}
	tInterface := TypeOfGeneric[I]()
	tStruct := TypeOfGeneric[S]()
	return tStruct.Implements(tInterface)
}
