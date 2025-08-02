// Copyright 2025 dywoq
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package polymorph

import (
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	"reflect"
	go_testing "testing"
)

func TestTypeOfGeneric(t *go_testing.T) {
	tests := []struct {
		name string
		want reflect.Type
	}{
		{"int", reflect.TypeOf(0)},
		{"string", reflect.TypeOf("")},
		{"bool", reflect.TypeOf(true)},
		{"float64", reflect.TypeOf(0.0)},
		{"*int", reflect.TypeOf((*int)(nil))},
		{"[]string", reflect.TypeOf([]string{})},
		{"map[string]int", reflect.TypeOf(map[string]int{})},
		{"struct{}", reflect.TypeOf(struct{}{})},
		{"struct{Name string; Age int}", reflect.TypeOf(struct {
			Name string
			Age  int
		}{})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *go_testing.T) {
			var got reflect.Type
			switch tt.name {
			case "int":
				got = TypeOfGeneric[int]()
			case "string":
				got = TypeOfGeneric[string]()
			case "bool":
				got = TypeOfGeneric[bool]()
			case "float64":
				got = TypeOfGeneric[float64]()
			case "*int":
				got = TypeOfGeneric[*int]()
			case "[]string":
				got = TypeOfGeneric[[]string]()
			case "map[string]int":
				got = TypeOfGeneric[map[string]int]()
			case "struct{}":
				got = TypeOfGeneric[struct{}]()
			case "struct{Name string; Age int}":
				got = TypeOfGeneric[struct {
					Name string
					Age  int
				}]()
			default:
				t.Fatalf("unhandled test case: %s", tt.name)
			}

			if got != tt.want {
				t.Errorf("TypeOfGeneric[%s]() got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func BenchmarkTypeOfGeneric(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		_ = TypeOfGeneric[int]()
	}
}

func BenchmarkTypeOfGenericPointer(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	for b.Loop() {
		_ = TypeOfGeneric[*int]()
	}
}
