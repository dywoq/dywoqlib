package stringn

import (
	"fmt"

	"github.com/dywoq/dywoqlib/iterator"
)

// Fixed represents a string with a fixed length constraint.
// It wraps a Dynamic string and enforces that operations only succeed
// if the underlying string matches the specified fixed length.
type Fixed struct {
	str         string
	err         error
	fixedLength int
	d           *Dynamic
}

// NewFixed creates a new Fixed string with the given fixed length and initial value.
func NewFixed(fixedLength int, str string) *Fixed {
	return &Fixed{str, nil, fixedLength, NewDynamic(str)}
}

// Length returns the length of the underlying string.
func (f *Fixed) Length() int {
	return f.d.Length()
}

// FixedLength returns the fixed length constraint for this Fixed string.
func (f *Fixed) FixedLength() int {
	return f.fixedLength
}

// Empty reports whether the underlying string is empty.
func (f *Fixed) Empty() bool {
	return f.d.Empty()
}

// Begin returns an iterator to the beginning of the string.
func (f *Fixed) Begin() *iterator.Iterator[rune] {
	return f.d.Begin()
}

// End returns an iterator to the end of the string.
func (f *Fixed) End() *iterator.Iterator[rune] {
	return f.d.End()
}

// Find searches for the specified rune in the string.
// Returns zero value if not found or if an error occurs.
func (f *Fixed) Find(reqRune rune) rune {
	if f.d.Err() != nil {
		f.err = f.d.Err()
		var zero rune
		return zero
	}
	if f.d.Length() != f.fixedLength {
		f.err = ErrOutOfFixedLength
		var zero rune
		return zero
	}
	r := f.d.Find(reqRune)
	if f.d.Err() != nil {
		f.err = f.d.Err()
		var zero rune
		return zero
	}
	return r
}

// At returns the rune at the specified index.
// Returns zero value if out of bounds or if an error occurs.
func (f *Fixed) At(i int) rune {
	if f.d.Err() != nil {
		f.err = f.d.Err()
		var zero rune
		return zero
	}
	if f.d.Length() != f.fixedLength {
		f.err = ErrOutOfFixedLength
		var zero rune
		return zero
	}
	r := f.d.At(i)
	if f.d.Err() != nil {
		f.err = f.d.Err()
		var zero rune
		return zero
	}
	return r
}

// Format implements custom formatting for the Fixed string.
func (f *Fixed) Format(state fmt.State, verb rune) {
	// default
	f.d.Format(state, verb)
	// custom
	if verb == 'r' && state.Flag('#') {
		fmt.Fprintf(state, "%s:%d", f.str, f.fixedLength)
	}
}

// Front returns the first rune of the string.
// Returns zero value if the string is empty or an error occurs.
func (f *Fixed) Front() rune {
	if f.d.Err() != nil {
		f.err = f.d.Err()
		var zero rune
		return zero
	}
	if f.d.Length() != f.fixedLength {
		f.err = ErrOutOfFixedLength
		var zero rune
		return zero
	}
	return f.d.Front()
}

// Back returns the last rune of the string.
// Returns zero value if the string is empty or an error occurs.
func (f *Fixed) Back() rune {
	if f.d.Err() != nil {
		f.err = f.d.Err()
		var zero rune
		return zero
	}
	if f.d.Length() != f.fixedLength {
		f.err = ErrOutOfFixedLength
		var zero rune
		return zero
	}
	return f.d.Back()
}

// AppendBack appends the provided strings to the end of the underlying string.
// Returns the resulting string, or an empty string if an error occurs.
func (f *Fixed) AppendBack(strs ...string) string {
	if f.d.Err() != nil {
		f.err = f.d.Err()
		return ""
	}
	if f.d.Length() != f.fixedLength {
		f.err = ErrOutOfFixedLength
		return ""
	}
	res := f.d.AppendBack(strs...)
	if f.d.Err() != nil {
		f.err = f.d.Err()
		return ""
	}
	return res
}

// Append appends the provided strings to the end of the underlying string.
// Does nothing if an error occurs.
func (f *Fixed) Append(strs ...string) {
	if f.d.Err() != nil {
		f.err = f.d.Err()
		return
	}
	if f.d.Length() != f.fixedLength {
		f.err = ErrOutOfFixedLength
		return
	}
	_ = f.AppendBack(strs...)
	if f.d.Err() != nil {
		f.err = f.d.Err()
		return
	}
}

// EraseBack erases the underlying string and returns its previous value.
// Returns an empty string if an error occurs.
func (f *Fixed) EraseBack() string {
	if f.d.Err() != nil {
		f.err = f.d.Err()
		return ""
	}
	if f.d.Length() != f.fixedLength {
		f.err = ErrOutOfFixedLength
		return ""
	}
	res := f.d.EraseBack()
	if f.d.Err() != nil {
		f.err = f.d.Err()
		return ""
	}
	return res
}

// Erase erases the underlying string. Does nothing if an error occurs.
func (f *Fixed) Erase() {
	if f.d.Err() != nil {
		f.err = f.d.Err()
		return
	}
	if f.d.Length() != f.fixedLength {
		f.err = ErrOutOfFixedLength
		return
	}
	_ = f.EraseBack()
	if f.d.Err() != nil {
		f.err = f.d.Err()
		return
	}
}
