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
	it := iterator.NewReserve(slice)
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
