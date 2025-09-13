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
	"reflect"
	"slices"
	go_testing "testing"

	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
)

func TestFifoNative(t *go_testing.T) {
	tests := []struct {
		name     string
		fifo     *Fifo[int]
		appended []int
		want     []int
	}{
		{"empty slice", NewFifo[int](), []int{}, []int{}},
		{"not empty slice", NewFifo[int](), []int{2, 3, 4}, []int{2, 3, 4}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			for _, appending := range test.appended {
				test.fifo.Append(appending)
			}
			got := test.fifo.Native()
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFifoEmpty(t *go_testing.T) {
	tests := []struct {
		name     string
		fifo     *Fifo[int]
		appended []int
		want     bool
	}{
		{"not empty", NewFifo[int](), []int{2, 3}, false},
		{"empty", NewFifo[int](), []int{}, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			for _, appending := range test.appended {
				test.fifo.Append(appending)
			}
			got := test.fifo.Empty()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFifoAppend(t *go_testing.T) {
	tests := []struct {
		name     string
		fifo     *Fifo[int]
		appended []int
		want     []int
	}{
		{"appending", NewFifo[int](), []int{2, 3, 4}, []int{2, 3, 4}},
		{"not appending", NewFifo[int](), []int{}, []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			for _, appending := range test.appended {
				test.fifo.Append(appending)
			}
			got := test.fifo.Native()
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFifoFront(t *go_testing.T) {
	tests := []struct {
		name     string
		fifo     *Fifo[int]
		appended []int
		want     int
	}{
		{"front", NewFifo[int](), []int{2, 3}, 2},
		{"empty front", NewFifo[int](), []int{}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			for _, appending := range test.appended {
				test.fifo.Append(appending)
			}
			got := test.fifo.Front()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFifoBack(t *go_testing.T) {
	tests := []struct {
		name     string
		fifo     *Fifo[int]
		appended []int
		want     int
	}{
		{"back", NewFifo[int](), []int{2, 3}, 3},
		{"empty back", NewFifo[int](), []int{}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			for _, appending := range test.appended {
				test.fifo.Append(appending)
			}
			got := test.fifo.Back()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFifoPop(t *go_testing.T) {
	tests := []struct {
		name     string
		fifo     *Fifo[int]
		appended []int
		want     []int
	}{
		{"popping", NewFifo[int](), []int{2, 3, 4}, []int{2, 3}},
		{"popping on empty slice", NewFifo[int](), []int{}, []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			for _, appending := range test.appended {
				test.fifo.Append(appending)
			}
			test.fifo.Pop()
			got := test.fifo.Native()
			if !slices.Equal(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFifoString(t *go_testing.T) {
	tests := []struct {
		name     string
		fifo     *Fifo[int]
		appended []int
		want     string
	}{
		{"not empty slice", NewFifo[int](), []int{10, 12}, "[10, 12]"},
		{"empty slice", NewFifo[int](), []int{}, "[]"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			for _, appending := range test.appended {
				test.fifo.Append(appending)
			}
			got := test.fifo.String()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func BenchmarkFifoNative(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	fifo := NewFifo[int]()
	fifo.Append(20)
	fifo.Append(10)
	for b.Loop() {
		_ = fifo.Native()
	}
}

func BenchmarkFifoEmpty(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	fifo := NewFifo[int]()
	fifo.Append(20)
	fifo.Append(10)
	for b.Loop() {
		_ = fifo.Empty()
	}
}

func BenchmarkFifoLength(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	fifo := NewFifo[int]()
	fifo.Append(20)
	fifo.Append(10)
	for b.Loop() {
		_ = fifo.Length()
	}
}

func BenchmarkFifoFront(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	fifo := NewFifo[int]()
	fifo.Append(20)
	fifo.Append(10)
	for b.Loop() {
		_ = fifo.Front()
	}
}

func BenchmarkFifoBack(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	fifo := NewFifo[int]()
	fifo.Append(20)
	fifo.Append(10)
	for b.Loop() {
		_ = fifo.Back()
	}
}

func BenchmarkFifoAppend(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	fifo := NewFifo[int]()
	fifo.Append(20)
	fifo.Append(10)
	for b.Loop() {
		_ = fifo.Append(3)
	}
}

func BenchmarkFifoPop(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	fifo := NewFifo[int]()
	fifo.Append(20)
	fifo.Append(10)
	for b.Loop() {
		_ = fifo.Pop()
	}
}

func BenchmarkFifoString(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	fifo := NewFifo[int]()
	fifo.Append(20)
	fifo.Append(10)
	for b.Loop() {
		_ = fifo.String()
	}
}
