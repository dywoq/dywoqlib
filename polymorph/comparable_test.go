package polymorph

import (
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	go_testing "testing"
)

func TestComparable(t *go_testing.T) {
	tests := []struct {
		body      string
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

func BenchmarkComparable(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		_ = Comparable[int]()
	}
}
