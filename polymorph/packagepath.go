package polymorph

// PackagePath returns the package path of T.
func PackagePath[T any]() string {
	return TypeOfGeneric[T]().PkgPath()
}
