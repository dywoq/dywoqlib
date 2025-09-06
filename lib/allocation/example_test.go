package allocation_test

import (
	"github.com/dywoq/dywoqlib/lib/allocation"
	"github.com/dywoq/dywoqlib/lib/stringn"
)

func ExampleSizer() {
	grow := func(i int, s allocation.Sizer) {
		s.Grow(i)
	}
	str := stringn.New("Hi!")
	grow(10, str)
}
