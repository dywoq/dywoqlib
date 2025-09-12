package numeric

import (
	"fmt"
	"math"

	"github.com/dywoq/dywoqlib/err"
	"github.com/dywoq/dywoqlib/numeric/base"
	"github.com/dywoq/dywoqlib/numeric/constraints"
	"github.com/dywoq/dywoqlib/numeric/limits"
)

// Base represents a wrapper around integer, which allows you to work with integers safely.
type Base[I constraints.Integral] interface {
	// Get returns the underlying integer.
	Get() I

	// Error returns the possible encountered error,
	// otherwise it would simply return err.NoneContext(.)
	Error() err.Context

	// Set sets the underlying integer to val and returns the modified version.
	// If error is present, the function is skipped.
	Set(val I) Base[I]

	// Add adds val to the underlying integer and returns the modified version.
	// If error is present, the function is skipped.
	Add(val I) Base[I]

	// Minus takes val from the underlying integer and returns the modified version.
	// If error is present, the function is skipped.
	Minus(val I) Base[I]

	// Divide divides the underlying integer by val and returns the modified version.
	// The operation is skipped if val is zero.
	// If error is present, the function is skipped.
	Divide(val I) Base[I]

	// Multiply multiplies the underlying integer by val and returns the modified version.
	// If error is present, the function is skipped.
	Multiply(val I) Base[I]

	// Limits returns a pair of two values, which are
	// the minimum and maximum values of the chosen integer type I.
	Limits() (I, I)

	// Zero reports whether the underlying integer is zero.
	// If error is present, the function is skipped.
	Zero() bool

	// Negative reports whether the underlying integer is negative.
	// If error is present, the function is skipped.
	Negative() bool

	// Positive reports whether the underlying integer is positive.
	// If error is present, the function is skipped.
	Positive() bool

	// Equal checks if the underlying integer is equal to val.
	// If error is present, the function is skipped.
	Equal(val I) bool

	// Compare returns the result of comparison
	// between val and the underlying integer.
	//
	// If the integer is greater than val, the function returns base.CompareGreater.
	//
	// If the integer is less than val,  the function returns base.CompareLess.
	//
	// If they're equal, the function returns base.CompareEqual.
	//
	// If error is present, the function is skipped.
	// If none of these conditions satisfied, Compare returns zero.
	Compare(val I) base.Compare

	// Odd reports whether the underlying integer is odd.
	// If error is present, the function is skipped.
	Odd() bool

	// Even reports whether the underlying integer is even.
	// If error is present, the function is skipped.
	Even() bool

	// Absolute returns the absolute value of the underlying integer.
	// If error is present, the function is skipped.
	Absolute() I

	// String returns the string representation of the underlying integer.
	// If error is present, the function is skipped.
	String() string

	// And preforms a bitwise AND operation with val and returns the modified version.
	// If error is present, the function is skipped.
	And(val I) Base[I]

	// Or performs a bitwise OR operation with val and returns the modified version.
	// If error is present, the function is skipped.
	Or(val I) Base[I]

	// Xor performs a bitwise XOR operation with val and returns the modified version.
	// If error is present, the function is skipped.
	Xor(val I) Base[I]

	// ShiftLeft performs a bitwise left shift by n positions and returns the modified version.
	// If error is present, the function is skipped.
	ShiftLeft(n I) Base[I]

	// ShiftRight performs a bitwise right shift by n positions and returns the modified version.
	// If error is present, the function is skipped.
	ShiftRight(n I) Base[I]

	// Sign returns the sign of the underlying number.
	//
	// If the number is negative, it returns base.SignNegative.
	//
	// If the number is positive, it returns base.SignPositive.
	//
	// If the number is zero, it returns base.SignZero.
	//
	// If error is present, the function is skipped.
	Sign() base.Sign

	// Prime reports whether the underlying integer is a prime number.
	// If error is present, the function is skipped.
	Prime() bool
}

// Int returns Base[int] integer.
func Int(val ...int) Base[int] { return baseFactoryMethod(val...) }

// Int8 returns Base[int8] integer.
func Int8(val ...int8) Base[int8] { return baseFactoryMethod(val...) }

// Int16 returns Base[int16] integer.
func Int16(val ...int16) Base[int16] { return baseFactoryMethod(val...) }

// Int32 returns Base[int32] integer.
func Int32(val ...int32) Base[int32] { return baseFactoryMethod(val...) }

// Int64 returns Base[int64] integer.
func Int64(val ...int64) Base[int64] { return baseFactoryMethod(val...) }

// UInt returns Base[uint] integer.
func UInt(val ...uint) Base[uint] { return baseFactoryMethod(val...) }

// UInt8 returns Base[uint8] integer.
func UInt8(val ...uint8) Base[uint8] { return baseFactoryMethod(val...) }

// UInt8 returns Base[uint16] integer.
func UInt16(val ...uint16) Base[uint16] { return baseFactoryMethod(val...) }

// UInt32 returns Base[uint32] integer.
func UInt32(val ...uint32) Base[uint32] { return baseFactoryMethod(val...) }

// UInt64 returns Base[uint64] integer.
func UInt64(val ...uint64) Base[uint64] { return baseFactoryMethod(val...) }

type baseImplementation[I constraints.Integral] struct {
	num I
	err err.Context
}

func (b baseImplementation[I]) Absolute() I {
	if b.num < 0 {
		return -b.num
	}
	return b.num
}

func (b baseImplementation[I]) Add(val I) Base[I] {
	if !b.err.Nil() {
		return b.none()
	}
	return b.new(b.num + val)
}

func (b baseImplementation[I]) And(val I) Base[I] {
	if !b.err.Nil() {
		return b.none()
	}
	return b.new(b.num & val)
}

func (b baseImplementation[I]) Compare(val I) base.Compare {
	switch {
	case !b.err.Nil():
		return 0
	case b.num > val:
		return base.CompareGreater
	case b.num < val:
		return base.CompareLess
	case b.num == val:
		return base.CompareEqual
	default:
		return 0
	}
}

func (b baseImplementation[I]) Divide(val I) Base[I] {
	if !b.err.Nil() {
		return b.none()
	}
	if val == 0 {
		return b.none()
	}
	return b.new(b.num / val)
}

func (b baseImplementation[I]) Equal(val I) bool {
	if !b.err.Nil() {
		return false
	}
	return b.num == 0
}

func (b baseImplementation[I]) Error() err.Context {
	return b.err
}

func (b baseImplementation[I]) Even() bool {
	return !(b.num%2 == 0)
}

func (b baseImplementation[I]) Get() I {
	return b.num
}

func (b baseImplementation[I]) Limits() (I, I) {
	if !b.err.Nil() {
		return 0, 0
	}
	return limits.Numeric[I]()
}

func (b baseImplementation[I]) Minus(val I) Base[I] {
	if !b.err.Nil() {
		return b.none()
	}
	return b.new(b.num - val)
}

func (b baseImplementation[I]) Multiply(val I) Base[I] {
	if !b.err.Nil() {
		return b.none()
	}
	return b.new(b.num * val)
}

func (b baseImplementation[I]) Negative() bool {
	if !b.err.Nil() {
		return false
	}
	return b.num < 0
}

func (b baseImplementation[I]) Odd() bool {
	return b.num%2 == 0
}

func (b baseImplementation[I]) Or(val I) Base[I] {
	if !b.err.Nil() {
		return b.none()
	}
	return b.new(b.num | val)
}

func (b baseImplementation[I]) Positive() bool {
	if !b.err.Nil() {
		return false
	}
	return b.num > 0
}

func (b baseImplementation[I]) Prime() bool {
	switch {
	case b.num <= 1:
		return false
	case b.num == 2:
		return true
	case b.num%2 == 0:
		return false
	}

	limit := I(math.Sqrt(float64(b.num)))
	for i := I(3); i <= limit; i += 2 {
		if b.num%i == 0 {
			return false
		}
	}
	return true
}

func (b baseImplementation[I]) Set(val I) Base[I] {
	if !b.err.Nil() {
		return b.none()
	}
	return b.new(val)
}

func (b baseImplementation[I]) ShiftLeft(n I) Base[I] {
	if !b.err.Nil() {
		return b.none()
	}
	if n < 0 {
		return b.none()
	}
	return b.new(b.num << n)
}

func (b baseImplementation[I]) ShiftRight(n I) Base[I] {
	if !b.err.Nil() {
		return b.none()
	}
	if n < 0 {
		return b.none()
	}
	return b.new(b.num >> n)
}

func (b baseImplementation[I]) Sign() base.Sign {
	switch {
	case !b.err.Nil():
		return 0
	case b.num < 0:
		return base.SignNegative
	case b.num > 0:
		return base.SignPositive
	case b.num == 0:
		return base.SignZero
	default:
		return 0
	}
}

func (b baseImplementation[I]) String() string {
	if !b.err.Nil() {
		return ""
	}
	return fmt.Sprint(b.num)
}

func (b baseImplementation[I]) Xor(val I) Base[I] {
	if !b.err.Nil() {
		return b.none()
	}
	return b.new(b.num ^ val)
}

func (b baseImplementation[I]) Zero() bool {
	if !b.err.Nil() {
		return false
	}
	return b.num == 0
}

func (b *baseImplementation[I]) none() Base[I] {
	return baseImplementation[I]{0, err.NoneContext()}
}

func (b *baseImplementation[I]) new(val I) Base[I] {
	return baseImplementation[I]{val, err.NoneContext()}
}

func baseFactoryMethod[I constraints.Integral](val ...I) Base[I] {
	if len(val) != 0 {
		return baseImplementation[I]{val[0], err.NoneContext()}
	}
	return baseImplementation[I]{0, err.NoneContext()}
}
