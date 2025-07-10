package sliceutil

import (
	"fmt"
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

func Find[T comparable](req T, it iterator.Forward[T]) (T, error) {
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
