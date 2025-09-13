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

package unique_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/container/unique"
)

func ExampleSlice() {
	s := unique.NewSlice(2, 2, 3)
	fmt.Printf("s: %v\n", s)

	length := s.Length()
	fmt.Printf("length: %v\n", length)

	it := s.Iterating().Forward()
	for it.Next() {
		fmt.Printf("it.Value(): %v\n", it.Value())
	}

	s.Append(3, 5)
	fmt.Printf("s: %v\n", s)

	at := s.At(0)
	fmt.Printf("at: %v\n", at)

	find := s.Find(2)
	fmt.Printf("find: %v\n", find)

	s.Set(8, 0)
	fmt.Printf("s: %v\n", s)

	s.Delete(1)
	fmt.Printf("s: %v\n", s)

	s.Insert(0, 9)
	fmt.Printf("s: %v\n", s)

	front := s.Front()
	fmt.Printf("front: %v\n", front)

	back := s.Back()
	fmt.Printf("back: %v\n", back)

	s.Pop()
	fmt.Printf("s: %v\n", s)

	// Output:
	// s: [2, 3]
	// length: 2
	// it.Value(): 2
	// it.Value(): 3
	// s: [2, 3, 5]
	// at: 2
	// find: 2
	// s: [8, 3, 5]
	// s: [8, 5, 0]
	// s: [9, 8, 5, 0]
	// front: 9
	// back: 0
	// s: [9, 8, 5, 0]
}

func ExampleLifo() {
	l := unique.NewLifo[int]()
	l.Append(2)
	l.Append(2)
	fmt.Printf("l: %v\n", l)

	empty := l.Empty()
	fmt.Printf("empty: %v\n", empty)

	length := l.Length()
	fmt.Printf("length: %v\n", length)

	l.Pop()
	fmt.Printf("l: %v\n", l)

	l.Append(30)

	top := l.Top()
	fmt.Printf("top: %v\n", top)
	// Output:
	// l: [2]
	// empty: false
	// length: 1
	// l: []
	// top: 30
}

func ExampleFifo() {
	l := unique.NewFifo[int]()
	l.Append(2)
	l.Append(2)
	fmt.Printf("l: %v\n", l)

	empty := l.Empty()
	fmt.Printf("empty: %v\n", empty)

	length := l.Length()
	fmt.Printf("length: %v\n", length)

	l.Pop()
	fmt.Printf("l: %v\n", l)

	l.Append(30)
	l.Append(50)

	front := l.Front()
	fmt.Printf("top: %v\n", front)

	back := l.Back()
	fmt.Printf("back: %v\n", back)

	// Output:
	// l: [2]
	// empty: false
	// length: 1
	// l: []
	// top: 30
	// back: 50
}
