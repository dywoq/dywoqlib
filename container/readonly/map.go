package readonly

import "github.com/dywoq/dywoqlib/err"

// Map is a generic thread-safe and readonly map container, 
// with K as the key type and V as the value type.
type Map[K, V comparable] struct{}

// Length returns the length of the underlying map.
// The method returns zero if error is present.
// The mutex locks and unlock after the completion of the function.
func (m *Map[K, V]) Length() int

// Error returns the possibly encountered error,
// otherwise it returns err.NoneContext.
// The mutex locks and unlock after the completion of the function.
func (m *Map[K, V]) Error() err.Context

// Exists reports whether reqkey exists.
// If error is present, it returns false.
// The mutex locks and unlock after the completion of the function.
func (m *Map[K, V]) Exists(reqkey K) bool

// Keys returns a slice of keys of the underlying map.
// If error is present, it returns empty slice.
// The mutex locks and unlock after the completion of the function.
func (m *Map[K, V]) Keys() []K

// Values returns a slice of values of the underlying map.
// If error is present, it returns empty slice.
// The mutex locks and unlock after the completion of the function.
func (m *Map[K, V]) Values() []V

// Get gets reqkey and returns reqkey and its value from the underlying map.
// If the method didn't find reqkey, it returns zero values.
// If error is present, it returns zero values.
// The mutex locks and unlock after the completion of the function. 
func (m *Map[K, V]) Get(reqkey K) (K, V)

// String returns a string representation of the underlying map.
// It returns an empty string if error is present.
// The mutex locks and unlock after the completion of the function.
func (m *Map[K, V]) String() string
