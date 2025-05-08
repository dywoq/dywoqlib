package dywoqlib

import "errors"

// Map is a generic wrapper around Go's built-in map type, providing methods to manage key-value pairs
// with error handling and iteration support. The KeyType must be comparable to support map key operations,
// while ValueType can be any type.
type Map[KeyType comparable, ValueType any] struct {
	data map[KeyType]ValueType
}

// IsEmpty checks if the map contains no key-value pairs.
//
// Returns:
//   - bool: true if the map is empty (i.e., has zero key-value pairs), false otherwise.
//
// Time Complexity:
//   - O(1), as it only checks the length of the underlying map.
func (m *Map[KeyType, ValueType]) IsEmpty() bool {
	return len(m.data) == 0
}

// NewMap creates a new Map instance with an optional initial map. If the provided map is empty or nil,
// it initializes an empty map.
//
// Parameters:
//   - data map[KeyType]ValueType: The initial map to use, or nil/empty to create a new empty map.
//
// Returns:
//   - *Map[KeyType, ValueType]: A pointer to the newly created Map instance.
//
// Time Complexity:
//   - O(1), as it only involves map initialization or assignment.
//
// Note:
//   - If the input map is empty, a new empty map is created to ensure the Map's data field is non-nil.
//   - The returned Map is ready for use with methods like Add, Remove, etc.
func NewMap[KeyType comparable, ValueType any](data map[KeyType]ValueType) *Map[KeyType, ValueType] {
	if len(data) == 0 {
		return &Map[KeyType, ValueType]{data: map[KeyType]ValueType{}}
	}
	return &Map[KeyType, ValueType]{data: data}
}

// Add inserts a key-value pair into the map. If the map is empty, it returns an error.
//
// Parameters:
//   - key KeyType: The key to associate with the value.
//   - value ValueType: The value to store.
//
// Returns:
//   - error: An error if the map is empty, nil otherwise.
//
// Time Complexity:
//   - O(1) on average for map insertion.
//
// Note:
//   - The empty map check may be counterintuitive, as maps can typically be populated from an empty state.
//     Consider initializing the map in a constructor or removing this check for standard map behavior.
func (m *Map[KeyType, ValueType]) Add(key KeyType, value ValueType) (e error) {
	if m.IsEmpty() {
		e = errors.New("map is empty")
		return
	}
	m.data[key] = value
	return
}

// Remove deletes a key-value pair from the map by key. It returns an error if the map is empty or the key is not found.
//
// Parameters:
//   - key KeyType: The key to remove.
//
// Returns:
//   - error: An error if the map is empty or the key is not found, nil otherwise.
//
// Time Complexity:
//   - O(1) on average for map deletion and lookup.
//
// Example:
//   m := &Map[string, int]{data: map[string]int{"a": 1}}
//   err := m.Remove("a") // Returns nil
//   err = m.Remove("b")  // Returns "key wasn't found in the map"
func (m *Map[KeyType, ValueType]) Remove(key KeyType) (e error) {
	if m.IsEmpty() {
		e = errors.New("map is empty")
		return
	}

	_, isFound := m.data[key]
	if !isFound {
		e = errors.New("key wasn't found in the map")
		return
	}
	delete(m.data, key)
	return
}

// IsNil checks if the underlying map is nil.
//
// Returns:
//   - bool: true if the map's data is nil, false otherwise.
//
// Time Complexity:
//   - O(1), as it only checks the nil status of the map.
func (m *Map[KeyType, ValueType]) IsNil() bool {
	return m.data == nil
}

// GetKeyValue retrieves the value associated with a given key. It returns an error if the key is not found.
//
// Parameters:
//   - key KeyType: The key to look up.
//
// Returns:
//   - value ValueType: The value associated with the key, or the zero value if not found.
//   - error: An error if the key is not found, nil otherwise.
//
// Time Complexity:
//   - O(1) on average for map lookup.
//
// Example:
//   m := &Map[string, int]{data: map[string]int{"a": 1}}
//   val, err := m.GetKeyValue("a") // Returns 1, nil
//   val, err = m.GetKeyValue("b")  // Returns 0, "key wasn't found in the map"
func (m *Map[KeyType, ValueType]) GetKeyValue(key KeyType) (value ValueType, e error) {
	value, ok := m.data[key]
	if !ok {
		e = errors.New("key wasn't found in the map")
	}
	return
}

// RemoveAll deletes multiple key-value pairs from the map by their keys. It silently ignores keys that do not exist.
//
// Parameters:
//   - keys []KeyType: A slice of keys to remove.
//
// Time Complexity:
//   - O(n) where n is the number of keys in the input slice, as each deletion is O(1) on average.
//
// Note:
//   - If the map is nil, the method returns immediately without modifying anything.
func (m *Map[KeyType, ValueType]) RemoveAll(keys []KeyType) {
	if m.data == nil {
		return
	}
	for _, key := range keys {
		delete(m.data, key)
	}
}

// Clear removes all key-value pairs from the map, leaving it empty but not nil.
//
// Time Complexity:
//   - O(1), as Go's clear function is optimized for maps.
//
// Note:
//   - If the map is nil, the method does nothing.
//   - The underlying map remains allocated and can be reused.
func (m *Map[KeyType, ValueType]) Clear() {
	if m.data != nil {
		clear(m.data)
	}
}

// ContainsKey checks if a given key exists in the map.
//
// Parameters:
//   - key KeyType: The key to check.
//
// Returns:
//   - bool: true if the key exists in the map, false otherwise.
//
// Time Complexity:
//   - O(1) on average for map lookup.
//
// Note:
//   - Returns false if the map is nil.
func (m *Map[KeyType, ValueType]) ContainsKey(key KeyType) bool {
	if m.data == nil {
		return false
	}
	_, ok := m.data[key]
	return ok
}

// Size returns the number of key-value pairs in the map.
//
// Returns:
//   - int: The number of key-value pairs in the map, or 0 if the map is nil.
//
// Time Complexity:
//   - O(1), as it only retrieves the length of the underlying map.
func (m *Map[KeyType, ValueType]) Size() int {
	if m.data == nil {
		return 0
	}
	return len(m.data)
}

// Keys returns a slice containing all keys in the map.
//
// Returns:
//   - []KeyType: A slice of all keys in the map, or nil if the map is nil.
//
// Time Complexity:
//   - O(n) where n is the number of keys, as it iterates over the map to collect keys.
//
// Note:
//   - The order of keys in the returned slice is arbitrary, as Go maps do not guarantee iteration order.
func (m *Map[KeyType, ValueType]) Keys() []KeyType {
	if m.data == nil {
		return nil
	}
	keys := make([]KeyType, 0, len(m.data))
	for key := range m.data {
		keys = append(keys, key)
	}
	return keys
}

// Values returns a slice containing all values in the map.
//
// Returns:
//   - []ValueType: A slice of all values in the map, or nil if the map is nil.
//
// Time Complexity:
//   - O(n) where n is the number of values, as it iterates over the map to collect values.
//
// Note:
//   - The order of values in the returned slice is arbitrary, as Go maps do not guarantee iteration order.
func (m *Map[KeyType, ValueType]) Values() []ValueType {
	if m.data == nil {
		return nil
	}
	values := make([]ValueType, 0, len(m.data))
	for _, value := range m.data {
		values = append(values, value)
	}
	return values
}

// Range provides a way to iterate over all key-value pairs in the map.
// It returns a read-only channel that yields each key-value pair in the map
// in an arbitrary order, as maps in Go do not guarantee iteration order.
// The channel is closed after all pairs have been sent. This is similar to
// range-based for loops in C++ or Go's range keyword for maps.
//
// Returns:
//   - <-chan struct{ Key KeyType; Value ValueType }: A read-only channel yielding
//     the map's key-value pairs as anonymous structs.
//
// Time Complexity:
//   - O(n), where n is the number of key-value pairs in the map, as each pair
//     is sent to the channel exactly once.
//
// Example:
//   m := &Map[string, int]{data: make(map[string]int)}
//   m.Add("a", 1)
//   m.Add("b", 2)
//   for pair := range m.Range() {
//       fmt.Printf("Key: %s, Value: %d\n", pair.Key, pair.Value) // Prints key-value pairs
//   }
//
// Note:
//   - The channel is closed automatically after all pairs are yielded.
//   - If the map is empty or nil, the channel is closed immediately.
//   - Concurrent modification of the map during iteration is not safe and
//     may lead to undefined behavior.
//   - The order of iteration is arbitrary, as Go maps do not guarantee a
//     consistent iteration order.
func (m *Map[KeyType, ValueType]) Range() <-chan struct {
	Key   KeyType
	Value ValueType
} {
	ch := make(chan struct {
		Key   KeyType
		Value ValueType
	})
	go func() {
		defer close(ch)
		if m.data == nil {
			return
		}
		for key, value := range m.data {
			ch <- struct {
				Key   KeyType
				Value ValueType
			}{Key: key, Value: value}
		}
	}()
	return ch
}
