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

package unique

import (
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	"reflect"
	"slices"
	go_testing "testing"
)

func TestLifoLength(t *go_testing.T) {
	tests := []struct {
		name     string
		lifo     *Lifo[int]
		appended []int
		want     int
	}{
		{"not zero length", NewLifo[int](), []int{2, 3}, 2},
		{"zero length", NewLifo[int](), []int{}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			for _, appending := range test.appended {
				test.lifo.Append(appending)
			}
			got := test.lifo.Length()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestLifoEmpty(t *go_testing.T) {
	tests := []struct {
		name     string
		lifo     *Lifo[int]
		appended []int
		want     bool
	}{
		{"not empty", NewLifo[int](), []int{2, 3}, false},
		{"empty", NewLifo[int](), []int{}, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			for _, appending := range test.appended {
				test.lifo.Append(appending)
			}
			got := test.lifo.Empty()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestLifoAppend(t *go_testing.T) {
	tests := []struct {
		name     string
		lifo     *Lifo[int]
		appended []int
		want     []int
	}{
		{"appending", NewLifo[int](), []int{2, 3, 4}, []int{2, 3, 4}},
		{"not appending", NewLifo[int](), []int{}, []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			for _, appending := range test.appended {
				test.lifo.Append(appending)
			}
			got := test.lifo.Native()
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestLifoPop(t *go_testing.T) {
	tests := []struct {
		name     string
		lifo     *Lifo[int]
		appended []int
		want     []int
	}{
		{"popping", NewLifo[int](), []int{2, 3, 4}, []int{2, 3}},
		{"popping on empty slice", NewLifo[int](), []int{}, []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			for _, appending := range test.appended {
				test.lifo.Append(appending)
			}
			test.lifo.Pop()
			got := test.lifo.Native()
			if !slices.Equal(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestLifoTop(t *go_testing.T) {
	tests := []struct {
		name     string
		lifo     *Lifo[int]
		appended []int
		want     int
	}{
		{"successful top", NewLifo[int](), []int{2, 3}, 2},
		{"getting top from empty slice", NewLifo[int](), []int{}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			for _, appending := range test.appended {
				test.lifo.Append(appending)
			}
			test.lifo.Pop()
			got := test.lifo.Top()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestLifoString(t *go_testing.T) {
	tests := []struct {
		name     string
		lifo     *Lifo[int]
		appended []int
		want     string
	}{
		{"not empty slice", NewLifo[int](), []int{10, 12}, "[10, 12]"},
		{"empty slice", NewLifo[int](), []int{}, "[]"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			for _, appending := range test.appended {
				test.lifo.Append(appending)
			}
			got := test.lifo.String()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func BenchmarkLength(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	lifo := NewLifo[int]()
	lifo.Append(2)
	lifo.Append(3)
	for b.Loop() {
		_ = lifo.Length()
	}
}

func BenchmarkEmpty(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	lifo := NewLifo[int]()
	lifo.Append(2)
	lifo.Append(3)
	for b.Loop() {
		_ = lifo.Empty()
	}
}

func BenchmarkAppend(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	lifo := NewLifo[int]()
	for b.Loop() {
		_ = lifo.Append(2)
	}
}

func BenchmarkPop(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	lifo := NewLifo[int]()
	lifo.Append(2)
	lifo.Append(3)
	for b.Loop() {
		_ = lifo.Pop()
	}
}

func BenchmarkString(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	lifo := NewLifo[int]()
	lifo.Append(2)
	lifo.Append(3)
	for b.Loop() {
		_ = lifo.String()
	}
}
