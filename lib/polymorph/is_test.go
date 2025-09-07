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
	internal_testing "github.com/dywoq/dywoqlib/lib/internal/testing"
	go_testing "testing"
)

func TestIs(t *go_testing.T) {
	tests := []struct {
		name string
		got  bool
		want bool
	}{
		{"int satisfies 100", Is[int](100), true},
		{"int satisfies uint(100)", Is[int](uint(100)), false},
		{"string satisfies \"Hi!\"", Is[string]("Hi!"), true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			if test.got != test.want {
				t.Errorf("%v != %v", test.got, test.want)
			}
		})
	}
}

func BenchmarkIs(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		_ = Is[int](100)
	}
}
