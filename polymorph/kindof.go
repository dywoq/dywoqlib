package polymorph

import "reflect"

// KindOf returns the kind of the generic parameter T.
func KindOf[T any]() reflect.Kind {
	tTypeOf := TypeOfGeneric[T]()
	return tTypeOf.Kind()
}
