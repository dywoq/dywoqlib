package mapn

// New creates a new Map instance.
// It initializes the map with the provided data and no initial error.
func New[K comparable, V comparable](data map[K]V) *Map[K, V] {
	return &Map[K, V]{data, nil}
}
