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

package slice

import (
	"reflect"
	"slices"
	go_testing "testing"

	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
)

func TestFixedNative(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Fixed[int]
		want  []int
	}{
		{"not empty slice", NewFixed(10, 2, 3, 4), []int{2, 3, 4}},
		{"empty slice", NewFixed[int](10), []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.slice.Native()
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFixedLength(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Fixed[int]
		want  int
	}{
		{"non-zero length", NewFixed(10, 2, 3, 4), 3},
		{"zero length", NewFixed[int](10), 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.slice.Length()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFixedAppend(t *go_testing.T) {
	tests := []struct {
		name     string
		slice    *Fixed[int]
		appended []int
		want     []int
	}{
		{"appending elements", NewFixed(10, 2, 3, 4), []int{5, 6}, []int{2, 3, 4, 5, 6}},
		{"not-appending elements", NewFixed(10, 2, 3, 4), []int{}, []int{2, 3, 4}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			test.slice.Append(test.appended...)
			got := test.slice.Native()
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFixedAt(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Fixed[int]
		i     int
		want  int
	}{
		{"correct index", NewFixed(10, 2, 3, 4), 0, 2},
		{"wrong index", NewFixed(10, 2, 3, 4), 10, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.slice.At(test.i)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFixedFind(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Fixed[int]
		req   int
		want  int
	}{
		{"found", NewFixed(10, 2, 3, 4), 2, 2},
		{"didn't find", NewFixed(10, 2, 3, 4), 10, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.slice.Find(test.req)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFixedString(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Fixed[int]
		want  string
	}{
		{"empty slice", NewFixed[int](10), "[]"},
		{"not empty slice", NewFixed(10, 2, 3, 4), "[2, 3, 4]"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.slice.String()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFixedSet(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Fixed[int]
		i     int
		elem  int
		want  []int
	}{
		{"successful setting", NewFixed(10, 2, 3, 4), 2, 5, []int{2, 3, 5}},
		{"not successful setting", NewFixed(10, 2, 3, 4), 10, 5, []int{2, 3, 4}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			test.slice.Set(test.elem, test.i)
			got := test.slice.Native()
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFixedDelete(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Fixed[int]
		i     int
		want  []int
	}{
		{"successful deleting", NewFixed(10, 2, 3, 4), 2, []int{2, 3, 0}},
		{"not successful deleting", NewFixed(10, 2, 3, 4), 10, []int{2, 3, 4}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			test.slice.Delete(test.i)
			got := test.slice.Native()
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFixedInsert(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Fixed[int]
		i     int
		elem  int
		want  []int
	}{
		{"successful inserting", NewFixed(10, 2, 3, 4), 0, 4, []int{4, 2, 3, 4}},
		{"not successful inserting", NewFixed(10, 2, 3, 4), 10, 4, []int{2, 3, 4}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			test.slice.Insert(test.i, test.elem)
			got := test.slice.Native()
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFixedFront(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Fixed[int]
		want  int
	}{
		{"successful front", NewFixed(10, 2, 3, 4), 2},
		{"not successful front", NewFixed[int](10), 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.slice.Front()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFixedBack(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Fixed[int]
		want  int
	}{
		{"successful back", NewFixed(10, 2, 3, 4), 4},
		{"not successful back", NewFixed[int](10), 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.slice.Back()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFixedPop(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Fixed[int]
		want  []int
	}{
		{"successful popping", NewFixed(10, 2, 3, 4), []int{2, 3}},
		{"pop on empty slice", NewFixed[int](10), []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			test.slice.Pop()
			got := test.slice.Native()
			if !slices.Equal(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func BenchmarkFixedNative(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewFixed(10, 2, 3, 4)
	for b.Loop() {
		_ = slice.Native()
	}
}

func BenchmarkFixedLength(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewFixed(10, 2, 3, 4)
	for b.Loop() {
		_ = slice.Length()
	}
}

func BenchmarkFixedAppend(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewFixed(10, 2, 3, 4)
	for b.Loop() {
		_ = slice.Append(10, 19)
	}
}

func BenchmarkFixedAt(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewFixed(10, 2, 3, 4)
	for b.Loop() {
		_ = slice.At(2)
	}
}

func BenchmarkFixedFind(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewFixed(10, 2, 3, 4)
	for b.Loop() {
		_ = slice.Find(2)
	}
}

func BenchmarkFixedString(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewFixed(10, 2, 3, 4)
	for b.Loop() {
		_ = slice.String()
	}
}

func BenchmarkFixedSet(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewFixed(10, 2, 3, 4)
	for b.Loop() {
		_ = slice.Set(2, 2)
	}
}

func BenchmarkFixedDelete(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewFixed(10, 2, 3, 4)
	for b.Loop() {
		_ = slice.Delete(2)
	}
}

func BenchmarkFixedInsert(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewFixed(10, 2, 3, 4)
	for b.Loop() {
		_ = slice.Insert(2, 2)
	}
}

func BenchmarkFixedFront(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewFixed(10, 2, 3, 4)
	for b.Loop() {
		_ = slice.Front()
	}
}

func BenchmarkFixedBack(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewFixed(10, 2, 3, 4)
	for b.Loop() {
		_ = slice.Back()
	}
}

func BenchmarkFixedPop(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewFixed(10, 2, 3, 4)
	for b.Loop() {
		_ = slice.Pop()
	}
}
