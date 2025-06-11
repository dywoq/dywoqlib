package polymorphic

import "reflect"

// TypeOfGeneric returns the reflect.Type of the generic type parameter T.
func TypeOfGeneric[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}
