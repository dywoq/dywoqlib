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

// Implements checks if S implements I.
// Returns false if I is not an interface.
func Implements[I any, S any]() bool {
	sType := TypeOfGeneric[S]()
	iType := TypeOfGeneric[I]()
	if iType.Kind() != reflect.Interface {
		return false
	}
	return sType.Implements(iType)
}
