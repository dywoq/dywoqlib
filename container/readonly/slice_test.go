package readonly

import (
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	go_testing "testing"
)

func TestSliceLength(t *go_testing.T) {
	tests := []struct {
		name string
		s    *Slice[int]
		want int
	}{
		{"non-zero length", NewSlice(2, 3, 4), 3},
		{"zero length", NewSlice[int](), 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.s.Length()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestSliceAt(t *go_testing.T) {
	tests := []struct {
		name string
		s    *Slice[int]
		i    int
		want int
	}{
		{"successful at", NewSlice(2, 3, 4), 0, 2},
		{"not successful at", NewSlice(2, 3, 4), 5, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.s.Length()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestSliceFind(t *go_testing.T) {
	tests := []struct {
		name      string
		s         *Slice[int]
		req, want int
	}{
		{"successful finding", NewSlice(2, 3, 4), 2, 2},
		{"not successful finding", NewSlice(2, 3, 4), 10, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.s.Find(test.req)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestSliceString(t *go_testing.T) {
	tests := []struct {
		name string
		s    *Slice[int]
		want string
	}{
		{"successful formatting", NewSlice(2, 3, 4), "[2, 3, 4]"},
		{"not successful formatting", NewSlice[int](), "[]"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.s.String()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestSliceFront(t *go_testing.T) {
	tests := []struct {
		name string
		s    *Slice[int]
		want int
	}{
		{"successful front", NewSlice(2, 3, 4), 2},
		{"not successful front", NewSlice[int](), 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.s.Front()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestSliceBack(t *go_testing.T) {
	tests := []struct {
		name string
		s    *Slice[int]
		want int
	}{
		{"successful back", NewSlice(2, 3, 4), 4},
		{"not successful back", NewSlice[int](), 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.s.Back()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func BenchmarkTestLength(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	s := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = s.Length()
	}
}

func BenchmarkTestAt(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	s := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = s.At(2)
	}
}

func BenchmarkTestFind(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	s := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = s.Find(2)
	}
}

func BenchmarkTestString(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	s := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = s.String()
	}
}

func BenchmarkTestFront(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	s := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = s.Front()
	}
}

func BenchmarkTestBack(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	s := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = s.Back()
	}
}
