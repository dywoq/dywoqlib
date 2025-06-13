package slice

// ErrorChecker interface for types that can report an error state.
type ErrorChecker interface {
	Err() error
}

// LengthChecker interface for types that provide information about their length and capacity.
type LengthChecker interface {
	ActualLength() int
	Empty() bool
}

// OverInitialLengthChecker interface for types that can check if a slice is over the initial length.
type OverInitialLengthChecker interface {
	OverInitialLength() bool
}

// NegativeLengthChecker interface for types that can check if an initial length is negative.
type NegativeLengthChecker interface {
	Negative() bool
}

// InitialLengthGetter interface for types that provide information about their initial length.
type InitialLengthGetter interface {
	InitialLength() int
}

// ElementReader interface for types whose elements can be read.
type ElementReader[T comparable] interface {
	At(index int) T
	Front() T
	Back() T
	Contains(t T) bool
}

// ElementModifier interface for types whose elements can be modified, appended, or removed.
type ElementModifier[T comparable] interface {
	Set(index int, value T)
	Append(elements ...T)
	Remove(index int)
	Clear()
}
