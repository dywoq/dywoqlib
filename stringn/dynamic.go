package stringn

import (
	"fmt"
	"strings"

	"github.com/dywoq/dywoqlib/container/slice"
	"github.com/dywoq/dywoqlib/iterator"
	"github.com/dywoq/dywoqlib/sliceutil"
)

// Dynamic represents a mutable string with various utility methods.
type Dynamic struct {
	str string
	err error
}

// NewDynamic creates a new Dynamic string with the given initial value.
func NewDynamic(str string) *Dynamic {
	return &Dynamic{str, nil}
}

// Length returns the length of the string.
func (d *Dynamic) Length() int {
	return len(d.str)
}

// Err returns the last error encountered by the Dynamic string.
func (d *Dynamic) Err() error {
	return d.err
}

// Empty reports whether the string is empty.
func (d *Dynamic) Empty() bool {
	return len(d.str) == 0
}

// Begin returns an iterator to the beginning of the string.
func (d *Dynamic) Begin() *iterator.Iterator[rune] {
	it := iterator.New(0, []rune(d.str))
	if it.Err() != nil {
		d.err = it.Err()
		return nil
	}
	return it
}

// End returns an iterator to the end of the string.
func (d *Dynamic) End() *iterator.Iterator[rune] {
	it := iterator.New(len(d.str)-1, []rune(d.str))
	if it.Err() != nil {
		d.err = it.Err()
		return nil
	}
	return it
}

// Find searches for the specified rune in the string.
// Returns zero value if not found or if an error occurs.
func (d *Dynamic) Find(reqRune rune) rune {
	if d.err != nil {
		var zero rune
		return zero
	}
	m := sliceutil.Management[rune]{}
	m.SetIterableType(d)
	foundElem := m.Find(reqRune)
	if m.Err() != nil {
		d.err = m.Err()
		var zero rune
		return zero
	}
	return foundElem
}

// At returns the rune at the specified index.
// Returns zero value if out of bounds or if an error occurs.
func (d *Dynamic) At(i int) rune {
	if d.err != nil {
		var zero rune
		return zero
	}
	m := sliceutil.Management[rune]{}
	m.SetIterableType(d)
	foundElem := m.At(i)
	if m.Err() != nil {
		d.err = m.Err()
		var zero rune
		return zero
	}
	return foundElem
}

// Format implements custom formatting for the Dynamic string.
func (d *Dynamic) Format(state fmt.State, verb rune) {
	// outputting string
	if verb == 'v' || verb == 's' {
		fmt.Fprintln(state, d.str)
	}

	// outputting the length of string
	if verb == 'd' {
		fmt.Fprintln(state, d.Length())
	}

	// outputting runes
	if verb == 'r' {
		fixed := slice.NewFixed(len(d.str), []rune(d.str)...)
		fmt.Fprintln(state, fixed)
	}

	fmt.Fprintln(state, d.str)
}

// Front returns the first rune of the string.
// Returns zero value if the string is empty or an error occurs.
func (d *Dynamic) Front() rune {
	if d.err != nil {
		var zero rune
		return zero
	}
	if d.Length() == 0 {
		d.err = ErrEmpty
		var zero rune
		return zero
	}
	return d.At(0)
}

// Back returns the last rune of the string.
// Returns zero value if the string is empty or an error occurs.
func (d *Dynamic) Back() rune {
	if d.err != nil {
		var zero rune
		return zero
	}
	if d.Length() == 0 {
		d.err = ErrEmpty
		var zero rune
		return zero
	}
	return d.At(d.Length() - 1)
}

// AppendBack appends the provided strings to the end of the string.
// Returns the resulting string, or an empty string if an error occurs.
func (d *Dynamic) AppendBack(strs ...string) string {
	if d.err != nil {
		return ""
	}
	var b strings.Builder
	for _, str := range strs {
		b.WriteString(str)
	}
	d.str += b.String()
	return d.str
}

// Append appends the provided strings to the end of the string.
// Does nothing if an error occurs.
func (d *Dynamic) Append(strs ...string) {
	if d.err != nil {
		return
	}
	_ = d.AppendBack(strs...)
}

// EraseBack erases the string and returns its previous value.
// Returns an empty string if an error occurs.
func (d *Dynamic) EraseBack() string {
	if d.err != nil {
		return ""
	}
	saved := d.str
	d.str = ""
	return saved
}

// Erase erases the string. Does nothing if an error occurs.
func (d *Dynamic) Erase() {
	if d.err != nil {
		return
	}
	_ = d.EraseBack()
}
