package sliceutil

import (
	"slices"
	"testing"

	"github.com/dywoq/dywoqlib/container"
)

func TestFormat(t *testing.T) {
	got1, err := Format([]int{2, 3})
	if err != nil {
		t.Fatal(got1)
	}
	want1 := "[2, 3]"
	if got1 != want1 {
		t.Errorf("Format([]int{2, 3}) = %s, want %s", got1, want1)
	}

	got2, err := Format([]int{})
	if err != nil {
		t.Fatal(got2)
	}
	want2 := "[]"
	if got2 != want2 {
		t.Errorf("Format([]int{}) = %s, want %s", got1, want1)
	}
}

func TestFind(t *testing.T) {
	s := container.IterableSlice[int]{2, 3, 4}
	it := s.Iterating().Forward()
	got, _ := Find(3, it)
	want := 3
	if got != want {
		t.Errorf("Find(2, it) = %v, want %v", got, want)
	}

	_, err := Find(5, it)
	if err != ErrElementNotFound {
		t.Errorf("Find(5, it) is expected to send error %s", ErrElementNotFound.Error())
	}
}

func TestAt(t *testing.T) {
	s := container.IterableSlice[int]{2, 3, 4}
	got, _ := At(2, s)
	want := 4
	if got != want {
		t.Errorf("At(2, s) = %v, want %v", got, want)
	}

	_, err := At(1, s)
	if err != ErrWrongIndex {
		t.Errorf("At(1, s) is expected to send error %v", ErrWrongIndex)
	}
}

func TestSet(t *testing.T) {
	s := []int{2, 3, 4}
	Set(5, 0, s)
	if !slices.Equal(s, []int{5, 3, 4}) {
		t.Errorf("slices.Equal(s, []int{5, 3, 4}) = false, want true")
	}

	_, err := Set(5, 10, s)
	if err != ErrWrongIndex {
		t.Errorf("Set(5, 10, s) is expected to send error %v", ErrWrongIndex)
	}
}

func TestDelete(t *testing.T) {
	s := []int{2, 3, 4}
	Delete(0, s)
	if !slices.Equal(s, []int{3, 4}) {
		t.Errorf("slices.Equal(s, []int{3, 4}) = false, want true")
	}

	_, err := Delete(0, s)
	if err != ErrWrongIndex {
		t.Errorf("Delete(0, s) is expected to send error %v", ErrWrongIndex)
	}
}
func TestInsert(t *testing.T) {
	s := []int{2, 3, 4}
	Insert(0, 5, s)
	if !slices.Equal(s, []int{5, 2, 3, 4}) {
		t.Errorf("s[0] is not 3")
	}

	_, err := Insert(10, 5, s)
	if err != ErrWrongIndex {
		t.Errorf("Insert(10, 5, s) is expected to send error %v", ErrWrongIndex)
	}
}
