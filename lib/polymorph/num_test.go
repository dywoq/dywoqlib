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

type numStruct1 struct{}

func (numStruct1) A() {}
func (numStruct1) B() {}

type numStruct2 struct {
	A string
	B int
}

func TestNumMethods(t *go_testing.T) {
	tests := []struct {
		body      string
		got, want int
	}{
		{"NumMethods[numStruct1]()", NumMethods[numStruct1](), 2},
		{"NumMethods[numStruct2]()", NumMethods[numStruct2](), 0},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func TestNumFields(t *go_testing.T) {
	tests := []struct {
		body      string
		got, want int
	}{
		{"NumFields[numStruct1]()", NumFields[numStruct1](), 0},
		{"NumFields[numStruct2]()", NumFields[numStruct2](), 2},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func BenchmarkNumMethods(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		_ = NumMethods[numStruct1]()
	}
}

func BenchmarkNumFields(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	b.ReportAllocs()
	for b.Loop() {
		_ = NumFields[numStruct1]()
	}
}
