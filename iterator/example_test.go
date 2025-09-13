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

package iterator_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/iterator"
)

func ExampleForward() {
	slice := []int{2, 3, 4}
	it := iterator.NewForward(slice)
	for it.Next() {
		fmt.Printf("it.Value(): %v\n", it.Value())
	}

	it.Reset()

	for it.Next() {
		ptr := it.ValuePtr()
		*ptr *= 2
	}
	fmt.Printf("slice: %v\n", slice)

	// Output:
	// it.Value(): 2
	// it.Value(): 3
	// it.Value(): 4
	//slice: [2 6 8]
}

func ExampleReverse() {
	slice := []int{2, 3, 4}
	it := iterator.NewReverse(slice)
	for it.Next() {
		fmt.Printf("it.Value(): %v\n", it.Value())
	}

	it.Reset()

	for it.Next() {
		ptr := it.ValuePtr()
		*ptr *= 2
	}
	fmt.Printf("slice: %v\n", slice)

	// Output:
	// it.Value(): 4
	// it.Value(): 3
	// it.Value(): 2
	// slice: [4 6 8]
}

func ExampleReadonlyBase() {
	slice := []int{2, 3, 4}
	it := iterator.ReadonlyBase[int](iterator.NewForward(slice))
	for it.Next() {
		fmt.Printf("it.Value(): %v\n", it.Value())
	}

	// Output:
	// it.Value(): 2
	// it.Value(): 3
	// it.Value(): 4
}
