package limits_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/numeric/limits"
)

func ExampleNumeric() {
	min, max := limits.Numeric[int64]()
	fmt.Printf("min: %v\n", min)
	fmt.Printf("max: %v\n", max)

	// Output:
	// min: -9223372036854775808
	// max: 9223372036854775807
}
