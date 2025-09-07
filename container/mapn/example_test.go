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

	fmt.Printf("d: %v\n", d)

	length := d.Length()
	fmt.Printf("length: %v\n", length)

	exists := d.Exists("a")
	fmt.Printf("exists: %v\n", exists)

	d.Add("d", 3)
	fmt.Printf("d: %v\n", d)

	d.Set("a", 40)
	fmt.Printf("d: %v\n", d)

	keys := d.Keys()
	fmt.Printf("keys: %v\n", keys)

	values := d.Values()
	fmt.Printf("values: %v\n", values)

	d.Delete("a")
	fmt.Printf("d: %v\n", d)

	gotkey, gotvalue := d.Get("d")
	fmt.Printf("gotkey: %v, gotvalue: %v\n", gotkey, gotvalue)

	// Output:
	// d: {
	//   a: 10
	//   b: 20
	// }
	// length: 2
	// exists: true
	// d: {
	//   a: 10
	//   b: 20
	//   d: 3
	// }
	// d: {
	//   a: 10
	//   b: 20
	//   d: 3
	// }
	// keys: [a b d]
	// values: [10 20 3]
	// d: {
	//   b: 20
	//   d: 3
	// }
	// gotkey: d, gotvalue: 3
}

func ExampleFixed() {
	d := mapn.NewFixed(4, map[string]int{
		"a": 10,
		"b": 20,
	})

	fmt.Printf("d: %v\n", d)

	length := d.Length()
	fmt.Printf("length: %v\n", length)

	exists := d.Exists("a")
	fmt.Printf("exists: %v\n", exists)

	d.Add("d", 3)
	fmt.Printf("d: %v\n", d)

	d.Set("a", 40)
	fmt.Printf("d: %v\n", d)

	keys := d.Keys()
	fmt.Printf("keys: %v\n", keys)

	values := d.Values()
	fmt.Printf("values: %v\n", values)

	d.Delete("a")
	fmt.Printf("d: %v\n", d)

	gotkey, gotvalue := d.Get("d")
	fmt.Printf("gotkey: %v, gotvalue: %v\n", gotkey, gotvalue)

	// Output:
	// d: {
	//   a: 10
	//   b: 20
	// }
	// length: 2
	// exists: true
	// d: {
	//   a: 10
	//   b: 20
	//   d: 3
	// }
	// d: {
	//   a: 10
	//   b: 20
	//   d: 3
	// }
	// keys: [a b d]
	// values: [10 20 3]
	// d: {
	//   b: 20
	//   d: 3
	// }
	// gotkey: d, gotvalue: 3
}
