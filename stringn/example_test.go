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

package stringn_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/stringn"
)

func ExampleString() {
	str := stringn.New("Hi, Go!")

	fmt.Printf("length: %v\n", str.Length())

	it := str.Iterating().Forward()
	for it.Next() {
		v := string(it.Value())
		if v != " " {
			fmt.Printf("it.Value(): %v\n", v)
		} else {
			fmt.Printf("it.Value():\n")
		}
	}

	str.Append(" Bye, Go!")

	fmt.Printf("r: %v\n", str.At(0))
	fmt.Printf("front: %v\n", str.Front())
	fmt.Printf("back: %v\n", str.Back())
	fmt.Printf("empty: %v\n", str.Empty())

	fmt.Printf("hasRunePrefix: %v\n", str.HasRunePrefix('H'))
	fmt.Printf("hasStringPrefix: %v\n", str.HasStringPrefix("Hi"))
	fmt.Printf("hasRuneSuffix: %v\n", str.HasRuneSuffix('!'))
	fmt.Printf("hasStringSuffix: %v\n", str.HasStringSuffix("Go!"))

	str.Insert(0, 'H')
	fmt.Printf("str: %v\n", str)

	str.Set('J', 0)
	fmt.Printf("str: %v\n", str)

	fmt.Printf("containsRune: %v\n", str.ContainsRune('o'))
	fmt.Printf("containsString: %v\n", str.ContainsString("Go!"))

	str.Write([]byte("Hi!"))

	str.Clear()
	str.Prepend("Hi! Bye!")
	fmt.Printf("str: %v\n", str)

	str.RemoveRange(0, 1)
	fmt.Printf("str: %v\n", str)

	str.Replace("b", "y")
	fmt.Printf("str: %v\n", str)

	str.ToLower()
	fmt.Printf("lower: %v\n", str)
	str.ToUpper()
	fmt.Printf("upper: %v\n", str)
	fmt.Printf("compare: %v\n", str.Compare("Hi!"))
	fmt.Printf("equals: %v\n", str.Equals("Bye!"))
	fmt.Printf("split: %v\n", str.Split("y"))
	fmt.Printf("substr: %v\n", str.Substring(0, 1))

	// Output:
	// length: 7
	// it.Value(): H
	// it.Value(): i
	// it.Value(): ,
	// it.Value():
	// it.Value(): G
	// it.Value(): o
	// it.Value(): !
	// r: 72
	// front: 72
	// back: 33
	// empty: false
	// hasRunePrefix: true
	// hasStringPrefix: true
	// hasRuneSuffix: true
	// hasStringSuffix: true
	// str: HHi, Go! Bye, Go!
	// str: JHi, Go! Bye, Go!
	// containsRune: true
	// containsString: true
	// str: Hi! Bye!
	// str: i! Bye!
	// str: i! Bye!
	// lower: i! bye!
	// upper: I! BYE!
	// compare: 1
	// equals: false
	// split: [I! BYE!]
	// substr: I
}
