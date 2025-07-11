package sliceutil

import (
	"fmt"
	"slices"
	"strings"

	"github.com/dywoq/dywoqlib/iterator"
)

func Format[T comparable](s []T) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	var b strings.Builder
	for i, elem := range s {
		_, err := fmt.Fprintf(&b, "%v", elem)
		if err != nil {
			return "", err
		}
		if i != len(s)-1 {
			_, err = fmt.Fprintf(&b, ", ")
			if err != nil {
				return "", err
			}
		}
	}
	return b.String(), nil
}

func Find[T comparable](req T, it *iterator.Forward[T]) (T, error) {
	var val T
	for it.Next() {
		val = it.Value()
		if val == req {
			return val, nil
		}
	}
	if it.Error() != nil {
		return val, it.Error()
	}
	return val, ErrElementNotFound
}

func At[T comparable](i int, s []T) (T, error) {
	if i < 0 || i >= len(s) {
		var zero T
		return zero, ErrWrongIndex
	}
	return s[i], nil
}

func Set[T comparable](elem T, i int, s []T) (T, error) {
	if i < 0 || i >= len(s) {
		var zero T
		return zero, ErrWrongIndex
	}
	old := s[i]
	s[i] = elem
	return old, nil
}

func Delete[T comparable](i int, s []T) (T, error) {
	if i < 0 || i >= len(s) {
		var zero T
		return zero, ErrWrongIndex
	}
	removed := s[i]
	s = slices.Delete(s, i, i+1)
	return removed, nil
}

func Insert[T comparable](i int, elem T, s []T) (T, error) {
	if i < 0 || i >= len(s) {
		var zero T
		return zero, ErrWrongIndex
	}
	inserted := s[i]
	s = slices.Insert(s, i, elem)
	return inserted, nil
}
