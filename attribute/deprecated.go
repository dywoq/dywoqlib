package attribute

import "strings"

// Deprecated gives a warning to user of a function, which indicates
// the function is deprecated.
//
// If mode is set to SoftMode, then program will be not executed.
// Otherwise it's set to StrictMode, program will be executed.
//
// Example:
// 	func Deprecated() int {
//		attributes.Deprecated(attributes.SoftMode)
//		return 0
// 	}
func Deprecated(mode Mode) {
	target := functionName(1)
	source := functionName(2)
	message := strings.Join(
		[]string{
			"attributes.Deprecated: ",
			target,
			" is deprecated; source of the warning: ",
			source,
		}, "",
	)
	output(message, mode)
}

// DeprecatedN works the exactly same as Deprecated,
// but it allows to point at the new function of the deprecated version.
//
// Example:
// 	func Deprecated() int {
//		attributes.Deprecated(attributes.SoftMode)
//		return 0
// 	}
func DeprecatedN(newFunc string, mode Mode) {
	target := functionName(1)
	source := functionName(2)
	message := strings.Join(
		[]string{
			"attributes.DeprecatedN: ",
			target,
			" is deprecated; use ", newFunc, ". source of the warning: ",
			source,
		}, "",
	)
	output(message, mode)
}
