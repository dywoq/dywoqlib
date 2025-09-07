package constraints

// Complexable is a constraint of complex integral types (complex64, complex128).
type Complexable interface {
	~complex64 | ~complex128
}
