package atd

import (
	"testing"
)

func TestFifoBasic(t *testing.T) {
	f := NewFifo[int]()
	if !f.Empty() {
		t.Errorf("Empty() = false, want true")
	}
	f.Append(1, 2, 3)
	if f.Length() != 3 {
		t.Errorf("Length() = %d, want 3", f.Length())
	}
	if f.Front() != 1 {
		t.Errorf("Front() = %d, want 1", f.Front())
	}
	if f.Back() != 3 {
		t.Errorf("Back() = %d, want 3", f.Back())
	}
	f.Pop()
	if f.Front() != 2 {
		t.Errorf("Front() = %d, want 2", f.Front())
	}
	f.Pop()
	f.Pop()
	if !f.Empty() {
		t.Errorf("Empty() = false, want true")
	}
}

func TestFifoErrorHandling(t *testing.T) {
	var f *Fifo[int]
	f = nil
	f = NewFifo[int]()
	f.data = nil
	f.Append(42)
	if f.Err() != ErrNilData {
		if f.Err() != nil {
			t.Errorf("Err() = %v, want %v", f.Err(), ErrNilData)
		} else {
			t.Errorf("Err() = nil, want %v", ErrNilData)
		}
	}
	f = NewFifo[int]()
	f.Pop() // Should not error
	if f.Err() != nil {
		t.Errorf("Err() is not nil: %v", f.Err())
	}
}
