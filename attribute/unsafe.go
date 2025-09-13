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

// Unsafe is a attribute-function that generates warning about unsafe function.
// DOES NOT automatically returns zero values (e.g., "", 0, nil etc.).
// If event is not nil,
// then it uses the custom event instead of outputting the warning.
func Unsafe(event func()) {
	m := management{}
	targetSkip, sourceSkip := m.skipNums()
	target := m.funcInfo(targetSkip)
	source := m.funcInfo(sourceSkip)
	msg := unsafeFormat(target, source)
	if event != nil {
		event()
		return
	}
	fmt.Println(msg)
}

func unsafeFormat(target, source string) string {
	return fmt.Sprintf("attribute.Unsafe: %s is unsafe; source: %s", target, source)
}

