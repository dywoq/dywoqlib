package container

// Fixed indicates a fixed-length slice wrapper.
// Fixed is a generic struct that wraps a Go slice (`[]T`) and associates it
// with a predetermined `initialLength`. This `initialLength` acts as a cap
// for the maximum allowed size of the underlying `data` slice.
type Fixed[T any] struct {
	initialLength int
	data          []T
}

// NewFixed returns new instance of Fixed.
// Panics if the actual length of data is out of the initial length.
func NewFixed[T any](initialLength int, data []T) *Fixed[T] {
	if len(data) >= initialLength {
		Panic("the actual size of data is out of the initial length")
	}
	return &Fixed[T]{initialLength, data}
}
