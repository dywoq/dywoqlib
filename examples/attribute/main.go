package main

import (
	"fmt"

	"github.com/dywoq/dywoqlib/attribute"
)

// attribute package helps you mark code that needs attention,
// like deprecated functions, and provides a way to handle these events.

func main() {
	// if the custom event of attribute.Todo is nil - the function fallbacks to default event
	attribute.Todo(nil)

	// here is the same situation: no custom event = the function fallbacks to default event
	attribute.Deprecated(nil)

	// if we run the program, you would see the output:
	// attribute.Todo: todo in main.notImplemented; source: main.main
	// attribute.Deprecated: main.deprecated is deprecated; source: main.main

	// if you want to set the custom event, you need to pass the function.
	// attribute.Todo requires the function to not have parameters and return type.
	// custom events are useful especially if you want to specify which function developer need to use instead.
	attribute.Todo(func() { fmt.Println("still not implemented") })

	// we do the absolutely same thing as above
	attribute.Deprecated(func() { fmt.Println("deprecatedCustom() is deprecated, use deprecated()") })

	// now, after you set your custom events, you would see this output:
	// still not implemented
	// deprecatedCustom() is deprecated, use deprecated()
}
