package ansi_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/console/ansi"
)

func ExampleApplyFg() {
	s := ansi.ApplyFg("Hi!", ansi.Blue)
	fmt.Printf("s: %v\n", s)
	// Output:
	// s: [34mHi![0m
}

func ExampleApplyBg() {
	s := ansi.ApplyBg("Hi!", ansi.Blue)
	fmt.Printf("s: %v\n", s)
	// Output:
	// s: [44mHi![0m
}

func ExampleApplyBoth() {
	s := ansi.ApplyBoth("Hi!", ansi.Green, ansi.Black)
	fmt.Printf("s: %v\n", s)
	// Output:
	// s: [32m[40mHi![0m
}
