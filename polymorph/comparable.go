package polymorph

// Comparable checks for the comparability of T.
// The generic type parameter T is to be presented.
func Comparable[T any]() bool {
	return TypeOfGeneric[T]().Comparable()
}
