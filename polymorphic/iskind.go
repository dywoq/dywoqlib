// Under Apache License 2.0, see https://github.com/dywoq/dywoqlib/blob/main/LICENSE for more information.
// dywoq - 2025 year

package polymorphic

import "reflect"

// IsKind checks if T is equivalent to kind.
// The function doesn't panics.
func IsKind[T any](kind reflect.Kind) bool {
	t := TypeOfGeneric[T]()
	return t.Kind() == kind
}
