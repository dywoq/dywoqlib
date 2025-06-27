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

// Package attribute provides attribute-like functions such as Deprecated and Todo
// for marking functions as deprecated or unimplemented, with optional custom event handling.
package attribute

import (
	"fmt"
	"strings"
)

// Todo is an attribute function that generates warning about unimplemented function.
// DOES NOT automatically returns zero values (e.g., "", 0, nil etc.).
func Todo() {
	m := management{}
	targetSkip, sourceSkip := m.skipNums()
	target := m.funcInfo(targetSkip)
	source := m.funcInfo(sourceSkip)
	msg := todoFormat(target, source)
	if event != nil {
		event()
		return
	}
	fmt.Println(msg)
}

func todoFormat(target, source string) string {
	strs := make([]string, 5)
	strs = append(strs, "attribute.Todo: todo in ", target, "; ")
	strs = append(strs, "source: ", source)
	res := strings.Join(strs, "")
	return res
}
