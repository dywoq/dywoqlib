package container

import (
	"fmt"
	"strings"

	"github.com/dywoq/dywoqlib/iterator"
)

type management[T comparable] struct {
	it  iterator.Iterable[T]
	err error
}

func (m *management[T]) currentErr() error {
	return m.err
}

func (m *management[T]) find(reqElem T) T {
	if m.err != nil {
		return m.zero()
	}

	it := m.it.Begin()
	if it.Native() == nil {
		return m.zero()
	}

	found := false
	var elem T
	for it.Next() {
		if it.Value() == reqElem {
			found = true
			elem = it.Value()
			break
		}
	}

	if it.Err() != nil {
		m.err = it.Err()
		return m.zero()
	}

	if !found {
		m.err = ErrElementNotFound
		return m.zero()
	}

	m.err = nil
	return elem
}

func (m management[T]) at(i int) T {
	if m.err != nil {
		return m.zero()
	}

	it := m.it.Begin()
	n := it.Native()
	if n == nil {
		m.err = ErrSliceIsNil
		return m.zero()
	}

	if i < 0 || i >= len(n) {
		m.err = ErrIndexOutOfBounds
		return m.zero()
	}

	m.err = nil
	return n[i]
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

func (m management[T]) zero() T {
	var zero T
	return zero
}
