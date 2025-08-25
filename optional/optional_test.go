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

package optional

import (
	"errors"
	go_testing "testing"

	"github.com/dywoq/dywoqlib/err"
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
)

func TestPresent(t *go_testing.T) {
	cases := []struct {
		name string
		opt  Maybe[int]
		want bool
	}{
		{"with value", New(10), true},
		{"without value", None[int](), false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *go_testing.T) {
			got := tc.opt.Present()
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGet(t *go_testing.T) {
	cases := []struct {
		name        string
		opt         Maybe[int]
		wantVal     int
		wantPresent bool
	}{
		{"with value", New(10), 10, true},
		{"without value", None[int](), 0, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *go_testing.T) {
			gotVal, gotPresent := tc.opt.Get()
			if gotVal != tc.wantVal {
				t.Errorf("got value %v, want %v", gotVal, tc.wantVal)
			}
			if gotPresent != tc.wantPresent {
				t.Errorf("got present %v, want %v", gotPresent, tc.wantPresent)
			}
		})
	}
}

func TestElse(t *go_testing.T) {
	cases := []struct {
		name   string
		opt    Maybe[int]
		orElse int
		want   int
	}{
		{"with value", New(10), 20, 10},
		{"without value", None[int](), 40, 40},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *go_testing.T) {
			got := tc.opt.Else(tc.orElse)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestString(t *go_testing.T) {
	cases := []struct {
		name string
		opt  Maybe[int]
		want string
	}{
		{"with value", New(10), "10"},
		{"without value", Int(), "0"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *go_testing.T) {
			got := tc.opt.String()
			if got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestFilter(t *go_testing.T) {
	type user struct {
		age int
	}

	cases := []struct {
		name string
		opt  Maybe[user]
		fn   func(u user) bool
		want bool
	}{
		{"age is 19", New(user{19}), func(u user) bool { return u.age >= 18 }, true},
		{"age is 14", New(user{14}), func(u user) bool { return u.age >= 18 }, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *go_testing.T) {
			got := tc.opt.Filter(tc.fn).Present()
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestUnwrap(t *go_testing.T) {
	t.Run("no panic", func(t *go_testing.T) {
		opt := Int(10)
		defer func() {
			if r := recover(); r != nil {
				t.Error("recovered panic despite the presence of the Maybe value")
			}
		}()
		_ = opt.Unwrap()
	})

	t.Run("panic", func(t *go_testing.T) {
		opt := Int()
		defer func() {
			if r := recover(); r == nil {
				t.Error("no recovered panic despite no presence of the Maybe value")
			}
		}()
		_ = opt.Unwrap()
	})
}

func TestOr(t *go_testing.T) {
	f := func() uint {
		return 2 * 2
	}

	cases := []struct {
		name string
		opt  Maybe[uint]
		want uint
	}{
		{"return Maybe value", UInt(20), 20},
		{"return the result of the function", UInt(), f()},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *go_testing.T) {
			got := tc.opt.Or(f)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMap(t *go_testing.T) {
	f := func(s uint8) uint8 {
		return s * 2
	}

	t.Run("if Maybe value is present", func(t *go_testing.T) {
		opt := UInt8(20)
		got := Map(opt, f)
		gotVal, _ := got.Get()
		wantVal, _ := opt.Get()
		if gotVal != wantVal*2 {
			t.Errorf("got %v, want %v", gotVal, wantVal*2)
		}
	})

	t.Run("if Maybe value is not present", func(t *go_testing.T) {
		opt := UInt8()
		got := Map(opt, f)
		present := got.Present()
		wantPresent := false
		if present != wantPresent {
			t.Errorf("got %v, want %v", present, wantPresent)
		}
	})
}

func TestError(t *go_testing.T) {
	t.Run("error context is present", func(tt *go_testing.T) {
		someError := errors.New("some error")
		context := err.NewContext(someError, "you will give me some oreo!!")
		opt := NoneContext[int](context)
		got := opt.Error().Nil()
		want := false
		if got != want {
			tt.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("error context is not present", func(tt *go_testing.T) {
		opt := Int()
		got := opt.Error().Nil()
		want := true
		if got != want {
			tt.Errorf("got %v, want %v", got, want)
		}
	})
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

func BenchmarkFilter(b *go_testing.B) {
	type user struct {
		age int
	}
	internal_testing.SetBase().Benchmark(b)
	opt := New(user{19})
	for b.Loop() {
		_ = opt.Filter(func(u user) bool {
			return u.age >= 18
		})
	}
}

func BenchmarkUnwrap(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	opt := New(20)
	for b.Loop() {
		_ = opt.Unwrap()
	}
}

func BenchmarkOr(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	f := func() uint {
		return 2 * 2
	}
	opt := UInt(20)
	for b.Loop() {
		_ = opt.Or(f)
	}
}

func BenchmarkMap(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	f := func(s uint8) uint8 {
		return s * 2
	}
	opt := UInt8(20)
	for b.Loop() {
		_ = Map(opt, f)
	}
}

func BenchmarkError(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	someError := errors.New("some error")
	context := err.NewContext(someError, "you will give me some oreo!!")
	opt := NoneContext[int](context)
	for b.Loop() {
		_ = opt.Error()
	}
}
