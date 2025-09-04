package sliceutil_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/container"
	"github.com/dywoq/dywoqlib/sliceutil"
)

func ExampleFormat() {
	slice := []int{2, 3, 4}
	str, err := sliceutil.Format(slice)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
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
