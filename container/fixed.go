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
		Panic(ErrOffTheInitialLength.Error())
	}
	return &Fixed[T]{initialLength, data}
}

// OffInitialLength checks if the fixed-length slice is off the initial length.
func (f Fixed[T]) OffInitialLength() bool {
	return f.ActualLength() >= f.InitialLength()
}

// InitialLength returns the initial length.
func (f Fixed[T]) InitialLength() int {
	return f.initialLength
}

// ActualLength returns the actual length of the fixed-length slice.
func (f Fixed[T]) ActualLength() int {
	return len(f.data)
}

// Empty checks if the fixed-length slice is not empty,
func (f Fixed[T]) Empty() bool {
	return f.ActualLength() == 0
}

// Returns the first element of the slice.
func (f Fixed[T]) Front() T {
	if f.OffInitialLength() {
		Panic(ErrOffTheInitialLength.Error())
	}

	if f.Empty() {
		var zero T
		return zero
	}
	return f.data[0]
}

// Returns the last element of the slice.
func (f Fixed[T]) Back() T {
	if f.OffInitialLength() {
		Panic(ErrOffTheInitialLength.Error())
	}

	if f.Empty() {
		var zero T
		return zero
	}
	return f.data[f.ActualLength()]
}
