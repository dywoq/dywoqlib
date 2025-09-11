package mapn_test

import (
	"fmt"

	"github.com/dywoq/dywoqlib/container/mapn"
)

func ExampleDynamic() {
	d := mapn.NewDynamic(map[string]int{
		"a": 10,
		"b": 20,
	})

	// getting length
	length := d.Length()
	fmt.Printf("length: %v\n", length)

	// exists
	exists := d.Exists("a")
	fmt.Printf("exists: %v\n", exists)

	// adding
	d.Add("d", 3)

	// setting
	d.Set("a", 40)

	// Deleting
	d.Delete("a")

	// getting key
	gotkey, gotvalue := d.Get("d")
	fmt.Printf("gotkey: %v, gotvalue: %v\n", gotkey, gotvalue)

	// Output:
	// length: 2
	// exists: true
	// gotkey: d, gotvalue: 3
}

func ExampleFixed() {
	d := mapn.NewFixed(4, map[string]int{
		"a": 10,
		"b": 20,
	})

	// getting length
	length := d.Length()
	fmt.Printf("length: %v\n", length)

	// exists
	exists := d.Exists("a")
	fmt.Printf("exists: %v\n", exists)

	// adding
	d.Add("d", 3)

	// setting
	d.Set("a", 40)

	// deleting
	d.Delete("a")

	// getting key
	gotkey, gotvalue := d.Get("d")
	fmt.Printf("gotkey: %v, gotvalue: %v\n", gotkey, gotvalue)

	// Output:
	// length: 2
	// exists: true
	// gotkey: d, gotvalue: 3
}
