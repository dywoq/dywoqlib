package integer

import (
	"github.com/dywoq/dywoqlib/numeric/base"
	"github.com/dywoq/dywoqlib/numeric/constraints"
)

// Base defines an interface for generic integer types with arithmetic, bitwise,
// comparison, and utility operations.
type Base[I constraints.Integral] interface {
	// Native returns the underlying Go integer value.
	Native() I
	// String returns the string representation of the integer.
	String() string
	// Add returns a new Base with the sum of the current value and the given value.
	Add(I) Base[I]
	// Minus returns a new Base with the difference between the current value and the given value.
	Minus(I) Base[I]
	// Divide returns a new Base with the result of integer division by the given value.
	Divide(I) Base[I]
	// Multiply returns a new Base with the product of the current value and the given value.
	Multiply(I) Base[I]
	// Shift provides left and right bitwise shift operations.
	Shift() interface {
		// Left shifts the integer left by the given number of bits.
		Left(I) Base[I]
		// Right shifts the integer right by the given number of bits.
		Right(I) Base[I]
	}
	// Power returns the current value raised to the power of the given exponent.
	Power(I) Base[I]
	// Absolute returns the absolute value of the integer.
	Absolute(I) Base[I]
	// Mod returns the remainder after division by the given value.
	Mod(I) Base[I]
	// Limits returns the minimum and maximum possible values for the integer type.
	Limits() (I, I)
	// Even returns true if the integer is even.
	Even() bool
	// Compare returns a comparison result with another integer.
	Compare() base.Compare
	// Sign returns the sign of the integer (negative, zero, or positive).
	Sign() base.Sign
	// Sqrt returns the integer square root and a boolean indicating success.
	Sqrt() (Base[I], bool)
	// Bitwise provides standard bitwise operations.
	Bitwise() interface {
		// And returns the bitwise AND of the integer with another.
		And(I) Base[I]
		// Or returns the bitwise OR of the integer with another.
		Or(I) Base[I]
		// Xor returns the bitwise XOR of the integer with another.
		Xor(I) Base[I]
		// Not returns the bitwise NOT of the integer.
		Not(I) Base[I]
		// CountBits returns the number of set bits in the integer.
		CountBits() Base[I]
		// CountTraillingZero returns the number of trailing zero bits.
		CountTraillingZero() Base[I]
	}
	// Bytes returns the byte representation of the integer.
	Bytes() []byte
}
