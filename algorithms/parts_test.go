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

package algorithms_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/algorithms"
)

func ExampleParts() {
	parts := algorithms.Parts(100, 4)
	for _, num := range parts {
		fmt.Printf("num: %v\n", num)
	}
	// Output:
	// num: 20
	// num: 40
	// num: 60
	// num: 80
}
