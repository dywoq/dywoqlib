package polymorph

import "testing"

func TestComparable(t *testing.T) {
	tests := []struct {
		got, want bool
	}{
		{Comparable[int](), true},
		{Comparable[map[string]int](), false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("got %v, want %v", test.got, test.want)
		}
	}
}
