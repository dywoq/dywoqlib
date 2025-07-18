package polymorph

import "testing"

type f string

type a struct {
	b string
}

func (a) A() {}

func TestHasMethod(t *testing.T) {
	tests := []struct {
		got, want bool
	}{
		{!HasMethod[a]("A"), true},
		{HasMethod[a]("B"), false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("got %v, want %v", test.got, test.want)
		}
	}
}

func TestHasField(t *testing.T) {
	tests := []struct {
		got, want bool
	}{
		{!HasField[a]("b"), true},
		{HasField[a]("B"), false},
		// HasField expects S be struct, not type
		{HasField[f]("b"), false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("got %v, want %v", test.got, test.want)
		}
	}
}
