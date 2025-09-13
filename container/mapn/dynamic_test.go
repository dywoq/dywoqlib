package mapn

import (
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	"maps"
	go_testing "testing"
)

func TestDynamicLength(t *go_testing.T) {
	tests := []struct {
		name string
		m    *Dynamic[int, int]
		want int
	}{
		{"not empty map", NewDynamic(map[int]int{2: 2, 3: 3}), 2},
		{"empty map", NewDynamic(map[int]int{}), 0},
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

func TestDynamicExists(t *go_testing.T) {
	tests := []struct {
		name string
		m    *Dynamic[int, int]
		key  int
		want bool
	}{
		{"does exist", NewDynamic(map[int]int{2: 2}), 2, true},
		{"does not exist", NewDynamic(map[int]int{2: 2}), 3, false},
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

func TestDynamicAdd(t *go_testing.T) {
	tests := []struct {
		name string
		m    *Dynamic[int, int]
		add  struct{ key, value int }
		want map[int]int
	}{
		{"key does exist", NewDynamic(map[int]int{2: 2}), struct {
			key   int
			value int
		}{2, 2}, map[int]int{2: 2}},
		{"key does not exist", NewDynamic(map[int]int{2: 2}), struct {
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

func TestDynamicSet(t *go_testing.T) {
	tests := []struct {
		name string
		m    *Dynamic[int, int]
		set  struct{ key, value int }
		want map[int]int
	}{
		{"key does exist", NewDynamic(map[int]int{2: 2}), struct {
			key   int
			value int
		}{2, 3}, map[int]int{2: 2}},
		{"key does not exist", NewDynamic(map[int]int{2: 2}), struct {
			key   int
			value int
		}{3, 3}, map[int]int{2: 2, 3: 3}},
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

func TestDynamicDelete(t *go_testing.T) {
	tests := []struct {
		name    string
		m       *Dynamic[int, int]
		deleted int
		want    map[int]int
	}{
		{"key does exist", NewDynamic(map[int]int{2: 2}), 2, map[int]int{}},
		{"key does not exist", NewDynamic(map[int]int{2: 2}), 3, map[int]int{2: 2}},
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

func TestDynamicGet(t *go_testing.T) {
	tests := []struct {
		name string
		m    *Dynamic[int, int]
		key  int
		want struct {
			key, value int
		}
	}{
		{"key does exists", NewDynamic(map[int]int{2: 2}), 2, struct {
			key   int
			value int
		}{2, 2}},
		{"key does not exist", NewDynamic(map[int]int{2: 2}), 3, struct {
			key   int
			value int
		}{0, 0}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			gotkey, gotvalue := test.m.Get(test.key)
			want := test.want
			if gotkey != want.key {
				t.Errorf("got key %v, want key %v", gotkey, want.key)
			}

			if gotvalue != want.value {
				t.Errorf("got value %v, want value %v", gotvalue, want.value)
			}
		})
	}
}

func BenchmarkDynamicLength(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewDynamic(map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_ = m.Length()
	}
}

func BenchmarkDynamicExists(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewDynamic(map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_ = m.Exists(2)
	}
}

func BenchmarkDynamicAdd(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewDynamic(map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_, _ = m.Add(2, 2)
	}
}

func BenchmarkDynamicSet(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewDynamic(map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_, _ = m.Set(2, 2)
	}
}

func BenchmarkDynamicDelete(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewDynamic(map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_ = m.Delete(2)
	}
}

func BenchmarkDynamicGet(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	m := NewDynamic(map[int]int{2: 2, 3: 3})
	for b.Loop() {
		_, _ = m.Get(2)
	}
}
