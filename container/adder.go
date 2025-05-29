package container

// ElementAdder defines a method to add a element to a collection.
type ElementAdder[T any] interface {
	Add(val T) error
}
