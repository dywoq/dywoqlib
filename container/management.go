package container

import (
	"fmt"
	"strings"

	"github.com/dywoq/dywoqlib/container/iterator"
)

type management[T comparable] struct {
	it  iterator.Iterable[T]
	err error
}

func (m *management[T]) currentErr() error {
	return m.err
}

func (m *management[T]) find(reqElem T) (elem T) {
	it := m.it.Begin()
	for it.Next() {
		if it.Value() != reqElem {
			var zero T
			elem = zero
			m.err = ErrElementNotFound
			break
		}
		elem = it.Value()
	}
	return
}

func (m management[T]) at(i int) (elem T) {
	it := m.it.Begin()
	for it.Next() {
		if it.Position() != i {
			var zero T
			elem = zero
			m.err = ErrElementNotFound
			break
		}
		elem = it.Value()
	}
	return
}

// used to implement fmt.Stringer interface
func (m management[T]) format() string {
	var b strings.Builder
	it := m.it.Begin()
	b.WriteString("[")
	for it.Next() {
		b.WriteString(fmt.Sprintf("%v", it.Value()))
		if it.Position() != len(it.Native())-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString("]")
	return b.String()
}
