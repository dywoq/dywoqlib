package iterator

// Slice is a generic iterator for slices of any type T.
type Slice[T any] struct {
	data []T
	pos  int
	err  error
}

// NewSlice creates a new Slice iterator for the given data.
// It initializes the iterator's position to -1, indicating it's before the first element.
func NewSlice[T any](data []T) *Slice[T] {
	return &Slice[T]{
		data: data,
		pos:  -1,
		err:  nil,
	}
}

// Err returns the first error encountered during iteration, if any.
// This allows consumers to check for errors after the iteration is complete.
func (s *Slice[T]) Err() error {
	return s.err
}

// Position returns the current index of the iterator.
// It returns -1 if the iterator is exhausted or an error has occurred.
func (s *Slice[T]) Position() int {
	if s.err != nil || !s.isValidPosition(s.pos) {
		return -1
	}
	return s.pos
}

// Next advances the iterator to the next element.
// It returns true if there is a next element, or false if the iterator is exhausted or an error occurred.
func (s *Slice[T]) Next() bool {
	if s.err != nil {
		return false
	}

	s.pos++

	if !s.isValidPosition(s.pos) {
		return false
	}
	return true
}

// Value returns the element at the current iterator position.
// If the iterator is not at a valid position (e.g., before the first element, after the last, or an error occurred),
// it sets an ErrInvalidPosition and returns the zero value of type T.
func (s *Slice[T]) Value() T {
	if s.err != nil || !s.isValidPosition(s.pos) {
		return s.zeroValueAndSetError()
	}
	return s.data[s.pos]
}

func (s *Slice[T]) isValidPosition(pos int) bool {
	return pos >= 0 && pos < len(s.data)
}

func (s *Slice[T]) zeroValueAndSetError() T {
	var zero T
	if s.err == nil {
		s.err = ErrInvalidPosition
	}
	return zero
}
