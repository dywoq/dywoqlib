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

	// retrieving values (you can get keys too by d.Keys())
	values := d.Values()
	fmt.Printf("values: %v\n", values)

	// Deleting
	d.Delete("a")

	// getting key
	gotkey, gotvalue := d.Get("d")
	fmt.Printf("gotkey: %v, gotvalue: %v\n", gotkey, gotvalue)

	// Output:
	// length: 2
	// exists: true
	// values: [10 20 3]
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

	// retrieving values (you can get keys too by d.Keys())
	values := d.Values()
	fmt.Printf("values: %v\n", values)

	// deleting
	d.Delete("a")

	// getting key
	gotkey, gotvalue := d.Get("d")
	fmt.Printf("gotkey: %v, gotvalue: %v\n", gotkey, gotvalue)

	// Output:
	// length: 2
	// exists: true
	// values: [0 0 0 40 20 3]
	// gotkey: d, gotvalue: 3
}
