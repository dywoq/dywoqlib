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

// DirectDataGetter defines a method for getting an private slice directly.
type DirectDataGetter[T any] interface {
	Direct() []T
}

// Cleaner defines a method to clear a whole collection.
type Cleaner interface {
	Clear()
}

// FullnessChecker defines methods for checking if a container is full.
type FullnessChecker interface {
	IsFull() bool
}

// InitialLengthChecker defines a method that checks if a container is over the initial length.
type InitialLengthChecker interface {
	IsOverInitialLength() bool
}

// Filterer defines a method that filters elements of the S container,
// and returns the new S container instance that holds chosen elements
// that satisfy predicate.
type Filterer[T any, S any] interface {
	Filter(predicate func(T) bool) *S
}

// Replacer defines a method that replaces newVal over oldVal.
// The method returns the number of replacements made.
type Replacer[T any] interface {
	Replace(oldVal, newVal T) int
}
