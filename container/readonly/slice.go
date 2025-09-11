package readonly

import (
	"github.com/dywoq/dywoqlib/err"
)

// Slice is a generic readonly wrapper around the standard Go slice.
type Slice[T comparable] struct{}

// Error returns the possibly encountered current error.
// If error doesn't present, the function returns err.NoneContext().
// The mutex locks and unlocks after the function completed.
func (s *Slice[T]) Error() err.Context

// Length returns the length of the underlying slice
// If error is present, it returns zero.
// The mutex locks and unlocks after the function completed.
func (s *Slice[T]) Length() int

// At returns the element at i.
// If error is present, it returns zero value.
// The mutex locks and unlocks after the function completed.
func (s *Slice[T]) At(i int) T

// Find finds req in the underlying slice
// and returns it if the finding was successful, otherwise,
// it updates the internal error state.
// If error is present, it returns zero value.
// The mutex locks and unlocks after the function completed.
func (s *Slice[T]) Find(req T) T

// String returns a string representation of the slice.
// It uses sliceutil.Format to format the underlying slice.
// Updates the error state if sliceutil.Format returned the error.
// The mutex locks and unlocks after the function completed.
func (s *Slice[T]) String() string

// Front returns the first element of the slice.
// It returns a zero value if the slice is empty or an error occurred.
// The mutex locks and unlocks after the function completed.
func (s *Slice[T]) Front() T

// Back returns the last element of the slice.
// It returns a zero value if the slice is empty or an error occurred.
// The mutex locks and unlocks after the function completed.
func (s *Slice[T]) Back() T 
