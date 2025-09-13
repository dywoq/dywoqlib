package unique

import (
	"reflect"
	"slices"
	go_testing "testing"

	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
)

func TestSliceNative(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Slice[int]
		want  []int
	}{
		{"not empty slice", NewSlice(3, 4, 5), []int{3, 4, 5}},
		{"empty slice", NewSlice[int](), []int{}},
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

func TestSliceLength(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Slice[int]
		want  int
	}{
		{"not zero length", NewSlice(2, 3, 4), 3},
		{"zero length", NewSlice[int](), 0},
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

func TestSliceAppend(t *go_testing.T) {
	tests := []struct {
		name     string
		slice    *Slice[int]
		appended []int
		want     []int
	}{
		{"appending elements", NewSlice(2, 3, 4), []int{5, 6}, []int{2, 3, 4, 5, 6}},
		{"not-appending elements", NewSlice(2, 3, 4), []int{}, []int{2, 3, 4}},
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

func TestSliceAt(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Slice[int]
		i     int
		want  int
	}{
		{"correct index", NewSlice(2, 3, 4), 0, 2},
		{"wrong index", NewSlice(2, 3, 4), 10, 0},
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

func TestSliceFind(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Slice[int]
		req   int
		want  int
	}{
		{"found", NewSlice(2, 3, 4), 2, 2},
		{"didn't find", NewSlice(2, 3, 4), 10, 0},
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

func TestSliceString(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Slice[int]
		want  string
	}{
		{"empty slice", NewSlice[int](), "[]"},
		{"not empty slice", NewSlice(2, 3, 4), "[2, 3, 4]"},
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

func TestSliceSet(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Slice[int]
		i     int
		elem  int
		want  []int
	}{
		{"successful setting", NewSlice(2, 3, 4), 2, 5, []int{2, 3, 5}},
		{"not successful setting", NewSlice(2, 3, 4), 10, 5, []int{2, 3, 4}},
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

func TestSliceDelete(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Slice[int]
		i     int
		want  []int
	}{
		{"successful deleting", NewSlice(2, 3, 4), 2, []int{2, 3, 0}},
		{"not successful deleting", NewSlice(2, 3, 4), 10, []int{2, 3, 4}},
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

func TestSliceInsert(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Slice[int]
		i     int
		elem  int
		want  []int
	}{
		{"successful inserting", NewSlice(2, 3, 4), 0, 4, []int{4, 2, 3, 4}},
		{"not successful inserting", NewSlice(2, 3, 4), 10, 4, []int{2, 3, 4}},
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

func TestSliceFront(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Slice[int]
		want  int
	}{
		{"successful front", NewSlice(2, 3, 4), 2},
		{"not successful front", NewSlice[int](), 0},
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

func TestSliceBack(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Slice[int]
		want  int
	}{
		{"successful back", NewSlice(2, 3, 4), 4},
		{"not successful back", NewSlice[int](), 0},
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

func TestSlicePop(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Slice[int]
		want  []int
	}{
		{"successful popping", NewSlice(2, 3, 4), []int{2, 3}},
		{"pop on empty slice", NewSlice[int](), []int{}},
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

func BenchmarkSliceNative(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = slice.Native()
	}
}

func BenchmarkSliceLength(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = slice.Length()
	}
}

func BenchmarkSliceAppend(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = slice.Append(10, 19)
	}
}

func BenchmarkSliceAt(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = slice.At(2)
	}
}

func BenchmarkSliceFind(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = slice.Find(2)
	}
}

func BenchmarkSliceString(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = slice.String()
	}
}

func BenchmarkSliceSet(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = slice.Set(2, 2)
	}
}

func BenchmarkSliceDelete(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = slice.Delete(2)
	}
}

func BenchmarkSliceInsert(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = slice.Insert(2, 2)
	}
}

func BenchmarkSliceFront(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = slice.Front()
	}
}

func BenchmarkSliceBack(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = slice.Back()
	}
}

func BenchmarkSlicePop(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewSlice(2, 3, 4)
	for b.Loop() {
		_ = slice.Pop()
	}
}
