package container

type NativeGetter[T comparable] interface {
	Native() []T
}
