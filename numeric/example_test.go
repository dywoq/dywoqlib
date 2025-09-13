package numeric_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/numeric"
)

func ExampleBase() {
	num := numeric.Int(32)

	// setting
	num = num.Set(23)
	fmt.Printf("num: %v\n", num)

	// adding
	num = num.Add(2)
	fmt.Printf("num: %v\n", num)

	// minus
	num = num.Minus(5)
	fmt.Printf("num: %v\n", num)

	// dividing
	num = num.Divide(4)
	fmt.Printf("num: %v\n", num)

	// multiplying
	num = num.Multiply(4)
	fmt.Printf("num: %v\n", num)

	// limits
	min, max := num.Limits()
	fmt.Printf("min: %v\n", min)
	fmt.Printf("max: %v\n", max)

	// zero checking
	zero := num.Zero()
	fmt.Printf("zero: %v\n", zero)

	// negative checking
	negative := num.Negative()
	fmt.Printf("negative: %v\n", negative)

	// positive checking
	positive := num.Positive()
	fmt.Printf("positive: %v\n", positive)

	// equal checking
	equal := num.Equal(32)
	fmt.Printf("equal: %v\n", equal)

	// comparison
	comparison := num.Compare(32)
	fmt.Printf("comparison: %v\n", comparison)

	// odd, even
	odd := num.Odd()
	even := num.Even()
	fmt.Printf("odd: %v\n", odd)
	fmt.Printf("even: %v\n", even)

	// absolute
	abs := num.Absolute()
	fmt.Printf("abs: %v\n", abs)

	// and bitwise
	and := num.And(3)
	fmt.Printf("and: %v\n", and)

	// or bitwise
	or := num.Or(5)
	fmt.Printf("or: %v\n", or)

	// xor bitwise
	xor := num.Xor(5)
	fmt.Printf("xor: %v\n", xor)

	// shifting left/right
	shiftLeft := num.ShiftLeft(3)
	shiftRight := num.ShiftRight(3)
	fmt.Printf("shiftLeft: %v\n", shiftLeft)
	fmt.Printf("shiftRight: %v\n", shiftRight)

	// sign
	sign := num.Sign()
	fmt.Printf("sign: %v\n", sign)

	// prime checking
	prime := num.Prime()
	fmt.Printf("prime: %v\n", prime)

	// Output:
	// num: 23
	// num: 25
	// num: 20
	// num: 5
	// num: 20
	// min: -9223372036854775808
	// max: 9223372036854775807
	// zero: false
	// negative: false
	// positive: true
	// equal: false
	// comparison: -1
	// odd: false
	// even: true
	// abs: 20
	// and: 0
	// or: 21
	// xor: 17
	// shiftLeft: 160
	// shiftRight: 2
	// sign: 1
	// prime: false
}
