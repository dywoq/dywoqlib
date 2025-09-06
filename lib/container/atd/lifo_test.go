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

package atd

import (
	"slices"
	"testing"
)

func TestLifoLength(t *testing.T) {
	l := NewLifo[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)

	got := l.Length()
	want := 4

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLifoEmpty(t *testing.T) {
	l := NewLifo[int]()

	got := l.Empty()
	want := true

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLifoAppend(t *testing.T) {
	l := NewLifo[int]()
	l.Append(2)
	l.Append(4)

	got := l.Native()
	want := []int{2, 4}

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLifoPop(t *testing.T) {
	l := NewLifo[int]()
	l.Append(2)
	l.Append(4)
	l.Pop()

	got := l.Native()
	want := []int{2}

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLifoTop(t *testing.T) {
	l := NewLifo[int]()
	l.Append(2)
	l.Append(4)

	got := l.Top()
	want := 4

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLifoString(t *testing.T) {
	l := NewLifo[int]()
	l.Append(2)
	l.Append(4)

	got := l.String()
	want := "[2, 4]"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
