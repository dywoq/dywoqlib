package attribute_test

import "github.com/dywoq/dywoqlib/attribute"

func todo() {
	attribute.Todo(nil)
}

func ExampleTodo() {
	todo()
	// Output:
	// attribute.Todo: todo in github.com/dywoq/dywoqlib/attribute_test.todo; source: github.com/dywoq/dywoqlib/attribute_test.ExampleTodo
}
