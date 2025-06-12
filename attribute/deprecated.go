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

import "strings"

// Deprecated is an attribute-function that generates a warning about deprecated function.
//   - name of the deprecated function.
//   - name of the source of the warning.
//
// If mode is set to:
//   - SoftMode - the warning will be just simply outputted into console.
//   - StrictMode - the warning will be outputted into console and cause program terminate.
func Deprecated(mode Mode) {
	m := management{}
	target := m.functionName(m.targetNumberSkip())
	source := m.functionName(m.sourceNumberSkip())
	elems := []string{
		"attribute.Deprecated: ",
		target,
		" is deprecated; the source of the warning: ", source,
	}
	message := strings.Join(elems, "")
	m.output(message, mode)
}
