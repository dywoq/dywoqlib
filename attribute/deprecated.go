// Under Apache License 2.0, see https://github.com/dywoq/dywoqlib/blob/main/LICENSE for more information.
// dywoq - 2025 year

package attribute

import "strings"

// Deprecated is an attribute-function that generates a warning about deprecated function.
// 	- name of the deprecated function.
// 	- name of the source of the warning.
//
// If mode is set to:
// 	- SoftMode - the warning will be just simply outputted into console.
// 	- StrictMode - the warning will be outputted into console and cause program terminate.
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
