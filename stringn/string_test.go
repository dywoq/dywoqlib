package stringn

import (
	"fmt"
	"io"
	"slices"
	"testing"
)

func TestAppend(t *testing.T) {
	str := New("HI!")
	tests := []struct {
		body      string
		got, want []string
	}{
		{"str.Append(\"bye\", \"hi1\")", str.Append("bye", "hi1"), []string{"bye", "hi1"}},
		{"str.Append(\"sd\", \"sd\", \"sd\")", str.Append("sd", "sd", "sd"), []string{"sd", "sd", "sd"}},
	}

	for _, test := range tests {
		if !slices.Equal(test.got, test.want) {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func TestFront(t *testing.T) {
	str1 := New("HI!")
	str2 := New("I")
	tests := []struct {
		body      string
		got, want rune
	}{
		{"str1.Front()", str1.Front(), 'H'},
		{"str2.Front()", str2.Front(), 'I'},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func TestBack(t *testing.T) {
	str1 := New("HI!")
	str2 := New("I")
	tests := []struct {
		body      string
		got, want rune
	}{
		{"str1.Back()", str1.Back(), '!'},
		{"str2.Back()", str2.Back(), 'I'},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func TestHasRunePrefix(t *testing.T) {
	str1 := New("HI!")
	tests := []struct {
		body      string
		got, want bool
	}{
		{"str1.HasRunePrefix('H')", str1.HasRunePrefix('H'), true},
		{"str1.HasRunePrefix('H')", str1.HasRunePrefix('I'), false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func TestHasStringPrefix(t *testing.T) {
	str1 := New("HI!")
	tests := []struct {
		body      string
		got, want bool
	}{
		{"str1.HasStringPrefix(\"HI\")", str1.HasStringPrefix("HI"), true},
		{"str1.HasStringPrefix(\"I!\")", str1.HasStringPrefix("I!"), false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func TestHasRuneSuffix(t *testing.T) {
	str1 := New("HI!")
	tests := []struct {
		body      string
		got, want bool
	}{
		{"str1.HasRuneSuffix('!')", str1.HasRuneSuffix('!'), true},
		{"str1.HasRuneSuffix('H')", str1.HasRuneSuffix('H'), false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func TestHasStringSuffix(t *testing.T) {
	str1 := New("bye")
	tests := []struct {
		body      string
		got, want bool
	}{
		{"str1.HasStringSuffix(\"e\")", str1.HasStringSuffix("e"), true},
		{"str1.HasStringSuffix(\"b\")", str1.HasStringSuffix("b"), false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func TestInsert(t *testing.T) {
	str1 := New("bye")
	str1.Insert(0, 'H')
	if str1.Error() != nil {
		t.Fatal(str1.Error())
	}
	want := "Hbye"
	if str1.Native() != want {
		t.Errorf("got %s, want %s", str1.Native(), want)
	}
}

func TestСontainsRune(t *testing.T) {
	str1 := New("bye")
	tests := []struct {
		body      string
		got, want bool
	}{
		{"str1.ContainsRune('b')", str1.ContainsRune('b'), true},
		{"str1.ContainsRune('a')", str1.ContainsRune('a'), false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func TestСontainsString(t *testing.T) {
	str1 := New("hello")
	tests := []struct {
		body      string
		got, want bool
	}{
		{"str1.ContainsString(\"hel\")", str1.ContainsString("hel"), true},
		{"str1.ContainsString(\"llo\")", str1.ContainsString("llo"), true},
		{"str1.ContainsString(\"j\")", str1.ContainsString("j"), false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func TestWrite(t *testing.T) {
	str1 := New("hello")
	fmt.Fprintf(str1, ". bye")

	got := str1.Native()
	want := "hello. bye"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRead(t *testing.T) {
	str1 := New("hello")
	want := str1.Native()

	got, _ := io.ReadAll(str1)

	if str1.Error() != nil {
		t.Fatal(str1.Error())
	}

	if string(got) != want {
		t.Errorf("got %v, want %v", string(got), want)
	}
}
func TestEmpty(t *testing.T) {
	str1 := New("hello")
	str2 := New("")
	tests := []struct {
		body      string
		got, want bool
	}{
		{"str1.Empty()", str1.Empty(), false},
		{"str2.Empty()", str2.Empty(), true},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, %v", test.body, test.got, test.want)
		}
	}
}

func TestNative(t *testing.T) {
	str1 := New("hello")
	got := str1.Native()
	want := "hello"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestSet(t *testing.T) {
	str1 := New("hello")
	str1.Set('H', 0)
	if str1.Error() != nil {
		t.Fatal(str1.Error())
	}
	want := "Hello"
	if str1.Native() != want {
		t.Errorf("got %s, want %s", str1.Native(), want)
	}
}
