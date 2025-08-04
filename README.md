# dywoqlib

`dywoqlib` is a comprehensive Go library designed to provide a collection of utility packages that enhance and simplify common programming tasks. It includes generic data structures, iterators, optional types, and various helper functions for slices, strings, and more.

## Installation

To use `dywoqlib` in your project, you can install it using `go get`:

```sh
go get github.com/dywoq/dywoqlib
```

## Features

*   **Generic Data Structures**: Type-safe data structures like queues, stacks, dynamic maps, and slices.
*   **Fluent APIs**: Many packages offer a fluent and chainable API, with error handling built-in.
*   **Iterators**: A flexible iterator pattern for traversing collections.
*   **Optional Type**: An `Optional` type to handle nullable values gracefully.
*   **Polymorphism Utilities**: Helpers for working with Go generics at runtime.
*   **Slice and String Utilities**: A rich set of functions for manipulating slices and strings.
*   **Console Helper**: A simple way to run external commands.
*   **Panic Recovery**: A utility to safely recover from panics with logging.

## Package Overview

Here is a detailed overview of the packages available in `dywoqlib`.

### allocation

The `allocation` package provides a `Sizer` interface, which is used by data structures to manage their memory capacity.

### attribute

The `attribute` package is used for internal purposes, such as getting information about the caller function.

### console

The `console` package and its subpackages provides a functions to work with external commands and console.

#### `console/ansi`

This package provides a clean and structured way to manage ANSI-colored strings in Go by using an interface `Base`.  

**Example:**

```go
package main

import (
	"fmt"

	"github.com/dywoq/dywoqlib/console/ansi"
)

func main() {
	msg := ansi.New("Hi!").SetBgColor(ansi.Black).SetBgColor(ansi.Cyan)
	fmt.Printf("msg: %v\n", msg)
}
```

### container

The `container` package and its subpackages provide various generic data structures.

#### `container/atd`

This package offers abstract data types like `Fifo` (queue) and `Lifo` (stack).

**Example (`Fifo`):**

```go
package main

import (
	"fmt"
	"github.com/dywoq/dywoqlib/container/atd"
)

func main() {
	q := atd.NewFifo[int]()
	q.Append(1)
	q.Append(2)
	fmt.Println(q.Pop()) // 1
	fmt.Println(q.Pop()) // 2
}
```

#### `container/mapn`

This package provides `Dynamic` and `Fixed` map implementations with a fluent API.

**Example (`Dynamic`):**

```go
package main

import (
	"fmt"
	"github.com/dywoq/dywoqlib/container/mapn"
)

func main() {
	m := mapn.NewDynamic(make(map[string]int))
	m.Add("one", 1)
	m.Add("two", 2)

	val, _ := m.Get("one")
	fmt.Println(val) // 1
}
```

#### `container/slice`

This package provides `Dynamic` and `Fixed` slice implementations, as well as wrappers like `IterableSlice`.

**Example (`Dynamic`):**

```go
package main

import (
	"fmt"
	"github.com/dywoq/dywoqlib/container/slice"
)

func main() {
	s := slice.NewDynamic[string]()
	s.Append("hello")
	s.Append("world")
	fmt.Println(s.At(0)) // "hello"
}
```

### iterator

The `iterator` package provides a generic iterator interface and implementations for forward, reverse, and combined iteration.

**Example:**

```go
package main

import (
	"fmt"
	"github.com/dywoq/dywoqlib/container"
)

func main() {
	s := container.IterableSlice[int]{1, 2, 3}
	it := s.Iterating()

	for it.Next() {
		fmt.Printf("Index: %d, Value: %d\n", it.Position(), it.Value())
	}
}
```

### optional

The `optional` package provides a type-safe `Maybe` that can be used to represent values that may or may not be present.

**Example:**

```go
package main

import (
	"fmt"
	"github.com/dywoq/dywoqlib/optional"
)

func main() {
	opt1 := optional.New("hello")
	fmt.Println(opt1.Present()) // true
	fmt.Println(opt1.Get())     // "hello", true

	opt2 := optional.None[string]()
	fmt.Println(opt2.Present())   // false
	fmt.Println(opt2.Else("default")) // "default"
}
```

### polymorph

The `polymorph` package contains utilities for working with Go generics at runtime.

**Example:**

```go
package main

import (
	"fmt"
	"github.com/dywoq/dywoqlib/polymorph"
)

func main() {
	fmt.Println(polymorph.Comparable[int]())    // true
	fmt.Println(polymorph.Comparable[[]int]()) // false
}
```

### recovering

The `recovering` package provides a function to recover from panics while logging the panic message.

**Example:**

```go
package main

import (
	"fmt"
	"github.com/dywoq/dywoqlib/recovering"
)

func main() {
	defer recovering.Recover();
	panic("something went wrong")
}
```

### sliceutil

The `sliceutil` package offers a collection of helper functions for working with slices.

**Example:**

```go
package main

import (
	"fmt"
	"github.com/dywoq/dywoqlib/sliceutil"
)

func main() {
	s := []int{1, 2, 3}
	val, err := sliceutil.At(1, s)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(val) // 2
}
```

### stringn

The `stringn` package provides a mutable `String` type with a rich API for string manipulation.

**Example:**

```go
package main

import (
	"fmt"
	"github.com/dywoq/dywoqlib/stringn"
)

func main() {
	s := stringn.New("hello")
	s.Append(" world")
	s.Replace("world", "Go")
	fmt.Println(s.String()) // "hello Go"
}
```

## License

`dywoqlib` is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for more details.
