package slice

func MergeDynamic[T comparable](first *Dynamic[T], second *Dynamic[T]) *Dynamic[T] {
	new := NewDynamic[T]()
	if new.Error() != nil {
		return nil
	}
	new.Grow(first.Length() + second.Length())

	it := first.Iterating().Forward()
	for it.Next() {
		new.Append(it.Value())
	}

	it = second.Iterating().Forward()
	for it.Next() {
		new.Append(it.Value())
	}

	return new
}

func MergeFixed[T comparable](first *Fixed[T], second *Fixed[T]) *Fixed[T] {
	new := NewFixed[T](first.Length() + second.Length())
	if new.Error() != nil {
		return nil
	}

	it := first.Iterating().Forward()
	for it.Next() {
		new.Append(it.Value())
	}

	it = second.Iterating().Forward()
	for it.Next() {
		new.Append(it.Value())
	}

	return new
}
