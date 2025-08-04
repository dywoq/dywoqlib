package optional

import "fmt"

// Maybe is an interface representing a optional value.
type Maybe[T any] interface {
	fmt.Stringer
	// Present checks if the optional value is present.
	Present() bool
	// Get returns the value and a boolean indicating its presence.
	// The boolean always match what Present() returns.
	Get() (T, bool)
	// Else returns the value if it's present, otherwise it returns a default value.
	Else(T) T
}

// New retruns a new Maybe with a value of a generic parameter T.
func New[T any](val T) Maybe[T] {
	return &implementation[T]{val, true}
}

// None creates a new Maybe with no value,
// but a generic parameter T must be still present.
func None[T any]() Maybe[T] {
	return &implementation[T]{present: false}
}
