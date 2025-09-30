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

// filter provides a functionality to filter slices and maps.
//
// # Purpose
//
// The purpose is to reduce rewriting the same logic code.
//
// The developers have to rewrite the same logic of filtering slices and maps every time,
// which is redundant and time-wasting. filter package solves this problem.
//
// # Example of using
//
//  package main
//
//  import (
//    "fmt"
//    "github.com/dywoq/dywoqlib/filter"
//  )
//
//  func main() {
//     s := filter.Slice([]int{1, 2, 3, 4, 5, 6}, func(elem int) bool { return elem%2 == 0 })
//     fmt.Println(s) // Output: [2 4 6]
//  }
//
package filter
