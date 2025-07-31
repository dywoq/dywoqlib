package polymorph

import "testing"

func TestComparable(t *testing.T) {
	tests := []struct {
		body string
		got, want bool
	}{
		{"Comparable[int]()", Comparable[int](), true},
		{"Comparable[map[string]int]()", Comparable[map[string]int](), false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func BenchmarkComparable(b *testing.B) {
	for b.Loop() {
		_ = Comparable[int]()
	}
}
