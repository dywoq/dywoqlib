package slice

// LengthChecker interface for types that provide information about their length and capacity.
type LengthChecker[T comparable] interface {
	InitialLength() int
	ActualLength() int
	Empty() bool
	OverInitialLength() bool
	Negative() bool
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
