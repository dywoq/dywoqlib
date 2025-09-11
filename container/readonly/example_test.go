package readonly_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/container/readonly"
)

func ExampleSlice() {
	s := readonly.NewSlice(2, 3, 4, 5)

	// length getting
	length := s.Length()
	fmt.Printf("length: %v\n", length)

	// getting element at index
	at := s.At(0)
	fmt.Printf("at: %v\n", at)

	// finding element
	found := s.Find(3)
	fmt.Printf("found: %v\n", found)

	// into string
	fmt.Printf("s: %v\n", s)

	// front
	front := s.Front()
	fmt.Printf("front: %v\n", front)

	// back
	back := s.Back()
	fmt.Printf("back: %v\n", back)

	// error checking
	if !s.Error().Nil() {
		fmt.Printf("s.Error(): %v\n", s.Error())
	}

	// Output:
	// length: 4
	// at: 2
	// found: 3
	// s: [2, 3, 4, 5]
	// front: 2
	// back: 5
}
