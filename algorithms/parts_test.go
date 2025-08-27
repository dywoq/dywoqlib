package algorithms_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/algorithms"
)

func ExampleParts() {
	parts := algorithms.Parts(100, 4)
	for _, num := range parts {
		fmt.Printf("num: %v\n", num)
	}
	// Output:
	// num: 20
	// num: 40
	// num: 60
	// num: 80
}
