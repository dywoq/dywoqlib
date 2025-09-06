package ansi_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/lib/console/ansi"
)

func ExampleBase() {
	msg := ansi.New("Hi!")
	msg = msg.SetBgColor(ansi.Green)
	fmt.Printf("msg: %v\n", msg)

	msg = msg.SetFgColor(ansi.Cyan)
	fmt.Printf("msg: %v\n", msg)

	// Output:
	// msg: [310m[42mHi![0m
	// msg: [36m[42m[310m[42mHi![0m[0m
}
