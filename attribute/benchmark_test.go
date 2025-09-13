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

package attribute_test

import (
	go_testing "testing"

	"github.com/dywoq/dywoqlib/attribute"
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
)

func BenchmarkDeprecated(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		attribute.Deprecated(func() {})
	}
}

func BenchmarkExperimental(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		attribute.Experimental(func() {})
	}
}

func BenchmarkRemoved(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		attribute.Removed(func() {})
	}
}

func BenchmarkTodo(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		attribute.Todo(func() {})
	}
}

func BenchmarkUnsafe(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		attribute.Unsafe(func() {})
	}
}
