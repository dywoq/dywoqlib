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
	"testing"
)

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := New("hello")
		s.Append(" world")
	}
}

func BenchmarkAt(b *testing.B) {
	s := New("hello world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.At(6)
	}
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := New("hello world")
		s.Insert(6, '!')
	}
}

func BenchmarkRemoveRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := New("hello world")
		s.RemoveRange(5, 11)
	}
}

func BenchmarkReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := New("hello world")
		s.Replace("world", "go")
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := New("hello world")
		s.Reverse()
	}
}

func BenchmarkToLower(b *testing.B) {
	s := New("HELLO WORLD")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.ToLower()
	}
}

func BenchmarkToUpper(b *testing.B) {
	s := New("hello world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.ToUpper()
	}
}

func BenchmarkCompare(b *testing.B) {
	s := New("hello world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Compare("hello world")
	}
}

func BenchmarkEquals(b *testing.B) {
	s := New("hello world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Equals("hello world")
	}
}

func BenchmarkSplit(b *testing.B) {
	s := New("hello world, how are you?")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Split(" ")
	}
}

func BenchmarkSubstring(b *testing.B) {
	s := New("hello world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Substring(6, 11)
	}
}

func BenchmarkPrepend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := New("world")
		s.Prepend("hello ")
	}
}

func BenchmarkContainsRune(b *testing.B) {
	s := New("hello world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.ContainsRune('w')
	}
}

func BenchmarkContainsString(b *testing.B) {
	s := New("hello world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.ContainsString("world")
	}
}

func BenchmarkFront(b *testing.B) {
	s := New("hello world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Front()
	}
}

func BenchmarkBack(b *testing.B) {
	s := New("hello world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Back()
	}
}

func BenchmarkHasRunePrefix(b *testing.B) {
	s := New("hello world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.HasRunePrefix('h')
	}
}

func BenchmarkHasStringPrefix(b *testing.B) {
	s := New("hello world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.HasStringPrefix("hello")
	}
}

func BenchmarkHasRuneSuffix(b *testing.B) {
	s := New("hello world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.HasRuneSuffix('d')
	}
}

func BenchmarkHasStringSuffix(b *testing.B) {
	s := New("hello world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.HasStringSuffix("world")
	}
}

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := New("hello world")
		s.Set('W', 6)
	}
}

func BenchmarkNative(b *testing.B) {
	s := New("hello")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = s.Native()
	}
}

func BenchmarkString(b *testing.B) {
	s := New("hello")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = s.String()
	}
}

func BenchmarkLength(b *testing.B) {
	s := New("hello")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = s.Length()
	}
}

func BenchmarkError(b *testing.B) {
	s := New("hello")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = s.Error()
	}
}

func BenchmarkClear(b *testing.B) {
	s := New("hello")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Clear()
	}
}

func BenchmarkEmpty(b *testing.B) {
	s := New("hello")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = s.Empty()
	}
}

func BenchmarkGrow(b *testing.B) {
	s := New("hello")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Grow(10)
	}
}

func BenchmarkWrite(b *testing.B) {
	data := []byte(" world")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := New("hello")
		b.StartTimer()
		_, _ = s.Write(data)
	}
}

func BenchmarkRead(b *testing.B) {
	s := New("hello world")
	buf := make([]byte, s.Length())
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s = New("hello world")
		b.StartTimer()
		_, _ = s.Read(buf)
	}
}

func BenchmarkIterating(b *testing.B) {
	s := New("hello")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it := s.Iterating().Forward()
		for it.Next() {
			_ = it.Value()
		}
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New("hello")
	}
}
