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

package readonly

import (
	"reflect"
	"sort"
	go_testing "testing"

	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
)

func TestMapLength(t *go_testing.T) {
	m := NewMap(map[int]int{2: 2, 3: 3})
	got := m.Length()
	want := 2
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMapExists(t *go_testing.T) {
	m := NewMap(map[int]int{2: 2, 3: 3})
	got := m.Exists(2)
	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMapGet(t *go_testing.T) {
	m := NewMap(map[int]int{2: 2, 3: 3})
	got, _ := m.Get(2)
	want := 2
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMapKeys(t *go_testing.T) {
	m := NewMap(map[int]int{2: 2, 3: 3})

	got := m.Keys()
	sort.Ints(got)
	want := []int{2, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMapValues(t *go_testing.T) {
	m := NewMap(map[int]int{2: 2, 3: 3})

	got := m.Values()
	sort.Ints(got)
	want := []int{2, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func BenchmarkLength(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewMap(map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_ = m.Length()
	}
}

func BenchmarkExists(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewMap(map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_ = m.Exists(2)
	}
}

func BenchmarkKeys(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewMap(map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_ = m.Keys()
	}
}

func BenchmarkValues(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewMap(map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_ = m.Values()
	}
}

func BenchmarkGet(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewMap(map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_, _ = m.Get(2)
	}
}

func BenchmarkString(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewMap(map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_ = m.String()
	}
}
