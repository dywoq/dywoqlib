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

package sliceutil_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/container"
	"github.com/dywoq/dywoqlib/sliceutil"
)

func ExampleFormat() {
	slice := []int{2, 3, 4}
	str := sliceutil.Format(slice)
	fmt.Printf("str: %v\n", str)
	// Output:
	// str: [2, 3, 4]
}

func ExampleFind() {
	slice := container.IterableSlice[int]{2, 3, 4}
	elem, err := sliceutil.Find(2, slice.Iterating().Forward())
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("elem: %v\n", elem)
	// Output:
	// elem: 2
}

func ExampleAt() {
	slice := []int{2, 3, 4}
	elem, err := sliceutil.At(1, slice)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("elem: %v\n", elem)
	// Output:
	// elem: 3
}

func ExampleSet() {
	slice := []int{2, 3, 4}
	_, err := sliceutil.Set(5, 0, slice)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("slice: %v\n", slice)
	// Output:
	// slice: [5 3 4]
}

func ExampleDelete() {
	slice := []int{2, 3, 4}
	_, err := sliceutil.Delete(0, slice)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("slice: %v\n", slice)
	// Output:
	// slice: [3 4 0]
}

func ExampleInsert() {
	slice := []int{2, 3, 4}
	_, err := sliceutil.Insert(1, &slice, 1)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("slice: %v\n", slice)
	// Output:
	// slice: [2 1 3 4]
}

func ExamplePop() {
	slice := []int{2, 3, 4}
	sliceutil.Pop(&slice)
	fmt.Printf("slice: %v\n", slice)
	// Output:
	// slice: [2 3]
}

func ExampleFront() {
	slice := []int{2, 3, 4}
	front := sliceutil.Front(slice)
	fmt.Printf("front: %v\n", front)
	// Output:
	// front: 2
}

func ExampleBack() {
	slice := []int{2, 3, 4}
	back := sliceutil.Back(slice)
	fmt.Printf("back: %v\n", back)
	// Output:
	// back: 4
}
