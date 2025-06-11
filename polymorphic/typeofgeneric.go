// Under Apache License 2.0, see https://github.com/dywoq/dywoqlib/blob/main/LICENSE for more information.
// dywoq - 2025 year

package polymorphic

import "reflect"

// TypeOfGeneric returns the reflect.Type of the generic type parameter T.
func TypeOfGeneric[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}
