package traits

import "github.com/dywoq/dywoqlib/iterator"

type Base[T comparable] interface {
	Error() error
	Format(s []T) string
	Find(req T, it iterator.Forward[T]) T
	At(i int, s []T) T
}
