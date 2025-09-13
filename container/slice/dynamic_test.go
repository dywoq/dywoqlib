package slice

import (
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	"reflect"
	"slices"
	go_testing "testing"
)

func TestDynamicNative(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Dynamic[int]
		want  []int
	}{
		{"not empty slice", NewDynamic(2, 3, 4), []int{2, 3, 4}},
		{"empty slice", NewDynamic[int](), []int{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.slice.Native()
			if !slices.Equal(got, test.want) {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestDynamicLength(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Dynamic[int]
		want  int
	}{
		{"non-zero length", NewDynamic(2, 3, 4), 3},
		{"zero length", NewDynamic[int](), 0},
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

func TestDynamicAppend(t *go_testing.T) {
	tests := []struct {
		name     string
		slice    *Dynamic[int]
		appended []int
		want     []int
	}{
		{"appending elements", NewDynamic(2, 3, 4), []int{5, 6}, []int{2, 3, 4, 5, 6}},
		{"not-appending elements", NewDynamic(2, 3, 4), []int{}, []int{2, 3, 4}},
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

func TestDynamicAt(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Dynamic[int]
		i     int
		want  int
	}{
		{"correct index", NewDynamic(2, 3, 4), 0, 2},
		{"wrong index", NewDynamic(2, 3, 4), 10, 0},
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

func TestDynamicFind(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Dynamic[int]
		req   int
		want  int
	}{
		{"found", NewDynamic(2, 3, 4), 2, 2},
		{"didn't find", NewDynamic(2, 3, 4), 10, 0},
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

func TestDynamicString(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Dynamic[int]
		want  string
	}{
		{"empty slice", NewDynamic[int](), "[]"},
		{"not empty slice", NewDynamic(2, 3, 4), "[2, 3, 4]"},
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

func TestDynamicSet(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Dynamic[int]
		i     int
		elem  int
		want  []int
	}{
		{"successful setting", NewDynamic(2, 3, 4), 2, 5, []int{2, 3, 5}},
		{"not successful setting", NewDynamic(2, 3, 4), 10, 5, []int{2, 3, 4}},
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

func TestDynamicDelete(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Dynamic[int]
		i     int
		want  []int
	}{
		{"successful deleting", NewDynamic(2, 3, 4), 2, []int{2, 3, 0}},
		{"not successful deleting", NewDynamic(2, 3, 4), 10, []int{2, 3, 4}},
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

func TestDynamicInsert(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Dynamic[int]
		i     int
		elem  int
		want  []int
	}{
		{"successful inserting", NewDynamic(2, 3, 4), 0, 4, []int{4, 2, 3, 4}},
		{"not successful inserting", NewDynamic(2, 3, 4), 10, 4, []int{2, 3, 4}},
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

func TestDynamicFront(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Dynamic[int]
		want  int
	}{
		{"successful front", NewDynamic(2, 3, 4), 2},
		{"not successful front", NewDynamic[int](), 0},
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

func TestDynamicBack(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Dynamic[int]
		want  int
	}{
		{"successful back", NewDynamic(2, 3, 4), 4},
		{"not successful back", NewDynamic[int](), 0},
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

func TestDynamicPop(t *go_testing.T) {
	tests := []struct {
		name  string
		slice *Dynamic[int]
		want  []int
	}{
		{"successful popping", NewDynamic(2, 3, 4), []int{2, 3}},
		{"pop on empty slice", NewDynamic[int](), []int{}},
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

func BenchmarkDynamicNative(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewDynamic(2, 3, 4)
	for b.Loop() {
		_ = slice.Native()
	}
}

func BenchmarkDynamicLength(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewDynamic(2, 3, 4)
	for b.Loop() {
		_ = slice.Length()
	}
}

func BenchmarkDynamicAppend(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewDynamic(2, 3, 4)
	for b.Loop() {
		_ = slice.Append(10, 19)
	}
}

func BenchmarkDynamicAt(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewDynamic(2, 3, 4)
	for b.Loop() {
		_ = slice.At(2)
	}
}

func BenchmarkDynamicFind(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewDynamic(2, 3, 4)
	for b.Loop() {
		_ = slice.Find(2)
	}
}

func BenchmarkDynamicString(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewDynamic(2, 3, 4)
	for b.Loop() {
		_ = slice.String()
	}
}

func BenchmarkDynamicSet(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewDynamic(2, 3, 4)
	for b.Loop() {
		_ = slice.Set(2, 2)
	}
}

func BenchmarkDynamicDelete(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewDynamic(2, 3, 4)
	for b.Loop() {
		_ = slice.Delete(2)
	}
}

func BenchmarkDynamicInsert(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewDynamic(2, 3, 4)
	for b.Loop() {
		_ = slice.Insert(2, 2)
	}
}

func BenchmarkDynamicFront(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewDynamic(2, 3, 4)
	for b.Loop() {
		_ = slice.Front()
	}
}

func BenchmarkDynamicBack(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewDynamic(2, 3, 4)
	for b.Loop() {
		_ = slice.Back()
	}
}

func BenchmarkDynamicPop(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	slice := NewDynamic(2, 3, 4)
	for b.Loop() {
		_ = slice.Pop()
	}
}
