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
