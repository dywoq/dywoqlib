package myvector

import (
	"errors"
	"fmt"
)

// Vector is a dynamic array that can hold elements of any type.
// It is similar to std::vector in C++.
//
// The zero value for Vector is an empty vector.
type Vector[T any] struct {
	data []T // The underlying slice used to store the elements.
}

// NewVector creates a new Vector with a specified initial capacity.
//
// Parameters:
//   - capacity: The initial capacity of the vector. If capacity is less than or equal to 0,
//     the vector will be initialized with an empty slice (length and capacity 0).
//
// Returns:
//   - *Vector[T]: A pointer to the newly created Vector.
//
// Example:
//
//	v := NewVector[int](10) // Creates a new vector of integers with initial capacity 10.
//	v2 := NewVector[string](0) // Creates a new vector of strings with initial capacity 0.
func NewVector[T any](capacity int) *Vector[T] {
	if capacity <= 0 {
		return &Vector[T]{data: make([]T, 0)} // Initializes with zero length and zero capacity.
	}
	return &Vector[T]{data: make([]T, 0, capacity)} // Initializes with zero length and specified capacity.
}

// Size returns the number of elements currently stored in the vector.
// This is equivalent to the size() method in C++.
//
// Returns:
//   - int: The number of elements in the vector.
//
// Example:
//
//	v := NewVector[int](10)
//	v.PushBack(1)
//	v.PushBack(2)
//	size := v.Size() // size will be 2
func (v *Vector[T]) Size() int {
	return len(v.data)
}

// Capacity returns the total number of elements the vector can store
// without reallocating memory.  This is equivalent to the capacity()
// method in C++.
//
// Returns:
//   - int: The capacity of the vector.
//
// Example:
//
//	v := NewVector[int](10)
//	capacity := v.Capacity() // capacity will be 10
//	v2 := NewVector[int](0)
//	capacity2 := v2.Capacity() // capacity2 will be 0
func (v *Vector[T]) Capacity() int {
	return cap(v.data)
}

// IsEmpty checks if the vector is empty (contains no elements).
// This is equivalent to the empty() method in C++.
//
// Returns:
//   - bool: true if the vector is empty, false otherwise.
//
// Example:
//
//	v := NewVector[int](10)
//	isEmpty := v.IsEmpty() // isEmpty will be true
//	v.PushBack(1)
//	isEmpty = v.IsEmpty() // isEmpty will be false
func (v *Vector[T]) IsEmpty() bool {
	return len(v.data) == 0
}

// Get returns the element at the specified index.
// This is similar to the at() method in C++, but returns an error
// instead of panicking.
//
// Parameters:
//   - index: The index of the element to retrieve.
//
// Returns:
//   - T: The element at the specified index.  Will be the zero value
//     of type T if the index is out of range.
//   - error: An error if the index is out of range, nil otherwise.
//
// Example:
//
//	v := NewVector[int](10)
//	v.PushBack(100)
//	val, err := v.Get(0) // val will be 100, err will be nil
//	val, err = v.Get(1) // val will be 0, err will be "index out of range"
func (v *Vector[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= len(v.data) {
		return zero, errors.New("index out of range")
	}
	return v.data[index], nil
}

// Set sets the element at the specified index to the given value.
// This is similar to using the [] operator in C++ to modify an element.
//
// Parameters:
//   - index: The index of the element to set.
//   - value: The new value for the element.
//
// Returns:
//   - error: An error if the index is out of range, nil otherwise.
//
// Example:
//
//	v := NewVector[int](10)
//	v.PushBack(100)
//	err := v.Set(0, 200) // v[0] is now 200, err is nil
//	err = v.Set(1, 300) // err will be "index out of range"
func (v *Vector[T]) Set(index int, value T) error {
	if index < 0 || index >= len(v.data) {
		return errors.New("index out of range")
	}
	v.data[index] = value
	return nil
}

// PushBack adds a new element to the end of the vector.
// This is equivalent to the push_back() method in C++.
//
// Parameters:
//   - value: The value to add to the vector.
//
// Example:
//
//	v := NewVector[int](10)
//	v.PushBack(1) // v is now [1]
//	v.PushBack(2) // v is now [1 2]
//
// Time Complexity:
//
//	Amortized O(1).  PushBack usually takes constant time.  However, if the
//	vector's capacity is full, it needs to reallocate memory, which takes
//	O(n) time, where n is the current size of the vector.  Because
//	reallocations happen infrequently (capacity grows exponentially), the
//	average time complexity over many PushBack operations is O(1).
func (v *Vector[T]) PushBack(value T) {
	v.data = append(v.data, value)
}

// PopBack removes the last element from the vector.
// This is equivalent to the pop_back() method in C++.
//
// Returns:
//   - T: The value of the last element before it was removed.  Will be the
//     zero value of type T if the vector is empty.
//   - error: An error if the vector is empty, nil otherwise.
//
// Example:
//
//	v := NewVector[int](10)
//	v.PushBack(1)
//	v.PushBack(2)
//	val, err := v.PopBack() // val is 2, v is now [1], err is nil
//	val, err = v.PopBack() // val is 1, v is now [], err is nil
//	val, err = v.PopBack() // val is 0, err is "vector is empty"
func (v *Vector[T]) PopBack() (T, error) {
	var zero T
	if len(v.data) == 0 {
		return zero, errors.New("vector is empty")
	}
	last := v.data[len(v.data)-1]
	v.data = v.data[:len(v.data)-1]
	return last, nil
}

// Clear removes all elements from the vector, making it empty.
// This is equivalent to the clear() method in C++.  However, unlike the C++
// version, this does not change the capacity of the vector.
//
// Example:
//
//	v := NewVector[int](10)
//	v.PushBack(1)
//	v.PushBack(2)
//	v.Clear() // v is now [], but its capacity is still 10
func (v *Vector[T]) Clear() {
	v.data = v.data[:0]
}

// Insert inserts a new element at the specified index.
// This is similar to the insert() method in C++.
//
// Parameters:
//   - index: The index at which to insert the new element.
//   - value: The value of the new element.
//
// Returns:
//   - error: An error if the index is out of range, nil otherwise.
//
// Example:
//
//	v := NewVector[int](10)
//	v.PushBack(1)     // v is [1]
//	v.PushBack(3)     // v is [1 3]
//	err := v.Insert(1, 2) // v is [1 2 3], err is nil
//	err = v.Insert(0, 0) // v is [0 1 2 3], err is nil
//	err = v.Insert(4, 4) // v is [0 1 2 3 4], err is nil
//	err = v.Insert(5, 5) // err is "index out of range"
//
// Time Complexity:
//
//	O(n), where n is the number of elements after the insertion point.
//	Inserting an element requires shifting all subsequent elements to make
//	space, which takes linear time.
func (v *Vector[T]) Insert(index int, value T) error {
	if index < 0 || index > len(v.data) {
		return errors.New("index out of range")
	}
	v.data = append(v.data, *new(T))
	copy(v.data[index+1:], v.data[index:])
	v.data[index] = value
	return nil
}

// Erase removes the element at the specified index.
// This is similar to the erase() method in C++.
//
// Parameters:
//   - index: The index of the element to remove.
//
// Returns:
//   - error: An error if the index is out of range, nil otherwise.
//
// Example:
//
//	v := NewVector[int](10)
//	for i := 0; i < 5; i++ {
//		v.PushBack(i) // v is [0 1 2 3 4]
//	}
//	err := v.Erase(2) // v is [0 1 3 4], err is nil
//	err = v.Erase(0) // v is [1 3 4], err is nil
//	err = v.Erase(2) // v is [1 3], err is nil
//	err = v.Erase(2) // err is "index out of range"
//
// Time Complexity:
//
//	O(n), where n is the number of elements after the deletion point.
//	Erasing an element requires shifting all subsequent elements to close
//	the gap, which takes linear time.
func (v *Vector[T]) Erase(index int) error {
	if index < 0 || index >= len(v.data) {
		return errors.New("index out of range")
	}
	v.data = append(v.data[:index], v.data[index+1:]...)
	return nil
}

// Reserve increases the capacity of the vector to at least the specified capacity.
// If the new capacity is less than the current capacity, this method does nothing.
// This is similar to the reserve() method in C++.
//
// Parameters:
//   - capacity: The new minimum capacity of the vector.
//
// Example:
//
//	v := NewVector[int](5)      // capacity is 5
//	v.Reserve(10)             // capacity is now at least 10
//	v.Reserve(3)              // capacity remains unchanged (at least 10)
//	v2 := NewVector[int](0)     // capacity is 0
//	v2.Reserve(100)           // capacity is now 100
func (v *Vector[T]) Reserve(capacity int) {
	if capacity > cap(v.data) {
		newData := make([]T, len(v.data), capacity)
		copy(newData, v.data)
		v.data = newData
	}
}

// Resize changes the number of elements in the vector.
//
// If the new size is greater than the current size, new elements are
// added to the end of the vector and initialized with the zero value
// of type T.
// If the new size is less than the current size, the vector is truncated,
// and the extra elements are discarded.
// This is similar to the resize() method in C++.
//
// Parameters:
//   - size: The new size of the vector.
//
// Panics:
//   - If size is negative.
//
// Example:
//
//	v := NewVector[int](5)
//	v.PushBack(1) // v is [1]
//	v.PushBack(2) // v is [1 2]
//	v.Resize(5)   // v is [1 2 0 0 0] (capacity is at least 5)
//	v.Resize(1)   // v is [1]
//	v.Resize(0)   // v is []
//
// Go's behavior with resize is different from C++.  If you resize to a larger
// size, Go will automatically initialize the new elements to the zero value
// for the type.  C++ may leave the values uninitialized.
func (v *Vector[T]) Resize(size int) {
	if size < 0 {
		panic("vector: negative resize argument")
	}
	if size > len(v.data) {
		newData := make([]T, size)
		copy(newData, v.data)
		v.data = newData
	} else if size < len(v.data) {
		v.data = v.data[:size]
	}
}

// Front returns the first element in the vector.
// This is similar to the front() method in C++.
//
// Returns:
//   - T: The value of the first element. Will be the zero value of type T
//     if the vector is empty.
//   - error: An error if the vector is empty, nil otherwise.
//
// Example:
//
//	v := NewVector[int](10)
//	v.PushBack(10)
//	val, err := v.Front() // val is 10, err is nil
//	v.Clear()
//	val, err = v.Front() // val is 0, err is "vector is empty"
func (v *Vector[T]) Front() (T, error) {
	var zero T
	if v.IsEmpty() {
		return zero, errors.New("vector is empty")
	}
	return v.data[0], nil
}

// Back returns the last element in the vector.
// This is similar to the back() method in C++.
//
// Returns:
//   - T: The value of the last element. Will be the zero value of type T
//     if the vector is empty.
//   - error: An error if the vector is empty, nil otherwise.
//
// Example:
//
//	v := NewVector[int](10)
//	v.PushBack(10)
//	v.PushBack(20)
//	val, err := v.Back() // val is 20, err is nil
//	v.Clear()
//	val, err = v.Back() // val is 0, err is "vector is empty"
func (v *Vector[T]) Back() (T, error) {
	var zero T
	if v.IsEmpty() {
		return zero, errors.New("vector is empty")
	}
	return v.data[len(v.data)-1], nil
}

// String returns a string representation of the vector.
// This is similar to how vectors are often printed in C++ using operator<<.
//
// Returns:
//   - string: A string representation of the vector, which is the same as
//     the string representation of the underlying slice.
//
// Example:
//
//	v := NewVector[int](10)
//	v.PushBack(1)
//	v.PushBack(2)
//	str := v.String() // str is "[1 2]"
func (v *Vector[T]) String() string {
	return fmt.Sprintf("%v", v.data)
}
