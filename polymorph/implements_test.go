package polymorph

import "testing"

type someInterfaceForV interface {
	A() string
}

type v struct{} // implements someInterfaceForV
type g struct{} // doesn't implement someInterfaceForV

func (v) A() string { return "A" }

func TestImplements(t *testing.T) {
	tests := []struct {
		body      string
		got, want bool
	}{
		{"Implements[someInterfaceForV, v]()", Implements[someInterfaceForV, v](), true},
		{"Implements[someInterfaceForV, g]()", Implements[someInterfaceForV, g](), false},
		// Implements expects I to be an interface.
		// If it is not, Implements returns false.
		{"Implements[g, g]()", Implements[g, g](), false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func BenchmarkImplements(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_ = Implements[someInterfaceForV, v]()
	}
}
