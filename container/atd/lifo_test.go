package atd

import (
	"slices"
	"testing"
)

func TestLength(t *testing.T) {
	l := NewLifo[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)

	got := l.Length()
	want := 4

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestEmpty(t *testing.T) {
	l := NewLifo[int]()

	got := l.Empty()
	want := true

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestAppend(t *testing.T) {
	l := NewLifo[int]()
	l.Append(2)
	l.Append(4)

	got := l.Native()
	want := []int{2, 4}

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPop(t *testing.T) {
	l := NewLifo[int]()
	l.Append(2)
	l.Append(4)
	l.Pop()

	got := l.Native()
	want := []int{2}

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTop(t *testing.T) {
	l := NewLifo[int]()
	l.Append(2)
	l.Append(4)

	got := l.Top()
	want := 4

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestString(t *testing.T) {
	l := NewLifo[int]()
	l.Append(2)
	l.Append(4)

	got := l.String()
	want := "[2, 4]"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
