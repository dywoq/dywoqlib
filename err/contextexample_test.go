package err_test

import (
	"errors"
	"fmt"

	"github.com/dywoq/dywoqlib/err"
)

func ExampleContext() {
	c := err.NewContext(errors.New("something wrong"), "source is example context")
	fmt.Printf("c: %v\n", c)

	json, errjson := c.Marshal()
	if errjson != nil {
		fmt.Printf("errjson: %v\n", errjson)
	}
	fmt.Printf("json: %v\n", string(json))

	c.SetError(nil)
	c.SetMore("")
	fmt.Printf("c: %v\n", c)

	anotherC := err.NewContext(errors.New("two something wrong"), "source is example context")
	c.Copy(anotherC)

	fmt.Printf("c: %v\n", c)

	// Output:
	// c: something wrong: source is example context
	// json: {"error":"something wrong","more":"source is example context"}
	// c: <nil>
	// c: two something wrong: source is example context
}
