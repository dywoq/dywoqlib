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

package stringn

import (
	go_testing "testing"

	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
)

func BenchmarkAppend(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	for b.Loop() {
		s := New("hello")
		s.Append(" world")
	}
}

func BenchmarkAt(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello world")
	for b.Loop() {
		s.At(6)
	}
}

func BenchmarkInsert(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	for b.Loop() {
		s := New("hello world")
		s.Insert(6, '!')
	}
}

func BenchmarkRemoveRange(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	for b.Loop() {
		s := New("hello world")
		s.RemoveRange(5, 11)
	}
}

func BenchmarkReplace(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	for b.Loop() {
		s := New("hello world")
		s.Replace("world", "go")
	}
}

func BenchmarkReverse(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	for b.Loop() {
		s := New("hello world")
		s.Reverse()
	}
}

func BenchmarkToLower(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("HELLO WORLD")
	for b.Loop() {
		s.ToLower()
	}
}

func BenchmarkToUpper(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello world")
	for b.Loop() {
		s.ToUpper()
	}
}

func BenchmarkCompare(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello world")
	for b.Loop() {
		s.Compare("hello world")
	}
}

func BenchmarkEquals(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello world")
	for b.Loop() {
		s.Equals("hello world")
	}
}

func BenchmarkSplit(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello world, how are you?")
	for b.Loop() {
		s.Split(" ")
	}
}

func BenchmarkSubstring(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello world")
	for b.Loop() {
		s.Substring(6, 11)
	}
}

func BenchmarkPrepend(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	for b.Loop() {
		s := New("world")
		s.Prepend("hello ")
	}
}

func BenchmarkContainsRune(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello world")
	for b.Loop() {
		s.ContainsRune('w')
	}
}

func BenchmarkContainsString(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello world")
	for b.Loop() {
		s.ContainsString("world")
	}
}

func BenchmarkFront(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello world")
	for b.Loop() {
		s.Front()
	}
}

func BenchmarkBack(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello world")
	for b.Loop() {
		s.Back()
	}
}

func BenchmarkHasRunePrefix(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello world")
	for b.Loop() {
		s.HasRunePrefix('h')
	}
}

func BenchmarkHasStringPrefix(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello world")
	for b.Loop() {
		s.HasStringPrefix("hello")
	}
}

func BenchmarkHasRuneSuffix(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello world")
	for b.Loop() {
		s.HasRuneSuffix('d')
	}
}

func BenchmarkHasStringSuffix(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello world")
	for b.Loop() {
		s.HasStringSuffix("world")
	}
}

func BenchmarkSet(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	for b.Loop() {
		s := New("hello world")
		s.Set('W', 6)
	}
}

func BenchmarkNative(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello")

	for b.Loop() {
		_ = s.Native()
	}
}

func BenchmarkString(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello")

	for b.Loop() {
		_ = s.String()
	}
}

func BenchmarkLength(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello")

	for b.Loop() {
		_ = s.Length()
	}
}

func BenchmarkError(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello")

	for b.Loop() {
		_ = s.Error()
	}
}

func BenchmarkClear(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello")

	for b.Loop() {
		s.Clear()
	}
}

func BenchmarkEmpty(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello")

	for b.Loop() {
		_ = s.Empty()
	}
}

func BenchmarkGrow(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello")

	for b.Loop() {
		s.Grow(10)
	}
}

func BenchmarkWrite(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	data := []byte(" world")
	for b.Loop() {
		s := New("hello")
		_, _ = s.Write(data)
	}
}

func BenchmarkRead(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello world")
	buf := make([]byte, s.Length())

	for b.Loop() {
		b.StopTimer()
		s = New("hello world")
		b.StartTimer()
		_, _ = s.Read(buf)
	}
}

func BenchmarkIterating(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	s := New("hello")
	for b.Loop() {
		it := s.Iterating().Forward()
		for it.Next() {
			_ = it.Value()
		}
	}
}

func BenchmarkNew(b *go_testing.B) {
	base := internal_testing.SetBase()
	base.Benchmark(b)
	for b.Loop() {
		_ = New("hello")
	}
}
