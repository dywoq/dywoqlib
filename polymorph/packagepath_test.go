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
	"io"
	"net/http"
	go_testing "testing"

	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
)

func TestPackagePath(t *go_testing.T) {
	tests := []struct {
		body      string
		got, want string
	}{
		{"PackagePath[http.Server]()", PackagePath[http.Server](), "net/http"},
		{"PackagePath[io.Writer]()", PackagePath[io.Writer](), "io"},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("%s = %s, want %s", test.body, test.got, test.want)
		}
	}
}

func BenchmarkPackagePath(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	b.ReportAllocs()
	for b.Loop() {
		_ = PackagePath[http.Server]()
	}
}
