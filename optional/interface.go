package optional

// Optional is an interface representing a optional value.
type Optional[T any] interface {
	// Present checks if the optional value is present.
	Present() bool
	// Get returns the optional value and a boolean indicating its presence.
	Get() (T, bool)
}

// New retruns a new Optional with a value of a generic parameter T.
func New[T any](val T) Optional[T] {
	return &implementation[T]{val, true}
}

// Empty creates a new Optional with no value, 
// but a generic parameter T must be still present.
func Empty[T any]() Optional[T] {
	return &implementation[T]{present: false}
}
