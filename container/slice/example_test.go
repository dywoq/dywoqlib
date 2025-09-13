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

package slice_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/container/slice"
)

func ExampleDynamic() {
	d := slice.NewDynamic(2, 3)
	fmt.Printf("d: %v\n", d)

	length := d.Length()
	fmt.Printf("length: %v\n", length)

	it := d.Iterating().Forward()
	for it.Next() {
		fmt.Printf("iterator.Value(): %v\n", it.Value())
	}

	d.Append(4, 5)
	fmt.Printf("d: %v\n", d)

	first := d.At(0)
	fmt.Printf("first: %v\n", first)

	found := d.Find(2)
	fmt.Printf("found: %v\n", found)

	d.Set(3, 0)
	fmt.Printf("d: %v\n", d)

	d.Delete(3)
	fmt.Printf("d: %v\n", d)

	d.Insert(2, 3)
	fmt.Printf("d: %v\n", d)

	front := d.Front()
	fmt.Printf("front: %v\n", front)

	back := d.Back()
	fmt.Printf("back: %v\n", back)

	d.Pop()
	fmt.Printf("d: %v\n", d)

	// Output:
	// d: [2, 3]
	// length: 2
	// iterator.Value(): 2
	// iterator.Value(): 3
	// d: [2, 3, 4, 5]
	// first: 2
	// found: 2
	// d: [3, 3, 4, 5]
	// d: [3, 3, 4, 0]
	// d: [3, 3, 3, 4, 0]
	// front: 3
	// back: 0
	//d: [3, 3, 3, 4]
}

func ExampleFixed() {
	d := slice.NewFixed(10, 2, 3)
	fmt.Printf("d: %v\n", d)

	length := d.Length()
	fmt.Printf("length: %v\n", length)

	it := d.Iterating().Forward()
	for it.Next() {
		fmt.Printf("iterator.Value(): %v\n", it.Value())
	}

	d.Append(4, 5)
	fmt.Printf("d: %v\n", d)

	first := d.At(0)
	fmt.Printf("first: %v\n", first)

	found := d.Find(2)
	fmt.Printf("found: %v\n", found)

	d.Set(3, 0)
	fmt.Printf("d: %v\n", d)

	d.Delete(3)
	fmt.Printf("d: %v\n", d)

	d.Insert(2, 3)
	fmt.Printf("d: %v\n", d)

	front := d.Front()
	fmt.Printf("front: %v\n", front)

	back := d.Back()
	fmt.Printf("back: %v\n", back)

	d.Pop()
	fmt.Printf("d: %v\n", d)

	// Output:
	// d: [2, 3]
	// length: 2
	// iterator.Value(): 2
	// iterator.Value(): 3
	// d: [2, 3, 4, 5]
	// first: 2
	// found: 2
	// d: [3, 3, 4, 5]
	// d: [3, 3, 4, 0]
	// d: [3, 3, 3, 4, 0]
	// front: 3
	// back: 0
	//d: [3, 3, 3, 4]
}
