package filter

// Map returns a filtered map of keys that satisfy pred.
// Returns an empty map if len(m) is 0.
func Map[K comparable, V any](m map[K]V, pred func(K, V) bool) map[K]V {
	if len(m) == 0 {
		return map[K]V{}
	}
	result := map[K]V{}
	for k, v := range m {
		if pred(k, v) {
			result[k] = v
		}
	}
	return result
}

// Map returns a filtered map of keys that don't satisfy pred.
// Returns an empty map if len(m) is 0.
func MapNot[K comparable, V any](m map[K]V, pred func(K, V) bool) map[K]V {
	return Map(m, func(k K, v V) bool { return !pred(k, v) })
}

// MapKeys returns a filtered map with the keys that satisfy pred.
// Returns an empty map if len(m) is 0.
func MapKeys[K comparable, V any](m map[K]V, pred func(K) bool) map[K]V {
	return Map(m, func(k K, v V) bool { return pred(k) })
}

// MapKeys returns a filtered map with the values, and their keys, that satisfy pred.
// Returns an empty map if len(m) is 0.
func MapValues[K comparable, V any](m map[K]V, pred func(V) bool) map[K]V {
	return Map(m, func(k K, v V) bool { return pred(v) })
}
