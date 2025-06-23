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

package attribute

import (
	"fmt"
	"strings"
)

// Deprecated is an attribute-function that generates warning about deprecated function.
// DOES NOT automatically returns zero values (e.g., "", 0, nil etc.).
//
// If custom event is not nil (it's set by SetEvent),
// then it uses the custom event instead of outputting the warning.
func Deprecated() {
	m := management{}
	targetSkip, sourceSkip := m.skipNums()
	target := m.funcInfo(targetSkip)
	source := m.funcInfo(sourceSkip)
	res := deprecatedFormat(target, source)
	if event != nil {
		event()
		return
	}
	fmt.Println(res)
}

func deprecatedFormat(target, source string) string {
	strs := make([]string, 5)
	strs = append(strs, "attribute.Deprecated: ", target, " is deprecated; ")
	strs = append(strs, "source: ", source)
	res := strings.Join(strs, "")
	return res
}
