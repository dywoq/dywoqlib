package allocation_test

import (
	"github.com/dywoq/dywoqlib/allocation"
	"github.com/dywoq/dywoqlib/stringn"
)

func ExampleSizer() {
	grow := func(i int, s allocation.Sizer) {
		s.Grow(i)
	}
	str := stringn.New("Hi!")
	grow(10, str)
}
