package numeric

// Numeric is a constraint of the integral, floating types (int, uint etc.).
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Integral is a constraint of the integral types.
type Integral interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Complexable is a constraint of complex integral types (complex64, complex128).
type Complexable interface {
	~complex64 | ~complex128
}
