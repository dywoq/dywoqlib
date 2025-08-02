package polymorph

import (
	"reflect"
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	go_testing "testing"
)

func TestKindOf(t *go_testing.T) {
	tests := []struct {
		body      string
		got, want reflect.Kind
	}{
		{"KindOf[bool]()", KindOf[bool](), reflect.Bool},

		{"KindOf[int]()", KindOf[int](), reflect.Int},
		{"KindOf[int8]()", KindOf[int8](), reflect.Int8},
		{"KindOf[int16]()", KindOf[int16](), reflect.Int16},
		{"KindOf[int32]()", KindOf[int32](), reflect.Int32},
		{"KindOf[int64]()", KindOf[int64](), reflect.Int64},

		{"KindOf[uint]()", KindOf[uint](), reflect.Uint},
		{"KindOf[uint8]()", KindOf[uint8](), reflect.Uint8},
		{"KindOf[uint16]()", KindOf[uint16](), reflect.Uint16},
		{"KindOf[uint32]()", KindOf[uint32](), reflect.Uint32},
		{"KindOf[uint64](),", KindOf[uint64](), reflect.Uint64},

		{"KindOf[string]()", KindOf[string](), reflect.String},

		{"KindOf[float32]()", KindOf[float32](), reflect.Float32},
		{"KindOf[float64]()", KindOf[float64](), reflect.Float64},

		{"KindOf[complex64]()", KindOf[complex64](), reflect.Complex64},
		{"KindOf[complex128]()", KindOf[complex128](), reflect.Complex128},

		{"KindOf[struct{ name string }]()", KindOf[struct{ name string }](), reflect.Struct},
		{"KindOf[*int]()", KindOf[*int](), reflect.Pointer},
		{"KindOf[[]int]()", KindOf[[]int](), reflect.Slice},
		{"KindOf[[10]int]()", KindOf[[10]int](), reflect.Array},
		{"KindOf[interface{ name() string }]()", KindOf[interface{ name() string }](), reflect.Interface},
		{"KindOf[chan int]()", KindOf[chan int](), reflect.Chan},
		{"KindOf[uintptr]()", KindOf[uintptr](), reflect.Uintptr},
		{"KindOf[func(string)]()", KindOf[func(string)](), reflect.Func},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %s, want %s", test.body, test.got.String(), test.want.String())
		}
	}
}

func BenchmarkKindOf(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		_ = KindOf[int64]()
	}
}
