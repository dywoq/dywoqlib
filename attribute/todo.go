package attribute

import "strings"

// Todo is used to output warning into console to user,
// who tries to use an unimplemented function.
// Does not returns zero values automatically (e.g., nil, 0, "").
//
// If mode is set to SoftMode, then program will be not executed.
// Otherwise it's set to StrictMode, program will be executed.
//
// Example:
// 	func NotImplemented() int {
//		attributes.Todo(attributes.SoftMode)
//		return 0
// 	}
func Todo(mode Mode) {
	target := functionName(1)
	source := functionName(2)
	message := strings.Join(
		[]string{
			"attributes.Todo: ",
			target,
			" is not implemented yet; source of the warning: ",
			source,
		}, "",
	)
	output(message, mode)
}
