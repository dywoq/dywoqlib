package algorithms_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/lib/algorithms"
)

func ExampleGcd() {
	a, b := float64(4), float64(4)
	divisor := algorithms.Gcd(a, b)
	fmt.Printf("The GCD of %f and %f is %f\n", a, b, divisor)
	// Output:
	// The GCD of 4.000000 and 4.000000 is 4.000000
}
