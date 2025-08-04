#### `allocation.Sizer`

Each struct in dywoqlib who support pre-allocation (such as `slice.Dynamic`, `mapn.Dynamic`, `stringn.String`) do implement a specific interface - `allocation.Sizer`. This interface contains method ```Grow(int)``` for managing memory capacity.

The interface itself:
```go
type Sizer interface {
	Grow(i int)
}
```

First example:
```go
package main

import (
	"github.com/dywoq/dywoqlib/allocation"
	"github.com/dywoq/dywoqlib/container/slice"
)

func grow(i int, sizer allocation.Sizer) {
	sizer.Grow(i)
}

func main() {
	d := slice.NewDynamic(1, 2, 3, 3)
	grow(10, d)
}
```

Second example:
```go
package main

import (
	"fmt"
	"github.com/dywoq/dywoqlib/allocation"
	"github.com/dywoq/dywoqlib/container/slice"
)

// growAndShowCapacity shows the capacity before and after calling Grow.
func growAndShowCapacity(n int, sizer allocation.Sizer) {
	if d, ok := sizer.(*slice.Dynamic); ok {
		fmt.Printf("Before Grow(%d): Len=%d, Cap=%d\n", n, d.Len(), d.Cap())
	}

	sizer.Grow(n)

	if d, ok := sizer.(*slice.Dynamic); ok {
		fmt.Printf("After Grow(%d): Len=%d, Cap=%d\n", n, d.Len(), d.Cap())
	}
}

func main() {
	d := slice.NewDynamic(1, 2, 3, 4)
	growAndShowCapacity(10, d)
}
```
