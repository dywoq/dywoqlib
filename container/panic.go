package container

import "strings"

// Panic generates a panic with a formatted message.
// It's not generally recommended to use this function as a user.
// The function is defined as public since there will be conflicts with naming, as panic function already exists.
func Panic(message string) {
	result := strings.Join([]string{"github.com/dywoq/dywoqlib/container: ", message}, "")
	panic(result)
}
