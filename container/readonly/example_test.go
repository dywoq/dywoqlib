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

func ExampleMap() {
	m := readonly.NewMap(map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	})

	// length getting
	length := m.Length()
	fmt.Printf("length: %v\n", length)

	// exists
	exists := m.Exists("one")
	fmt.Printf("exists: %v\n", exists)

	// getting values (you can retrieve keys too by m.Keys())
	values := m.Values()
	fmt.Printf("values: %v\n", values)

	// getting
	gotkey, gotvalue := m.Get("one")
	fmt.Printf("gotkey: %v\n", gotkey)
	fmt.Printf("gotvalue: %v\n", gotvalue)

	// Output:
	// length: 4
	// exists: true
	// values: [1 2 3 4]
	// gotkey: one
	// gotvalue: 1
}
