package attribute_test

import "github.com/dywoq/dywoqlib/lib/attribute"

func todo() {
	attribute.Todo(nil)
}

func ExampleTodo() {
	todo()
	// Output:
	// attribute.Todo: todo in github.com/dywoq/dywoqlib/lib/attribute_test.todo; source: github.com/dywoq/dywoqlib/lib/attribute_test.ExampleTodo
}
