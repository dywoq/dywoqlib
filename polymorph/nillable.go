package polymorph

import "reflect"

// Nillable checks if T can hold nil values.
func Nillable[T any]() bool {
	kind := KindOf[T]()
	switch kind {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return true
	default:
		return false
	}
}
