package iterator

// MapIterable represents an iterable collection that can produce a Map iterator.
// It defines methods to get the beginning and ending iterators for a map-like structure.
type MapIterable[K comparable, V comparable] interface {
	Begin() *Map[K, V]
	End() *Map[K, V]
}

// SliceIterable represents an iterable collection that can produce a Slice iterator.
// It defines methods to get the beginning and ending iterators for a slice-like structure.
type SliceIterable[T any] interface {
	Begin() *Slice[T]
	End() *Slice[T]
}
