package iterator

type NativeGetter[T any] interface {
	Native() T
}
