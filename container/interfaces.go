package container

// ElementAdder defines a method to add an element to a collection.
type ElementAdder[T any] interface {
	Add(val T) error
}

// ElementGetter defines a method to retrieve an element from a collection.
type ElementGetter[T any] interface {
	Get(index int) (T, error)
}

// ElementSetter defines a method to set an element to a collection at the index.
type ElementSetter[T any] interface {
	Set(index int, val T) error 	
}

// Peeker defines methods for inspecting the elements at the ends of a collection without removal.
type Peeker[T any] interface {
	Front() (T, error)
	Back() (T, error)
	Empty() bool
}

// ActualLengthGetter defines a method for getting the actual length.
type ActualLengthGetter interface {
	ActualLength() int
}

// InitialLengthGetter defines a method for getting the initial length.
// Usually it's implemented by Fixed structure.
type InitialLengthGetter interface {
	InitialLength() int
}
