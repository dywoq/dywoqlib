package attribute_test

import "github.com/dywoq/dywoqlib/attribute"

func hi() {
	attribute.Deprecated(nil)
}

func ExampleDeprecated() {
	hi()
	// Output:
	// attribute.Deprecated: github.com/dywoq/dywoqlib/attribute_test.hi is deprecated; source: github.com/dywoq/dywoqlib/attribute_test.ExampleDeprecated
}
