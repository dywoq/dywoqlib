package slice

// MergeDynamic merges two dynamic slices of comparable generic parameter T
// and returns a pointer to Dynamic[T].
// If any errors are encountered, it returns a nil instead of the pointer to Dynamic[T] 
// with the first encountered error.
func MergeDynamic[T comparable](first *Dynamic[T], second *Dynamic[T]) (*Dynamic[T], error)  {
	new := NewDynamic[T]()
	if new.Error() != nil {
		return nil, new.Error()
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

	return new, nil
}

// MergeFixed merged two fixed-length slices of comparable generic parameter T
// and returns a pointer to Fixed[T].
// If any errors are encountered, it returns a nil instead of the pointer to Fixed[T]
// with the first encountered error.
func MergeFixed[T comparable](first *Fixed[T], second *Fixed[T]) (*Fixed[T], error) {
	new := NewFixed[T](first.Length() + second.Length())
	if new.Error() != nil {
		return nil, new.Error()
	}

	it := first.Iterating().Forward()
	for it.Next() {
		new.Append(it.Value())
	}
	if it.Error() != nil {
		return nil, it.Error()
	}

	it = second.Iterating().Forward()
	for it.Next() {
		new.Append(it.Value())
	}
	if it.Error() != nil {
		return nil, it.Error()
	}

	return new, nil
}
