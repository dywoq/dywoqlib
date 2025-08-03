package optional

import (
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	go_testing "testing"
)

func TestPresent(t *go_testing.T) {
	// first test
	opt := New(10)
	got := opt.Present()
	want := true
	if got != want {
		t.Errorf("opt.Present() = %v, want %v", got, want)
	}

	// second test
	opt = None[int]()
	got = opt.Present()
	want = false
	if got != want {
		t.Errorf("opt.Present() = %v, want %v", got, want)
	}
}

func TestGet(t *go_testing.T) {
	opt := New(10)

	// first test
	_, got := opt.Get()
	want := true
	if got != want {
		t.Errorf("opt.Get() = %v, want %v", got, want)
	}

	// second test
	gotval, _ := opt.Get()
	wantval := 10
	if gotval != wantval {
		t.Errorf("opt.Get() = %v, want %v", gotval, wantval)
	}
}

func TestElse(t *go_testing.T) {
	// first test
	opt := New(10)
	got := opt.Else(20)
	want := 10
	if got != want {
		t.Errorf("opt.Else(20) = %d, want %d", got, want)
	}

	// second test
	opt = None[int]()
	got = opt.Else(40)
	want = 40
	if got != want {
		t.Errorf("opt.Else(40) = %d, want %d", got, want)
	}
}

func TestString(t *go_testing.T) {
	// first test
	opt := New(10)
	got := opt.String()
	want := "10"
	if got != want {
		t.Errorf("opt.String() = %s, want %s", got, want)
	}

	// second test
	opt = Int()
	got = opt.String()
	want = "0"
	if got != want {
		t.Errorf("opt.String() = %s, want %s", got, want)
	}
}

func BenchmarkPresent(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	opt := New(10)
	for b.Loop() {
		_ = opt.Present()
	}
}

func BenchmarkGet(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	opt := New(10)
	for b.Loop() {
		_, _ = opt.Get()
	}
}

func BenchmarkElse(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	opt := New(10)
	for b.Loop() {
		_ = opt.Else(20)
	}
}
