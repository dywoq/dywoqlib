package asm

import (
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	go_testing "testing"
	"unsafe"
)

func TestString(t *go_testing.T) {
	data := []byte("hello world")

	tests := []struct {
		title string
		ptr   unsafe.Pointer
		len   int
		want  string
	}{
		{"full string", unsafe.Pointer(&data[0]), len(data), "hello world"},
		{"partial string", unsafe.Pointer(&data[6]), 5, "world"},
		{"empty string", unsafe.Pointer(&data[0]), 0, ""},
	}

	for _, tc := range tests {
		t.Run(tc.title, func(t *go_testing.T) {
			got := String(tc.ptr, tc.len)
			if got != tc.want {
				t.Errorf("String() = %q, want %q", got, tc.want)
			}
		})
	}
}

func BenchmarkString(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	data := []byte("hello world")
	for b.Loop() {
		_ = String(unsafe.Pointer(&data[0]), len(data))
	}
}
