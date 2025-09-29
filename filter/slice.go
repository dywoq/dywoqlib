package filter

// Slice returns a filtered slice of elements that satisfy pred.
// Returns an empty slice if len(s) is 0.
func Slice[S any](s []S, pred func(S) bool) []S {
	if len(s) == 0 {
		return []S{}
	}
	result := []S{}
	for _, elem := range s {
		if pred(elem) {
			result = append(result, elem)
		}
	}
	return result
}

// SliceNot returns a filtered slice of elements that don't satisfy pred.
// Returns an empty slice if len(s) is 0.
func SliceNot[S any](s []S, pred func(S) bool) []S {
	return Slice(s, func(s S) bool { return !pred(s) })
}
