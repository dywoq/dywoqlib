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

package mapn_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/container/mapn"
)

func ExampleDynamic() {
	d := mapn.NewDynamic(map[string]int{
		"a": 10,
		"b": 20,
	})

	// getting length
	length := d.Length()
	fmt.Printf("length: %v\n", length)

	// exists
	exists := d.Exists("a")
	fmt.Printf("exists: %v\n", exists)

	// adding
	d.Add("d", 3)

	// setting
	d.Set("a", 40)

	// Deleting
	d.Delete("a")

	// getting key
	gotkey, gotvalue := d.Get("d")
	fmt.Printf("gotkey: %v, gotvalue: %v\n", gotkey, gotvalue)

	// Output:
	// length: 2
	// exists: true
	// gotkey: d, gotvalue: 3
}

func ExampleFixed() {
	d := mapn.NewFixed(4, map[string]int{
		"a": 10,
		"b": 20,
	})

	// getting length
	length := d.Length()
	fmt.Printf("length: %v\n", length)

	// exists
	exists := d.Exists("a")
	fmt.Printf("exists: %v\n", exists)

	// adding
	d.Add("d", 3)

	// setting
	d.Set("a", 40)

	// deleting
	d.Delete("a")

	// getting key
	gotkey, gotvalue := d.Get("d")
	fmt.Printf("gotkey: %v, gotvalue: %v\n", gotkey, gotvalue)

	// Output:
	// length: 2
	// exists: true
	// gotkey: d, gotvalue: 3
}
