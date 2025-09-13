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

package limits_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/numeric/limits"
)

func ExampleNumeric() {
	min, max := limits.Numeric[int64]()
	fmt.Printf("min: %v\n", min)
	fmt.Printf("max: %v\n", max)

	// Output:
	// min: -9223372036854775808
	// max: 9223372036854775807
}
