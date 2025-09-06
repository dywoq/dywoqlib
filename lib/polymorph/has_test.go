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

	internal_testing "github.com/dywoq/dywoqlib/lib/internal/testing"
)

type f string

type a struct {
	b string
}

func (a) A() {}

func TestHasMethod(t *go_testing.T) {
	tests := []struct {
		body      string
		got, want bool
	}{
		{"HasMethod[a](\"A\")", HasMethod[a]("A"), true},
		{"HasMethod[a](\"B\")", HasMethod[a]("B"), false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func TestHasField(t *go_testing.T) {
	tests := []struct {
		body      string
		got, want bool
	}{
		{"HasField[a](\"b\")", HasField[a]("b"), true},
		{"HasField[a](\"B\")", HasField[a]("B"), false},
		// HasField expects S be struct, not type
		{"HasField[f](\"b\")", HasField[f]("b"), false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func BenchmarkHasMethod(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		_ = HasMethod[a]("A")
	}
}

func BenchmarkHasField(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		_ = HasField[a]("b")
	}
}
