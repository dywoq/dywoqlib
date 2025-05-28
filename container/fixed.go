package container

// Fixed indicates a fixed-length slice wrapper.
// Fixed is a generic struct that wraps a Go slice (`[]T`) and associates it
// with a predetermined `initialLength`. This `initialLength` acts as a cap
// for the maximum allowed size of the underlying `data` slice.
type Fixed[T any] struct {
	initialLength int
	data          []T
}

// NewFixed returns a new instance of Fixed.
// Returns an error if initialLength is negative or if the actual length of
// data is greater than the initialLength.
func NewFixed[T any](initialLength int, data []T) (*Fixed[T], error) {
	if initialLength < 0 {
		return nil, ErrNegativeInitialLength
	}
	if len(data) > initialLength {
		return nil, ErrOffTheInitialLength
	}
	return &Fixed[T]{initialLength, data}, nil
}

// IsOverCapacity checks if the fixed-length slice has exceeded its initial length (capacity).
// This indicates an invalid state where the slice has grown beyond its intended maximum size.
func (f Fixed[T]) IsOverCapacity() bool {
	return f.ActualLength() > f.InitialLength()
}

// IsFull checks if the fixed-length slice has reached its initial length.
func (f Fixed[T]) IsFull() bool {
	return f.ActualLength() == f.InitialLength()
}

// InitialLength returns the initial length (capacity) of the fixed-length slice.
func (f Fixed[T]) InitialLength() int {
	return f.initialLength
}

// ActualLength returns the current number of elements in the fixed-length slice.
func (f Fixed[T]) ActualLength() int {
	return len(f.data)
}

// Empty checks if the fixed-length slice is empty.
func (f Fixed[T]) Empty() bool {
	return f.ActualLength() == 0
}

// Front returns the first element of the slice.
// Returns an error if the slice is empty.
func (f Fixed[T]) Front() (T, error) {
	if f.Empty() {
		var zero T
		return zero, ErrEmptyFixedSlice
	}

	if f.IsOverCapacity() {
		var zero T
		return zero, ErrOffTheInitialLength
	}
	return f.data[0], nil
}

// Back returns the last element of the slice.
// Returns an error if the slice is empty.
func (f Fixed[T]) Back() (T, error) {
	if f.Empty() {
		var zero T
		return zero, ErrEmptyFixedSlice
	}
	if f.IsOverCapacity() {
		var zero T
		return zero, ErrOffTheInitialLength
	}
	return f.data[f.ActualLength()-1], nil
}

// Add appends an element to the fixed-length slice.
// Returns an error if the slice is already full.
func (f *Fixed[T]) Add(val T) error {
	if f.IsFull() {
		return ErrFixedSliceFull
	}
	f.data = append(f.data, val)
	return nil
}

// Get returns the element at the specified index.
// Returns an error if the index is out of bounds.
func (f Fixed[T]) Get(index int) (T, error) {
	if index < 0 || index >= f.ActualLength() {
		var zero T
		return zero, ErrInvalidIndex
	}
	return f.data[index], nil
}

// Set sets the element at the specified index.
// Returns an error if the index is out of bounds.
func (f *Fixed[T]) Set(index int, val T) error {
	if index < 0 || index >= f.ActualLength() {
		return ErrInvalidIndex
	}
	f.data[index] = val
	return nil
}

// Native returns the current slice data.
func (f *Fixed[T]) Native() []T {
	return f.data
}
