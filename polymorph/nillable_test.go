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

package polymorph

import (
	go_testing "testing"

	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
)

func TestNillable(t *go_testing.T) {
	tests := []struct {
		body      string
		got, want bool
	}{
		{"Nillable[chan int]()", Nillable[chan int](), true},
		{"Nillable[func(string)]()", Nillable[func(string)](), true},
		{"Nillable[interface { A() string }]()", Nillable[interface{ A() string }](), true},
		{"Nillable[map[int]int]()", Nillable[map[int]int](), true},
		{"Nillable[*int]()", Nillable[*int](), true},
		{"Nillable[[]int]()", Nillable[[]int](), true},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, %v", test.body, test.got, test.want)
		}
	}
}

func BenchmarkNillable(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		_ = Nillable[chan int]()
	}
}
