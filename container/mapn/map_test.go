package mapn

import "testing"

var m = New(map[string]int{"one": 1, "two": 5})

func TestMapAt(t *testing.T) {
	var tests = []struct {
		key  string
		got  int
		want int
	}{
		{"one", m.At("one"), 1},
		{"two", m.At("two"), 5},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("m.At(\"%s\") = %d, want %d", test.key, test.got, test.want)
		}
	}
}
