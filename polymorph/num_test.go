package polymorph

import (
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	go_testing "testing"
)

type numStruct1 struct{}

func (numStruct1) A() {}
func (numStruct1) B() {}

type numStruct2 struct {
	A string
	B int
}

func TestNumMethods(t *go_testing.T) {
	tests := []struct {
		body      string
		got, want int
	}{
		{"NumMethods[numStruct1]()", NumMethods[numStruct1](), 2},
		{"NumMethods[numStruct2]()", NumMethods[numStruct2](), 0},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func TestNumFields(t *go_testing.T) {
	tests := []struct {
		body      string
		got, want int
	}{
		{"NumFields[numStruct1]()", NumFields[numStruct1](), 0},
		{"NumFields[numStruct2]()", NumFields[numStruct2](), 2},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func BenchmarkNumMethods(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		_ = NumMethods[numStruct1]()
	}
}

func BenchmarkNumFields(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	b.ReportAllocs()
	for b.Loop() {
		_ = NumFields[numStruct1]()
	}
}
