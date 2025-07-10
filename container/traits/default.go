package traits

import (
	"fmt"
	"strings"

	"github.com/dywoq/dywoqlib/iterator"
)

type Default[T comparable] struct {
	err error
}

func (d *Default[T]) Error() error {
	return d.err
}

func (d *Default[T]) Format(s []T) string {
	if d.err != nil {
		return ""
	}
	if len(s) == 0 {
		return ""
	}
	var b strings.Builder
	for i, elem := range s {
		_, err := fmt.Fprintf(&b, "%v", elem)
		if err != nil {
			d.err = err
			break
		}
		if i != len(s)-1 {
			_, err = fmt.Fprintf(&b, ", ")
			if err != nil {
				d.err = err
				break
			}
		}
	}
	return b.String()
}

func (d *Default[T]) Find(req T, it iterator.Forward[T]) T {
	var val T
	if d.err != nil {
		return val
	}
	for it.Next() {
		val = it.Value()
		if val == req {
			return val
		}
	}
	if it.Error() != nil {
		d.err = it.Error()
		return val
	}
	d.err = ErrElementNotFound
	return val
}

func (d *Default[T]) At(i int, s []T) T {
	if d.err != nil {
		var zero T
		return zero
	}
	if i < 0 || i >= len(s) {
		d.err = ErrWrongIndex
		var zero T
		return zero
	}
	return s[i]
}
