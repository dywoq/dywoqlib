package main

import (
	"fmt"

	"github.com/dywoq/dywoqlib/optional"
)

// optional package provides a clear and safe way to handle values that may or may not be present,
// without using nil pointers or zero values.

func main() {
	// optional.Maybe[T] is a generic structure, meaning it can work with any type.
	// to create a new optional.Maybe structure, you need to use optional.New[T any]():
	_ = optional.New(2)
	// you don't need to provide the type, optional.New automatically decides the type,
	// judging to the value you pass to the function.

	// if you want to pass an empty optional.Maybe value, you would use optional.None[T any]().
	// unlike optional.New, it requires the generic type T for zero value.
	_ = optional.None[int]()

	// however, for people who think this is too tiring to write optional.None[T]() every time,
	// you can use conversion methods, which include builtin types.
	// these methods are indentical to each other, for example: optional.Int8(val...int8).
	// if val arguments weren't provided, it returns optional.None[T]().
	// if they were, it returns optional.New(val[0]).
	_ = optional.Int8()   // returns optional.None[T]()
	_ = optional.Int32(1) // returns optional.New(val[0])

	// to get an optional value, you use optional.Maybe.Get().
	// the function returns two values: the optional value and the boolean.
	// the boolean always match what Present() returns:
	opt := optional.New(2)
	val, ok := opt.Get()
	if ok {
		fmt.Printf("we got the value!! it's %v: %d\n", ok, val)
	}
	fmt.Printf("val.Present() == ok: %v\n", opt.Present() == ok)

	// to set a default value if optional value is not present,
	// you would use if-else statement:
	opt2 := optional.String()
	if !opt2.Present() {
		fmt.Println("hi, guest!")
	} else {
		got, _ := opt2.Get()
		fmt.Printf("hi, %s!\n", got)
	}
	// but this may be waste of time to write if-else statements.
	// so, optional.Maybe has a solution: optional.Maybe[T].Else(T) T:
	opt2 = optional.String("dywoq")
	name := opt2.Else("guest")
	fmt.Printf("bye, %s!\n", name)
	// here, opt2.Else() returns either optional value if it's present, or default value set by you.

	// optional.Maybe.Filter(func(T) bool) returns Maybe[T] if optional value is present, and
	// it satisfies your predicate function. otherwise, it returns an empty Maybe.
	// it's useful when you want to avoid deeply nested if statements.
	res := opt2.Filter(func(s string) bool { return len(s) == 5 })
	if res.Present() {
		fmt.Printf("res: %v\n", res)
	} else {
		fmt.Printf("res: there's no res :*(\n")
	}
}
