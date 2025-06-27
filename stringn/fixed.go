package stringn

import (
	"fmt"

	"github.com/dywoq/dywoqlib/iterator"
)

type Fixed struct {
	str         string
	err         error
	fixedLength int
	d           *Dynamic
}

func NewFixed(fixedLength int, str string) *Fixed {
	return &Fixed{str, nil, fixedLength, NewDynamic(str)}
}

func (f *Fixed) Length() int {
	return f.d.Length()
}

func (f *Fixed) FixedLength() int {
	return f.fixedLength
}

func (f *Fixed) Empty() bool {
	return f.d.Empty()
}

func (f *Fixed) Begin() *iterator.Iterator[rune] {
	return f.d.Begin()
}

func (f *Fixed) End() *iterator.Iterator[rune] {
	return f.d.End()
}

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

func (f *Fixed) Format(state fmt.State, verb rune) {
	// default
	f.d.Format(state, verb)
	// custom
	if verb == 'r' && state.Flag('#') {
		fmt.Fprintf(state, "%s:%d", f.str, f.fixedLength)
	}
}

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
