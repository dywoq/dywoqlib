package atd

import (
	"slices"
	"testing"
)

func TestFifoLength(t *testing.T) {
	l := NewFifo[int]()
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

func TestFifoEmpty(t *testing.T) {
	l := NewFifo[int]()

	got := l.Empty()
	want := true

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFifoAppend(t *testing.T) {
	l := NewFifo[int]()
	l.Append(2)
	l.Append(4)

	got := l.Native()
	want := []int{2, 4}

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFifoPop(t *testing.T) {
	l := NewFifo[int]()
	l.Append(2)
	l.Append(4)
	l.Pop()

	got := l.Native()
	want := []int{2}

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFifoString(t *testing.T) {
	l := NewFifo[int]()
	l.Append(2)
	l.Append(4)

	got := l.String()
	want := "[2, 4]"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFifoFront(t *testing.T) {
	f := NewFifo[int]()
	f.Append(2)
	f.Append(3)

	tests := []struct {
		body      string
		got, want bool
	}{
		{"f.Front() == 2", f.Front() == 2, true},
		{"f.Front() == 3", f.Front() == 3, false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func TestFifoBack(t *testing.T) {
	f := NewFifo[int]()
	f.Append(2)
	f.Append(3)

	tests := []struct {
		body      string
		got, want bool
	}{
		{"f.Back() == 3", f.Back() == 3, true},
		{"f.Back() == 2", f.Back() == 2, false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}
