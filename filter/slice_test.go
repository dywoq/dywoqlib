package filter_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/dywoq/dywoqlib/filter"
)

func TestSlice(t *testing.T) {
	tests := []struct {
		slice []int
		pred  func(int) bool
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, func(i int) bool { return i%2 == 0 }, []int{2, 4, 6}},
		{[]int{1, 2, 3, 4, 5, 6}, func(i int) bool { return i%2 != 0 }, []int{1, 3, 5}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.slice), func(t *testing.T) {
			got := filter.Slice(test.slice, test.pred)
			if !slices.Equal(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestSliceNot(t *testing.T) {
	tests := []struct {
		slice []int
		pred  func(int) bool
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, func(i int) bool { return i%2 == 0 }, []int{1, 3, 5}},
		{[]int{1, 2, 3, 4, 5, 6}, func(i int) bool { return i%2 != 0 }, []int{2, 4, 6}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.slice), func(t *testing.T) {
			got := filter.SliceNot(test.slice, test.pred)
			if !slices.Equal(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func BenchmarkSliceEven(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	for i := 0; i < b.N; i++ {
		_ = filter.Slice(data, func(i int) bool { return i%2 == 0 })
	}
}

func BenchmarkSliceOdd(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	for i := 0; i < b.N; i++ {
		_ = filter.Slice(data, func(i int) bool { return i%2 != 0 })
	}
}

func BenchmarkSliceNotEven(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	for i := 0; i < b.N; i++ {
		_ = filter.SliceNot(data, func(i int) bool { return i%2 == 0 })
	}
}

func BenchmarkSliceNotOdd(b *testing.B) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}
	for i := 0; i < b.N; i++ {
		_ = filter.SliceNot(data, func(i int) bool { return i%2 != 0 })
	}
}

func ExampleSlice() {
	data := []int{1, 2, 3, 4, 5, 6}
	even := filter.Slice(data, func(i int) bool { return i%2 == 0 })
	fmt.Println(even)
	// Output: [2 4 6]
}

func ExampleSlice_odd() {
	data := []int{1, 2, 3, 4, 5, 6}
	odd := filter.Slice(data, func(i int) bool { return i%2 != 0 })
	fmt.Println(odd)
	// Output: [1 3 5]
}

func ExampleSliceNot() {
	data := []int{1, 2, 3, 4, 5, 6}
	notEven := filter.SliceNot(data, func(i int) bool { return i%2 == 0 })
	fmt.Println(notEven)
	// Output: [1 3 5]
}

func ExampleSliceNot_odd() {
	data := []int{1, 2, 3, 4, 5, 6}
	notOdd := filter.SliceNot(data, func(i int) bool { return i%2 != 0 })
	fmt.Println(notOdd)
	// Output: [2 4 6]
}

