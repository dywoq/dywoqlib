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

package mapnutil_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/mapnutil"
)

func ExampleExists() {
	m := map[int]string{
		1: "a",
		2: "b",
	}

	exists1 := mapnutil.Exists(m, 1)
	fmt.Printf("exists1: %v\n", exists1)

	exists2 := mapnutil.Exists(m, 2)
	fmt.Printf("exists2: %v\n", exists2)

	exists3 := mapnutil.Exists(m, 3)
	fmt.Printf("exists3: %v\n", exists3)

	// Output:
	// exists1: true
	// exists2: true
	// exists3: false
}
