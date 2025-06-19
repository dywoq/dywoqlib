package polymorph

import "reflect"

// Implements checks if S implements I.
// Returns false if I is not an interface.
func Implements[I any, S any]() bool {
	sType := TypeOfGeneric[S]()
	iType := TypeOfGeneric[I]()
	if iType.Kind() != reflect.Interface {
		return false
	}
	return sType.Implements(iType)
}
