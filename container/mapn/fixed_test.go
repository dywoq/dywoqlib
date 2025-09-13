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

package mapn

import (
	"maps"
	go_testing "testing"

	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
)

func TestFixedLength(t *go_testing.T) {
	tests := []struct {
		name string
		m    *Fixed[int, int]
		want int
	}{
		{"not empty map", NewFixed(10, map[int]int{2: 2, 3: 3}), 2},
		{"empty map", NewFixed(10, map[int]int{}), 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.m.Length()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFixedExists(t *go_testing.T) {
	tests := []struct {
		name string
		m    *Fixed[int, int]
		key  int
		want bool
	}{
		{"does exist", NewFixed(10, map[int]int{2: 2}), 2, true},
		{"does not exist", NewFixed(10, map[int]int{2: 2}), 3, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.m.Exists(test.key)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFixedAdd(t *go_testing.T) {
	tests := []struct {
		name string
		m    *Fixed[int, int]
		add  struct{ key, value int }
		want map[int]int
	}{
		{"key does exist", NewFixed(10, map[int]int{2: 2}), struct {
			key   int
			value int
		}{2, 2}, map[int]int{2: 2}},
		{"key does not exist", NewFixed(10, map[int]int{2: 2}), struct {
			key   int
			value int
		}{3, 3}, map[int]int{2: 2, 3: 3}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			test.m.Add(test.add.key, test.add.value)

			got := test.m.Native()
			if !maps.Equal(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFixedSet(t *go_testing.T) {
	tests := []struct {
		name string
		m    *Fixed[int, int]
		set  struct{ key, value int }
		want map[int]int
	}{
		{"key does exist", NewFixed(10, map[int]int{2: 2}), struct {
			key   int
			value int
		}{2, 3}, map[int]int{2: 3}},
		{"key does not exist", NewFixed(10, map[int]int{2: 2}), struct {
			key   int
			value int
		}{3, 3}, map[int]int{2: 2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			test.m.Set(test.set.key, test.set.value)

			got := test.m.Native()
			if !maps.Equal(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFixedDelete(t *go_testing.T) {
	tests := []struct {
		name    string
		m       *Fixed[int, int]
		deleted int
		want    map[int]int
	}{
		{"key does exist", NewFixed(10, map[int]int{2: 2}), 2, map[int]int{}},
		{"key does not exist", NewFixed(10, map[int]int{2: 2}), 3, map[int]int{2: 2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			test.m.Delete(test.deleted)

			got := test.m.Native()
			if !maps.Equal(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestFixedGet(t *go_testing.T) {
	tests := []struct {
		name string
		m    *Fixed[int, int]
		key  int
		want struct {
			key, value int
		}
	}{
		{"key does exists", NewFixed(10, map[int]int{2: 2}), 2, struct {
			key   int
			value int
		}{2, 2}},
		{"key does not exist", NewFixed(10, map[int]int{2: 2}), 3, struct {
			key   int
			value int
		}{0, 0}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			gotkey, gotvalue := test.m.Get(test.key)
			want := test.want
			if gotkey != want.key {
				t.Errorf("got key %v, want key %v", gotkey, want.value)
			}

			if gotvalue != want.value {
				t.Errorf("got value %v, want value %v", gotvalue, want.value)
			}
		})
	}
}

func BenchmarkFixedLength(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewFixed(10, map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_ = m.Length()
	}
}

func BenchmarkFixedExists(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewFixed(10, map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_ = m.Exists(2)
	}
}

func BenchmarkFixedAdd(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewFixed(10, map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_, _ = m.Add(2, 2)
	}
}

func BenchmarkFixedSet(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewFixed(10, map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_, _ = m.Set(2, 2)
	}
}

func BenchmarkFixedDelete(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewFixed(10, map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_ = m.Delete(2)
	}
}

func BenchmarkFixedGet(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewFixed(10, map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_, _ = m.Get(2)
	}
}
