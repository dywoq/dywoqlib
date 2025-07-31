package polymorph

import "testing"

type numStruct1 struct{}

func (numStruct1) A() {}
func (numStruct1) B() {}

type numStruct2 struct {
	A string
	B int
}

func TestNumMethods(t *testing.T) {
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

func TestNumFields(t *testing.T) {
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

func BenchmarkNumMethods(b *testing.B) {
	for b.Loop() {
		_ = NumMethods[numStruct1]()
	}
}

func BenchmarkNumFields(b *testing.B) {
	for b.Loop() {
		_ = NumFields[numStruct1]()
	}
}
