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

package slice

import (
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	"slices"
	go_testing "testing"
)

func TestMergeDynamic(t *go_testing.T) {
	var (
		first  = NewDynamic(1, 2)
		second = NewDynamic(3, 4)
	)
	got, err := MergeDynamic(first, second)
	if err != nil {
		t.Fatal(err)
	}
	expected := []int{1, 2, 3, 4}
	if !slices.Equal(got.Native(), expected) {
		t.Errorf("MergeDynamic(first, second) = %v, want %v", got.Native(), expected)
	}
}

func TestMergeFixed(t *go_testing.T) {
	var (
		first  = NewFixed(10, 1, 2)
		second = NewFixed(10, 3, 4)
	)
	if first.Error() != nil {
		t.Error(first.Error())
	}
	if second.Error() != nil {
		t.Error(second.Error())
	}

	got, err := MergeFixed(first, second)
	if got.Error() != nil {
		t.Fatal(got.Error())
	}
	if err != nil {
		t.Fatal(err)
	}

	expected := []int{1, 2, 3, 4}
	if !slices.Equal(got.Native(), expected) {
		t.Errorf("MergeFixed(first, second) = %v, want %v", got.Native(), expected)
	}
}

func TestMerge(t *go_testing.T) {
	var (
		first  = []int{10, 54, 3}
		second = []int{89, 23, 10}
		got    = Merge(first, second)
		want   = []int{10, 54, 3, 89, 23, 10}
	)
	if !slices.Equal(got, want) {
		t.Errorf("Merge(first, second) = %v, want %v", got, want)
	}
}

func BenchmarkMergeDynamic(b *go_testing.B) {
	var (
		first  = NewDynamic(10, 1, 2)
		second = NewDynamic(10, 3, 4)
	)
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		_, _ = MergeDynamic(first, second)
	}
}

func BenchmarkMergeFixed(b *go_testing.B) {
	var (
		first  = NewFixed(10, 1, 2)
		second = NewFixed(10, 3, 4)
	)
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		_, _ = MergeFixed(first, second)
	}
}

func BenchmarkMerge(b *go_testing.B) {
	var (
		first  = []int{10, 54, 3}
		second = []int{89, 23, 10}
	)
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		_ = Merge(first, second)
	}
}
