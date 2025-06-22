package polymorph

import "reflect"

// NumMethods returns the number of the T methods.
// Inside the function, it uses reflect.Type.NumMethod().
// See https://pkg.go.dev/reflect#Value.NumMethod for more.
func NumMethods[T any]() int {
	return TypeOfGeneric[T]().NumMethod()
}

// NumFields returns the number of the T fields.
// Returns zero if S is not structure,
// as it's required by https://pkg.go.dev/reflect#Value.NumField.
func NumFields[S any]() int {
	if KindOf[S]() != reflect.Struct {
		return 0
	}
	return TypeOfGeneric[S]().NumField()
}
