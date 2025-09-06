package limits

import (
	"math"

	"github.com/dywoq/dywoqlib/lib/numeric/constraints"
)

// Limits returns the minimal, and maximum limit of a generic parameter numeric I.
// Rarely, but the function panics if I is not supported type.
func Numeric[I constraints.Numeric]() (I, I) {
	var zero I
	switch any(zero).(type) {
	case int:
		min, max := int(math.MinInt), int(math.MaxInt)
		return I(min), I(max)
	case int8:
		min, max := int8(math.MinInt8), int8(math.MaxInt8)
		return I(min), I(max)
	case int16:
		min, max := int16(math.MinInt16), int16(math.MaxInt16)
		return I(min), I(max)
	case int32:
		min, max := int32(math.MinInt32), int32(math.MaxInt32)
		return I(min), I(max)
	case int64:
		min, max := int64(math.MinInt64), int64(math.MaxInt64)
		return I(min), I(max)
	case uint:
		min, max := uint(0), uint(math.MaxUint)
		return I(min), I(max)
	case uint8:
		min, max := uint8(0), uint8(math.MaxUint8)
		return I(min), I(max)
	case uint16:
		min, max := uint16(0), uint16(math.MaxUint16)
		return I(min), I(max)
	case uint32:
		min, max := uint32(0), uint32(math.MaxUint32)
		return I(min), I(max)
	case uint64:
		min, max := uint64(0), uint64(math.MaxUint64)
		return I(min), I(max)
	case float32:
		min, max := float32(-math.MaxFloat32), float32(math.MaxFloat32)
		return I(min), I(max)
	case float64:
		min, max := float64(-math.MaxFloat64), float64(math.MaxFloat64)
		return I(min), I(max)
	default:
		panic(ErrUnsupportedNumericType)
	}
}
