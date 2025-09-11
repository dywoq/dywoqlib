package mapnutil_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/mapnutil"
)

func ExampleExists() {
	m := map[int]string{
		1: "a",
		2: "b",
	}

	exists1 := mapnutil.Exists(m, 1)
	fmt.Printf("exists1: %v\n", exists1)

	exists2 := mapnutil.Exists(m, 2)
	fmt.Printf("exists2: %v\n", exists2)

	exists3 := mapnutil.Exists(m, 3)
	fmt.Printf("exists3: %v\n", exists3)

	// Output:
	// exists1: true
	// exists2: true
	// exists3: false
}
