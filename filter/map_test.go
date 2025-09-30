package filter_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/dywoq/dywoqlib/filter"
)

func TestMap(t *testing.T) {
	tests := []struct {
		m    map[string]int
		pred func(string, int) bool
		want map[string]int
	}{
		{map[string]int{"a": 2, "b": 3, "c": 4}, func(s string, i int) bool { return i%2 == 0 }, map[string]int{"a": 2, "c": 4}},
		{map[string]int{"a": 2, "b": 3, "c": 4}, func(s string, i int) bool { return i%2 != 0 }, map[string]int{"b": 3}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.m), func(t *testing.T) {
			got := filter.Map(test.m, test.pred)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestMapNot(t *testing.T) {
	tests := []struct {
		m    map[string]int
		pred func(string, int) bool
		want map[string]int
	}{
		{map[string]int{"a": 2, "b": 3, "c": 4}, func(s string, i int) bool { return i%2 == 0 }, map[string]int{"b": 3}},
		{map[string]int{"a": 2, "b": 3, "c": 4}, func(s string, i int) bool { return i%2 != 0 }, map[string]int{"a": 2, "c": 4}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.m), func(t *testing.T) {
			got := filter.MapNot(test.m, test.pred)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestMapKeys(t *testing.T) {
	tests := []struct {
		m    map[string]int
		pred func(string) bool
		want map[string]int
	}{
		{map[string]int{"a": 2, "bs": 3, "c": 4}, func(s string) bool { return len(s) == 1 }, map[string]int{"a": 2, "c": 4}},
		{map[string]int{"a": 2, "bs": 3, "c": 4}, func(s string) bool { return len(s) != 1 }, map[string]int{"bs": 3}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.m), func(t *testing.T) {
			got := filter.MapKeys(test.m, test.pred)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestMapValues(t *testing.T) {
	tests := []struct {
		m    map[string]int
		pred func(int) bool
		want map[string]int
	}{
		{map[string]int{"a": 2, "bs": 3, "c": 4}, func(i int) bool { return i%2 == 0 }, map[string]int{"a": 2, "c": 4}},
		{map[string]int{"a": 2, "bs": 3, "c": 4}, func(i int) bool { return i%2 != 0 }, map[string]int{"bs": 3}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.m), func(t *testing.T) {
			got := filter.MapValues(test.m, test.pred)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func BenchmarkMap(b *testing.B) {
	m := map[string]int{}
	for i := range 1000 {
		m[fmt.Sprintf("k%d", i)] = i
	}
	pred := func(_ string, v int) bool { return v%2 == 0 }

	for b.Loop() {
		_ = filter.Map(m, pred)
	}
}

func BenchmarkMapNot(b *testing.B) {
	m := map[string]int{}
	for i := range 1000 {
		m[fmt.Sprintf("k%d", i)] = i
	}
	pred := func(_ string, v int) bool { return v%2 == 0 }

	for b.Loop() {
		_ = filter.MapNot(m, pred)
	}
}

func BenchmarkMapKeys(b *testing.B) {
	m := map[string]int{}
	for i := range 1000 {
		m[fmt.Sprintf("k%d", i)] = i
	}
	pred := func(k string) bool { return len(k) > 2 }

	for b.Loop() {
		_ = filter.MapKeys(m, pred)
	}
}

func BenchmarkMapValues(b *testing.B) {
	m := map[string]int{}
	for i := range 1000 {
		m[fmt.Sprintf("k%d", i)] = i
	}
	pred := func(v int) bool { return v%2 == 0 }

	for b.Loop() {
		_ = filter.MapValues(m, pred)
	}
}

func ExampleMap() {
	m := map[string]int{"a": 2, "b": 3, "c": 4}
	even := func(_ string, v int) bool { return v%2 == 0 }

	got := filter.Map(m, even)
	fmt.Println(got)
	// Output: map[a:2 c:4]
}

func ExampleMapNot() {
	m := map[string]int{"a": 2, "b": 3, "c": 4}
	even := func(_ string, v int) bool { return v%2 == 0 }

	got := filter.MapNot(m, even)
	fmt.Println(got)
	// Output: map[b:3]
}

func ExampleMapKeys() {
	m := map[string]int{"a": 2, "bs": 3, "c": 4}
	oneChar := func(k string) bool { return len(k) == 1 }

	got := filter.MapKeys(m, oneChar)
	fmt.Println(got)
	// Output: map[a:2 c:4]
}

func ExampleMapValues() {
	m := map[string]int{"a": 2, "bs": 3, "c": 4}
	even := func(v int) bool { return v%2 == 0 }

	got := filter.MapValues(m, even)
	fmt.Println(got)
	// Output: map[a:2 c:4]
}
