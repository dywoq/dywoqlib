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

func ExampleFormat() {
	m := map[int]string{
		1: "a",
		2: "b",
	}

	emptyM := map[int]string{}

	formatted, _ := mapnutil.Format(m)
	fmt.Printf("formatted: %v\n", formatted)

	emptyFormatted, _ := mapnutil.Format(emptyM)
	fmt.Printf("emptyFormatted: %v\n", emptyFormatted)

	// Output:
	// formatted: {
	//   1: a
	//   2: b
	// }
	// emptyFormatted:
}
