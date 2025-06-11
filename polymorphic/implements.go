package polymorphic

import "reflect"

// Implements checks if I is implemented by S.
// Also, the function can return false if I is not an interface.
func Implements[I any, S any]() bool {
	if !IsKind[I](reflect.Interface) {
		return false
	}
	tInterface := reflect.TypeOf((*I)(nil)).Elem()
	tStruct := reflect.TypeOf((*S)(nil)).Elem()
	return tStruct.Implements(tInterface)
}
