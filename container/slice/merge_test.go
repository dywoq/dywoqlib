package slice

import (
	"slices"
	"testing"
)

func TestMergeDynamic(t *testing.T) {
	first := NewDynamic(1, 2)
	second := NewDynamic(3, 4)
	got := MergeDynamic(first, second)
	expected := []int{1, 2, 3, 4}
	if !slices.Equal(got.Native(), expected) {
		t.Errorf("MergeDynamic(first, second) = %v, %v", got.Native(), expected)
	}
}

func TestMergeFixed(t *testing.T) {
	first := NewFixed(10, 1, 2)
	if first.Error() != nil {
		t.Fatal(first.Error())
	}

	second := NewFixed(10, 3, 4)
	if second.Error() != nil {
		t.Fatal(second.Error())
	}

	got := MergeFixed(first, second)
	if got.Error() != nil {
		t.Fatal(got.Error())
	}

	expected := []int{1, 2, 3, 4}
	if !slices.Equal(got.Native(), expected) {
		t.Errorf("MergeFixed(first, second) = %v, %v", got.Native(), expected)
	}
}
