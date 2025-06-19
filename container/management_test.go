package container

import (
	"testing"

	"github.com/dywoq/dywoqlib/iterator"
)

type slice[T any] struct {
	data []T
}

func (s slice[T]) Begin() *iterator.Iterator[T] {
	return iterator.New(0, s.data)
}

func (s slice[T]) End() *iterator.Iterator[T] {
	return iterator.New(len(s.data)-1, s.data)
}

func TestFind(t *testing.T) {
	s := slice[int]{[]int{2, 3}}
	m := management[int]{s, nil}

	got := m.find(3)
	expected := 3
	if m.currentErr() != nil {
		t.Errorf("management.find[int](3) got error %s", m.currentErr())
	}

	if got != expected {
		t.Errorf("management.find[int](3) = %d, want %d", got, expected)
	}
}

func TestAt(t *testing.T) {
	s := slice[int]{[]int{2, 3}}
	m := management[int]{s, nil}

	got := m.at(0)
	expected := 2
	if m.currentErr() != nil {
		t.Errorf("management.at[int](0) got error %s", m.currentErr())
	}

	if got != expected {
		t.Errorf("management.at[int](0) = %d, want %d", got, expected)
	}
}
