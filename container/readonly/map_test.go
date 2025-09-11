package readonly

import (
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	"slices"
	go_testing "testing"
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

func TestMapKeys(t *go_testing.T) {
	m := NewMap(map[int]int{2: 2, 3: 3})
	got := m.Keys()
	want := []int{2, 3}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMapValues(t *go_testing.T) {
	m := NewMap(map[int]int{2: 2, 3: 3})
	got := m.Values()
	want := []int{2, 3}
	if !slices.Equal(got, want) {
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
