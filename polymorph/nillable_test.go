package polymorph

import (
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	go_testing "testing"
)

func TestNillable(t *go_testing.T) {
	tests := []struct {
		body      string
		got, want bool
	}{
		{"Nillable[chan int]()", Nillable[chan int](), true},
		{"Nillable[func(string)]()", Nillable[func(string)](), true},
		{"Nillable[interface { A() string }]()", Nillable[interface{ A() string }](), true},
		{"Nillable[map[int]int]()", Nillable[map[int]int](), true},
		{"Nillable[*int]()", Nillable[*int](), true},
		{"Nillable[[]int]()", Nillable[[]int](), true},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, %v", test.body, test.got, test.want)
		}
	}
}

func BenchmarkNillable(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		_ = Nillable[chan int]()
	}
}
