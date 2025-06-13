package polymorphic

import "reflect"

// IsKind checks if T is equivalent to kind.
// The function doesn't panics.
func IsKind[T any](kind reflect.Kind) bool {
	t := TypeOfGeneric[T]()
	return t.Kind() == kind
}
