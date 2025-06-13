package polymorphic

import "reflect"

// Implements checks if I is implemented by S.
// Also, the function can return false if I is not an interface.
// The function doesn't panics.
func Implements[I any, S any]() bool {
	if !IsKind[I](reflect.Interface) {
		return false
	}
	tInterface := TypeOfGeneric[I]()
	tStruct := TypeOfGeneric[S]()
	return tStruct.Implements(tInterface)
}
