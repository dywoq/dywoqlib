package attribute_test

import "github.com/dywoq/dywoqlib/lib/attribute"

func hi() {
	attribute.Deprecated(nil)
}

func ExampleDeprecated() {
	hi()
	// Output:
	// attribute.Deprecated: github.com/dywoq/dywoqlib/lib/attribute_test.hi is deprecated; source: github.com/dywoq/dywoqlib/lib/attribute_test.ExampleDeprecated
}
