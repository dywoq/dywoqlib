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

package polymorph_test

import (
	"fmt"
	"reflect"

	"github.com/dywoq/dywoqlib/polymorph"
)

func ExampleComparable() {
	isComparable := polymorph.Comparable[int]()
	fmt.Printf("isComparable: %v\n", isComparable)

	isComparable = polymorph.Comparable[[]int]()
	fmt.Printf("isComparable: %v\n", isComparable)
	// Output:
	// isComparable: true
	// isComparable: false
}

type test struct {
	a bool
}

func (t test) Method() bool   { return t.a }
func (t test) String() string { return fmt.Sprintf("%v", t.a) }

func ExampleHasMethod() {
	has := polymorph.HasMethod[test]("Method")
	fmt.Printf("has: %v\n", has)

	has = polymorph.HasMethod[test]("MethodA")
	fmt.Printf("has: %v\n", has)

	// Output:
	// has: true
	// has: false
}

func ExampleHasField() {
	has := polymorph.HasField[test]("a")
	fmt.Printf("has: %v\n", has)

	has = polymorph.HasField[test]("b")
	fmt.Printf("has: %v\n", has)

	// Output:
	// has: true
	// has: false
}

func ExampleImplements() {
	has := polymorph.Implements[fmt.Stringer, test]()
	fmt.Printf("has: %v\n", has)

	has = polymorph.Implements[fmt.Formatter, test]()
	fmt.Printf("has: %v\n", has)

	// Output:
	// has: true
	// has: false
}

func ExampleKindOf() {
	kind := polymorph.KindOf[int]()
	fmt.Printf("%v == %v: %v", kind.String(), reflect.Int, kind == reflect.Int)

	// Output:
	// int == int: true
}

func ExampleNillable() {
	nillable := polymorph.Nillable[chan int]()
	fmt.Printf("nillable: %v\n", nillable)

	// Output:
	// nillable: true
}

func ExampleNumMethods() {
	num := polymorph.NumMethods[test]()
	fmt.Printf("num: %v\n", num)

	// Output:
	// num: 2
}

func ExampleNumFields() {
	num := polymorph.NumFields[test]()
	fmt.Printf("num: %v\n", num)

	// Output:
	// num: 1
}

func ExamplePackagePath() {
	path := polymorph.PackagePath[fmt.Stringer]()
	fmt.Printf("path: %v\n", path)

	// Output:
	// path: fmt
}

func ExampleTypeOfGeneric() {
	typeof := polymorph.TypeOfGeneric[int]()
	fmt.Printf("typeof: %v\n", typeof)

	// Output:
	// typeof: int
}

func ExampleIs() {
	val := int(20)
	fmt.Printf("polymorph.Is[int](val): %v\n", polymorph.Is[int](val))

	// Output:
	// polymorph.Is[int](val): true
}
