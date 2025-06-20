package slice

import "testing"

var d = NewDynamic(2, 3, 4)

func TestFront(t *testing.T) {
	got := d.Front()
	expected := 2
	if d.Err() != nil {
		t.Errorf("slice.Dynamic[int].Err() is not nil: %s", d.err.Error())
		return 
	}

	if got != expected {
		t.Errorf("slice.Dynamic[int].Front() = %d, want %d", got, expected)
	}
}

func TestBack(t *testing.T) {
	got := d.Back()
	expected := 4
	if d.Err() != nil {
		t.Errorf("slice.Dynamic[int].Err() is not nil: %s", d.err.Error())
		return 
	}

	if got != expected {
		t.Errorf("slice.Dynamic[int].Back() = %d, want %d", got, expected)
	}
}

func TestAt(t *testing.T) {
	got := d.At(1)
	expected := 3
	if d.Err() != nil {
		t.Errorf("slice.Dynamic[int].Err() is not nil: %s", d.err.Error())
		return 
	}

	if got != expected {
		t.Errorf("slice.Dynamic[int].At(1) = %d, want %d", got, expected)
	}
}
