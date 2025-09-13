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

package atd_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/container/atd"
)

func ExampleLifo() {
	l := atd.NewLifo[int]()
	l.Append(2)
	l.Append(3)
	fmt.Printf("l: %v\n", l)

	l.Pop()
	fmt.Printf("l: %v\n", l)

	r := l.Top()
	fmt.Printf("r: %v\n", r)

	length := l.Length()
	fmt.Printf("length: %v\n", length)

	empty := l.Empty()
	fmt.Printf("empty: %v\n", empty)

	// Output:
	// l: [2, 3]
	// l: [2]
	// r: 2
	// length: 1
	// empty: false
}

func ExampleFifo() {
	f := atd.NewFifo[int]()
	f.Append(2)
	f.Append(6)
	fmt.Printf("f: %v\n", f)

	f.Pop()
	fmt.Printf("f: %v\n", f)

	front := f.Front()
	fmt.Printf("front: %v\n", front)

	back := f.Back()
	fmt.Printf("back: %v\n", back)

	length := f.Length()
	fmt.Printf("length: %v\n", length)

	empty := f.Empty()
	fmt.Printf("empty: %v\n", empty)

	// Output:
	// f: [2, 6]
	// f: [2]
	// front: 2
	// back: 2
	// length: 1
	//empty: false
}
