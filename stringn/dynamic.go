package stringn

import (
	"fmt"
	"strings"

	"github.com/dywoq/dywoqlib/container/slice"
	"github.com/dywoq/dywoqlib/iterator"
	"github.com/dywoq/dywoqlib/sliceutil"
)

type Dynamic struct {
	str string
	err error
}

func NewDynamic(str string) *Dynamic {
	return &Dynamic{str, nil}
}

func (d *Dynamic) Length() int {
	return len(d.str)
}

func (d *Dynamic) Err() error {
	return d.err
}

func (d *Dynamic) Empty() bool {
	return len(d.str) == 0
}

func (d *Dynamic) Begin() *iterator.Iterator[rune] {
	it := iterator.New(0, []rune(d.str))
	if it.Err() != nil {
		d.err = it.Err()
		return nil
	}
	return it
}

func (d *Dynamic) End() *iterator.Iterator[rune] {
	it := iterator.New(len(d.str)-1, []rune(d.str))
	if it.Err() != nil {
		d.err = it.Err()
		return nil
	}
	return it
}

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

func (d *Dynamic) Append(strs ...string) {
	if d.err != nil {
		return
	}
	_ = d.AppendBack(strs...)
}

func (d *Dynamic) EraseBack() string {
	if d.err != nil {
		return ""
	}
	saved := d.str
	d.str = ""
	return saved
}

func (d *Dynamic) Erase() {
	if d.err != nil {
		return
	}
	_ = d.EraseBack()
}
