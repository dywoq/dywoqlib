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

func ExampleGcd() {
	a, b := float64(4), float64(4)
	divisor := algorithms.Gcd(a, b)
	fmt.Printf("The GCD of %f and %f is %f\n", a, b, divisor)
	// Output:
	// The GCD of 4.000000 and 4.000000 is 4.000000
}
