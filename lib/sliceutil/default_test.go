package sliceutil

import (
	"errors"
	"slices"
	"testing"

	"github.com/dywoq/dywoqlib/lib/iterator"
)

func TestFormat(t *testing.T) {
	var (
		x = []int{10, 11, 44}
		y = []int{}
	)

	// first test
	got, _ := Format(x)
	want := "[10, 11, 44]"
	if got != want {
		t.Errorf("Format([]int{10, 11, 44}) = %s, want %s", got, want)
	}

	// second test
	got, _ = Format(y)
	want = "[]"
	if got != want {
		t.Errorf("Format([]int{}) = %s, want %s", got, want)
	}
}

// we could use container.IterableSlice, but container already uses the sliceutil package,
// that would cause a circular dependency.
type iteratableSlice[T comparable] []T

func (it iteratableSlice[T]) Iterating() *iterator.Combined[T] { return iterator.NewCombined(it) }

func TestFind(t *testing.T) {
	x := iteratableSlice[int]{10, 11, 44}

	// first test
	got, _ := Find(10, x.Iterating().Forward())
	want := 10
	if got != want {
		t.Errorf("Find(10, x.Iterating().Forward()) = %d, want %d", got, want)
	}

	// second test
	_, err := Find(55, x.Iterating().Forward())
	errwant := ErrElementNotFound
	if !errors.Is(err, errwant) {
		t.Errorf("Find(55, x.Iterating().Forward()) = %s, want %s", err.Error(), errwant.Error())
	}
}

func TestAt(t *testing.T) {
	x := []int{10, 11, 44}

	// first test
	got, _ := At(1, x)
	want := 11
	if got != want {
		t.Errorf("At(1, x) = %d, want %d", got, want)
	}

	// second test
	_, err := At(20, x)
	errwant := ErrWrongIndex
	if !errors.Is(err, errwant) {
		t.Errorf("At(1, x) = %s, want %s", err.Error(), errwant.Error())
	}
}

func TestSet(t *testing.T) {
	// first test
	x := []int{10, 11, 44}
	Set(10, 1, x)
	want := []int{10, 10, 44}
	if !slices.Equal(x, want) {
		t.Errorf("Set(10, 1, x) = %v, want %v", x, want)
	}

	// second test
	x = []int{10, 11, 44}
	_, err := Set(10, 1100, x)
	errwant := ErrWrongIndex
	if !errors.Is(err, errwant) {
		t.Errorf("Set(10, 1100, x) = %s, want %s", err, errwant)
	}
}

func TestDelete(t *testing.T) {
	// first test
	x := []int{10, 11, 44}
	Delete(0, x)
	want := []int{11, 44, 0}
	if !slices.Equal(x, want) {
		t.Errorf("Delete(0, x) = %v, want %v", x, want)
	}

	// second test
	x = []int{10, 11, 44}
	_, err := Delete(1000, x)
	errwant := ErrWrongIndex
	if !errors.Is(err, errwant) {
		t.Errorf("Delete(0, x) = %s, want %s", err.Error(), errwant.Error())
	}
}

func TestInsert(t *testing.T) {
	// first test
	x := []int{10, 11, 44}
	Insert(1, &x, 54)
	want := []int{10, 54, 11, 44}
	if !slices.Equal(x, want) {
		t.Errorf("Insert(1, x, 54) = %v, want %v", x, want)
	}

	// second test
	x = []int{10, 11, 44}
	_, err := Insert(1000, &x, 54)
	errwant := ErrWrongIndex
	if !errors.Is(err, errwant) {
		t.Errorf("Insert(1000, x, 54) = %s, want %s", err.Error(), errwant.Error())
	}
}
