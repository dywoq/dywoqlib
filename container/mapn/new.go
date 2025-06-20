package mapn

func New[K comparable, V comparable](data map[K]V) Map[K, V] {
	return Map[K, V]{data, nil}
}
