package polymorph

import (
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	go_testing "testing"
)

type f string

type a struct {
	b string
}

func (a) A() {}

func TestHasMethod(t *go_testing.T) {
	tests := []struct {
		body      string
		got, want bool
	}{
		{"HasMethod[a](\"A\")", HasMethod[a]("A"), true},
		{"HasMethod[a](\"B\")", HasMethod[a]("B"), false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func TestHasField(t *go_testing.T) {
	tests := []struct {
		body      string
		got, want bool
	}{
		{"HasField[a](\"b\")", HasField[a]("b"), true},
		{"HasField[a](\"B\")", HasField[a]("B"), false},
		// HasField expects S be struct, not type
		{"HasField[f](\"b\")", HasField[f]("b"), false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func BenchmarkHasMethod(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		_ = HasMethod[a]("A")
	}
}

func BenchmarkHasField(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		_ = HasField[a]("b")
	}
}
