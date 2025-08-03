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
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	go_testing "testing"
)

type someInterfaceForV interface {
	A() string
}

type v struct{} // implements someInterfaceForV
type g struct{} // doesn't implement someInterfaceForV

func (v) A() string { return "A" }

func TestImplements(t *go_testing.T) {
	tests := []struct {
		body      string
		got, want bool
	}{
		{"Implements[someInterfaceForV, v]()", Implements[someInterfaceForV, v](), true},
		{"Implements[someInterfaceForV, g]()", Implements[someInterfaceForV, g](), false},
		// Implements expects I to be an interface.
		// If it is not, Implements returns false.
		{"Implements[g, g]()", Implements[g, g](), false},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %v, want %v", test.body, test.got, test.want)
		}
	}
}

func BenchmarkImplements(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		_ = Implements[someInterfaceForV, v]()
	}
}
