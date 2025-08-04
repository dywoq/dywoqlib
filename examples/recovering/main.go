package main

import "github.com/dywoq/dywoqlib/recovering"

// recovering package provides Recover() to recover from panics with logging.

func main() {
	// the thing that differs recovering.Recover() is logging.
	// the main goal of recovering.Recover() is making using recover() shorter, and print readable log message.
	// it uses log.Printf internally for logging.
	defer recovering.Recover()
	panic("something went wrong!!")
	// after you run the program, you would see something like this:
	// recovering.Recover(): caught the panic: something went wrong!!
	// also, recovering.Recover() returns the caught panic value if you need it.
}
