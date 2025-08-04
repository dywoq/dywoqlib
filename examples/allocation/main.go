package main

import (
	"fmt"

	"github.com/dywoq/dywoqlib/allocation"
	"github.com/dywoq/dywoqlib/container/slice"
)

func grow(i int, s allocation.Sizer) {
	s.Grow(i)
}

func main() {
	// pre-allocate capacity
	d := slice.NewDynamic(1, 2, 3, 4)
	grow(10, d)
	// printing capacity
	native := d.Native()
	fmt.Printf("cap(native): %v\n", cap(native)) // cap(native): 10
}
