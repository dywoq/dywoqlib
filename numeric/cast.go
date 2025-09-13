package numeric

import (
	"fmt"
	"github.com/dywoq/dywoqlib/err"
	"github.com/dywoq/dywoqlib/numeric/constraints"
	"github.com/dywoq/dywoqlib/numeric/limits"
)

// Cast safely converts b to Base[To] integer. First, it checks if b fits the
// limits of integer To. If b doesn't, it returns zero number, and error.
// If b does, Cast converts b into Base[To].
func Cast[From, To constraints.Integral](b Base[From]) (Base[To], err.Context) {
	val := b.Get()
	minTo, maxTo := limits.Numeric[To]()

	if isSigned[From]() {
		if isSigned[To]() {
			if int64(val) < int64(minTo) || int64(val) > int64(maxTo) {
				return baseFactoryMethod[To](), err.NewContext(ErrOverflow, fmt.Sprintf("%d overflows %T", val, *new(To)))
			}
		} else {
			if val < 0 {
				return baseFactoryMethod[To](), err.NewContext(ErrOverflow, fmt.Sprintf("%d overflows %T", val, *new(To)))
			}
			if uint64(val) > uint64(maxTo) {
				return baseFactoryMethod[To](), err.NewContext(ErrOverflow, fmt.Sprintf("%d overflows %T", val, *new(To)))
			}
		}
	} else {
		if isSigned[To]() {
			if uint64(val) > uint64(maxTo) {
				return baseFactoryMethod[To](), err.NewContext(ErrOverflow, fmt.Sprintf("%d overflows %T", val, *new(To)))
			}
		} else {
			if uint64(val) > uint64(maxTo) {
				return baseFactoryMethod[To](), err.NewContext(ErrOverflow, fmt.Sprintf("%d overflows %T", val, *new(To)))
			}
		}
	}

	return baseFactoryMethod(To(val)), nil
}

func isSigned[T constraints.Integral]() bool {
	return T(0)-T(1) < T(0)
}
