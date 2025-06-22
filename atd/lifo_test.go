package atd

import (
	"testing"
)

func TestLifoBasic(t *testing.T) {
	l := NewLifo[int]()
	if !l.Empty() {
		t.Errorf("Empty() = false, want true")
	}
	l.Append(1)
	l.Append(2)
	l.Append(3)
	if l.Length() != 3 {
		t.Errorf("Length() = %d, want 3", l.Length())
	}
	if l.Top() != 3 {
		t.Errorf("Top() = %d, want 3", l.Top())
	}
	l.Pop()
	if l.Top() != 2 {
		t.Errorf("Top() = %d, want 2", l.Top())
	}
	l.Pop()
	l.Pop()
	if !l.Empty() {
		t.Errorf("Empty() = false, want true")
	}
}

func TestLifoErrorHandling(t *testing.T) {
	var l *Lifo[int]
	l = nil
	l = NewLifo[int]()
	l.data = nil
	l.Append(42)
	if l.Err() != ErrNilData {
		if l.Err() != nil {
			t.Errorf("Err() = %v, want %v", l.Err(), ErrNilData)
		} else {
			t.Errorf("Err() = nil, want %v", ErrNilData)
		}
	}
	l = NewLifo[int]()
	l.Pop() // Should not error
	if l.Err() != nil {
		t.Errorf("Err() is not nil: %v", l.Err())
	}
}
