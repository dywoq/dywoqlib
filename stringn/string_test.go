// Copyright 2025 dywoq
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func TestGrow(t *testing.T) {
	str := New("hello")
	str.Grow(10)
	if str.b.Cap() < 15 {
		t.Errorf("expected capacity >= 15, got %d", str.b.Cap())
	}
}

func TestToLower(t *testing.T) {
	str := New("HeLLo")
	str.ToLower()
	want := "hello"
	if str.Native() != want {
		t.Errorf("got %s, want %s", str.Native(), want)
	}
}

func TestToUpper(t *testing.T) {
	str := New("HeLLo")
	str.ToUpper()
	want := "HELLO"
	if str.Native() != want {
		t.Errorf("got %s, want %s", str.Native(), want)
	}
}

func TestInsert_OutOfBounds(t *testing.T) {
	str := New("hello")
	str.Insert(10, '!')
	if str.Error() == nil {
		t.Errorf("expected an error for out-of-bounds insert")
	}
	want := "hello"
	if str.Native() != want {
		t.Errorf("string should not be modified on error. got %s, want %s", str.Native(), want)
	}
}

func TestRemoveRange(t *testing.T) {
	str := New("hello")
	str.RemoveRange(1, 4)
	want := "ho"
	if str.Native() != want {
		t.Errorf("got %s, want %s", str.Native(), want)
	}
}

func TestReverse(t *testing.T) {
	str := New("hello")
	str.Reverse()
	want := "olleh"
	if str.Native() != want {
		t.Errorf("got %s, want %s", str.Native(), want)
	}
}

func TestSubstring(t *testing.T) {
	str := New("hello")
	substr := str.Substring(1, 4)
	want := "ell"
	if substr != want {
		t.Errorf("got %s, want %s", substr, want)
	}
}

func TestSplit(t *testing.T) {
	str := New("a,b,c")
	parts := str.Split(",")
	want := []string{"a", "b", "c"}
	if !slices.Equal(parts, want) {
		t.Errorf("got %v, want %v", parts, want)
	}
}

func TestCompare(t *testing.T) {
	str := New("a")
	if str.Compare("a") != 0 {
		t.Errorf("expected 0")
	}
	if str.Compare("b") != -1 {
		t.Errorf("expected -1")
	}
	if str.Compare("A") != 1 {
		t.Errorf("expected 1")
	}
}

func TestEquals(t *testing.T) {
	str := New("hello")
	if !str.Equals("hello") {
		t.Errorf("expected true")
	}
	if str.Equals("world") {
		t.Errorf("expected false")
	}
}

func TestAt(t *testing.T) {
	str := New("hello")
	r := str.At(1)
	if r != 'e' {
		t.Errorf("got %c, want %c", r, 'e')
	}
	str.At(10)
	if str.Error() == nil {
		t.Errorf("expected error for out-of-bounds At")
	}
}

func TestPrepend(t *testing.T) {
	str := New("world")
	str.Prepend("hello ", "beautiful ")
	want := "hello beautiful world"
	if str.Native() != want {
		t.Errorf("got %s, want %s", str.Native(), want)
	}
}

func TestContainsRune(t *testing.T) {
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

func TestContainsString(t *testing.T) {
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
