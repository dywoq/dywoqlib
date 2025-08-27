package mapnutil

// Exists reports whether reqkey exists in map m.
// If map is empty, it returns false.
func Exists[K, V comparable](m map[K]V, reqkey K) bool {
	if len(m) == 0 {
		return false
	}
	for key := range m {
		if key == reqkey {
			return true
		}
	}
	return false
}
