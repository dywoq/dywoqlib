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
