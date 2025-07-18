package polymorph

import "testing"

func TestComparable(t *testing.T) {
	if !Comparable[int]() {
		t.Errorf("!Comparable[int]() = false, want true")
	}
	if Comparable[map[string]int]() {
		t.Errorf("!Comparable[int]() = false, want true")
	}
}
