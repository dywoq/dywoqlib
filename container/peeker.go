package container

// Peeker defines methods for inspecting the elements at the ends of a collection without removal.
type Peeker[T any] interface {
	Front() (T, error)
	Back() (T, error)
	Empty() bool
}
