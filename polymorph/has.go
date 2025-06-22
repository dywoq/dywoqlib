package polymorph

import "reflect"

// HasMethod checks if T has a method with the given name.
func HasMethod[T any](name string) bool {
	_, found := TypeOfGeneric[T]().MethodByName(name)
	return found
}

// HasMethod checks if struct S has a method with the given name.
// Returns false if S is not a structure.
func HasField[S any](name string) bool {
	tType := TypeOfGeneric[S]()
	if tType.Kind() != reflect.Struct {
		return false
	}
	_, found := tType.FieldByName(name)
	return found
}
