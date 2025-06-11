// Under Apache License 2.0, see https://github.com/dywoq/dywoqlib/blob/main/LICENSE for more information.
// dywoq - 2025 year

package attribute

import "strings"

// Todo is an attribute-function that generates a warning,
// which indicates that a function is not implemented yet.
// The warning includes:
//   - name of the not implemented function.
//   - name of the source of the warning.
//
// If mode is set to:
//   - SoftMode - the warning will be just simply outputted into console.
//   - StrictMode - the warning will be outputted into console and cause program terminate.
func Todo(mode Mode) {
	m := management{}
	target := m.functionName(m.targetNumberSkip())
	source := m.functionName(m.sourceNumberSkip())
	elems := []string{
		"attribute.Todo: ",
		target,
		" is not implemented yet; the source of the warning: ", source,
	}
	message := strings.Join(elems, "")
	m.output(message, mode)
}
