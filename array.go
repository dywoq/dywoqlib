package dywoqlib

import (
	"errors"
	"fmt"
	"slices"
)

// Array is a generic data structure that represents a resizable array.
// It provides methods for common array operations such as getting, setting,
// appending, deleting, and iterating over elements.
// The underlying data is stored in a dynamically sized slice.
type Array[T any] struct {
	size int // The initial size of the array (currently unused and might be removed).
	data []T // The underlying slice that holds the array elements.
}

// NewArray creates and returns a new Array with the specified initial size
// and underlying data. It panics if the length of the provided data slice
// does not match the specified size.
//
// Parameters:
//   - size int: The initial size of the array. Note: This parameter is currently
//     checked against the length of the data slice and might be redundant.
//   - data []T: The initial data to populate the array. The length of this
//     slice must match the 'size' parameter.
//
// Returns:
//   - *Array[T]: A pointer to the newly created Array.
//
// Panics:
//   - If the length of the 'data' slice is not equal to the 'size' parameter.
//
// Example:
//
//	arr := NewArray[int](3, []int{1, 2, 3})
//	fmt.Println(arr) // Output: &[3 [1 2 3]]
func NewArray[T any](size int, data []T) *Array[T] {
	if len(data) != size {
		panic("size of a slice doesn't match initial size")
	}
	return &Array[T]{size, data}
}

// Size returns the current number of elements in the Array.
//
// Returns:
//   - int: The number of elements in the Array.
//
// Example:
//
//	arr := NewArray[int](3, []int{1, 2, 3})
//	fmt.Println(arr.Size()) // Output: 3
func (array *Array[T]) Size() int {
	return len(array.data)
}

// Capacity returns the capacity of the underlying slice of the Array.
// This is the maximum number of elements the Array can hold before
// a new allocation might be necessary.
//
// Returns:
//   - int: The capacity of the underlying slice.
//
// Example:
//
//	arr := NewArray[int](3, make([]int, 3, 5))
//	fmt.Println(arr.Capacity()) // Output: 5
func (array *Array[T]) Capacity() int {
	return cap(array.data)
}

// IsEmpty checks if the Array contains no elements.
//
// Returns:
//   - bool: True if the Array is empty, false otherwise.
//
// Example:
//
//	arr1 := NewArray[int](0, []int{})
//	fmt.Println(arr1.IsEmpty()) // Output: true
//	arr2 := NewArray[int](1, []int{1})
//	fmt.Println(arr2.IsEmpty()) // Output: false
func (array *Array[T]) IsEmpty() bool {
	return len(array.data) == 0
}

// Get retrieves the element at the specified index in the Array.
//
// Parameters:
//   - index int: The index of the element to retrieve. Must be within
//     the bounds of the Array (0 <= index < Size()).
//
// Returns:
//   - T: The element at the specified index.
//   - error: An error if the index is out of range.
//
// Example:
//
//	arr := NewArray[int](3, []int{10, 20, 30})
//	val, err := arr.Get(1)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(val) // Output: 20
//	}
//	_, err = arr.Get(5)
//	if err != nil {
//		fmt.Println(err) // Output: index out of range
//	}
func (array *Array[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= len(array.data) {
		return zero, errors.New("index out of range")
	}
	return array.data[index], nil
}

// Set updates the element at the specified index in the Array with the given value.
//
// Parameters:
//   - index int: The index of the element to update. Must be within
//     the bounds of the Array (0 <= index < Size()).
//   - value T: The new value to set at the specified index.
//
// Returns:
//   - error: An error if the index is out of range.
//
// Example:
//
//	arr := NewArray[int](3, []int{1, 2, 3})
//	err := arr.Set(0, 10)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(arr) // Output: &[3 [10 2 3]]
//	}
//	err = arr.Set(5, 10)
//	if err != nil {
//		fmt.Println(err) // Output: index out of range
//	}
func (array *Array[T]) Set(index int, value T) error {
	if index < 0 || index >= len(array.data) {
		return errors.New("index out of range")
	}
	array.data[index] = value
	return nil
}

// PushBack appends a new element to the end of the Array.
// It might reallocate the underlying slice if the capacity is not sufficient.
// Note: It currently panics if the length of the underlying slice does not
// match the initial size, which might be an unintended behavior.
//
// Parameters:
//   - value T: The element to append to the Array.
//
// Panics:
//   - If the length of the underlying 'data' slice is not equal to the 'size' field.
//
// Example:
//
//	arr := NewArray[int](3, []int{1, 2, 3})
//	arr.PushBack(4)
//	fmt.Println(arr) // Output: &[3 [1 2 3 4]] (Note: size remains 3)
func (array *Array[T]) PushBack(value T) {
	if len(array.data) != array.size {
		panic("size of a slice doesn't match initial size")
	}
	array.data = append(array.data, value)
}

// PopBack removes and returns the last element of the Array.
//
// Returns:
//   - T: The last element of the Array.
//   - error: An error if the Array is empty.
//
// Example:
//
//	arr := NewArray[int](3, []int{1, 2, 3})
//	val, err := arr.PopBack()
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println("Popped:", val)   // Output: Popped: 3
//		fmt.Println("Array:", arr) // Output: Array: &[3 [1 2]]
//	}
//	arrEmpty := NewArray[int](0, []int{})
//	_, err = arrEmpty.PopBack()
//	if err != nil {
//		fmt.Println(err) // Output: array ies empty
//	}
func (array *Array[T]) PopBack() (T, error) {
	var zero T
	if len(array.data) == 0 {
		return zero, errors.New("array is empty")
	}
	last := array.data[len(array.data)-1]
	array.data = array.data[:len(array.data)-1]
	return last, nil
}

// Clear removes all elements from the Array, making it empty.
// The underlying capacity of the slice remains unchanged.
//
// Example:
//
//	arr := NewArray[int](3, []int{1, 2, 3})
//	arr.Clear()
//	fmt.Println(arr) // Output: &[3 []]
func (array *Array[T]) Clear() {
	array.data = array.data[:0]
}

// Insert inserts a new element at the specified index in the Array.
// Elements at and after the index are shifted to make space for the new element.
// It might reallocate the underlying slice if the capacity is not sufficient.
//
// Parameters:
//   - index int: The index at which to insert the new element. Must be within
//     the bounds of the Array (0 <= index <= Size()).
//   - value T: The element to insert.
//
// Returns:
//   - error: An error if the index is out of range.
//
// Example:
//
//	arr := NewArray[int](3, []int{1, 3, 4})
//	err := arr.Insert(1, 2)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(arr) // Output: &[3 [1 2 3 4]]
//	}
//	err = arr.Insert(5, 5)
//	if err != nil {
//		fmt.Println(err) // Output: index out of range
//	}
func (array *Array[T]) Insert(index int, value T) error {
	if index < 0 || index > len(array.data) {
		return errors.New("index out of range")
	}
	array.data = append(array.data, *new(T))
	copy(array.data[index+1:], array.data[index:])
	array.data[index] = value
	return nil
}

// Erase removes the element at the specified index from the Array.
// Elements after the index are shifted to fill the gap.
//
// Parameters:
//   - index int: The index of the element to remove. Must be within
//     the bounds of the Array (0 <= index < Size()).
//
// Returns:
//   - error: An error if the index is out of range.
//
// Example:
//
//	arr := NewArray[int](4, []int{1, 2, 3, 4})
//	err := arr.Erase(1)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(arr) // Output: &[4 [1 3 4]]
//	}
//	err = arr.Erase(5)
//	if err != nil {
//		fmt.Println(err) // Output: index out of range
//	}
func (array *Array[T]) Erase(index int) error {
	if index < 0 || index >= len(array.data) {
		return errors.New("index out of range")
	}
	array.data = slices.Delete(array.data, index, index+1)
	return nil
}

// Reserve increases the capacity of the underlying slice to at least the
// specified capacity. If the provided capacity is less than or equal to the
// current capacity, no action is taken.
//
// Parameters:
//   - capacity int: The desired minimum capacity of the Array.
//
// Example:
//
//	arr := NewArray[int](3, []int{1, 2, 3})
//	fmt.Println("Capacity before reserve:", arr.Capacity()) // Output: Capacity before reserve: 3
//	arr.Reserve(10)
//	fmt.Println("Capacity after reserve:", arr.Capacity())  // Output: Capacity after reserve: 10
//	arr.Reserve(5)
//	fmt.Println("Capacity after second reserve:", arr.Capacity()) // Output: Capacity after second reserve: 10
func (array *Array[T]) Reserve(capacity int) {
	if capacity > cap(array.data) {
		newData := make([]T, len(array.data), capacity)
		copy(newData, array.data)
		array.data = newData
	}
}

// Back returns the last element of the Array.
//
// Returns:
//   - T: The last element of the Array.
//   - error: An error if the Array is empty.
//
// Example:
//
//	arr := NewArray[int](3, []int{1, 2, 3})
//	val, err := arr.Back()
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(val) // Output: 3
//	}
//	arrEmpty := NewArray[int](0, []int{})
//	_, err = arrEmpty.Back()
//	if err != nil {
//		fmt.Println(err) // Output: array is empty
//	}
func (array *Array[T]) Back() (T, error) {
	var zero T
	if array.IsEmpty() {
		return zero, errors.New("array is empty")
	}
	return array.data[len(array.data)-1], nil
}

// String returns a string representation of the Array.
//
// Returns:
//   - string: A string representation of the Array, typically the
//     string representation of its underlying slice.
//
// Example:
//
//	arr := NewArray[int](3, []int{1, 2, 3})
//	fmt.Println(arr.String()) // Output: [1 2 3]
func (array *Array[T]) String() string {
	return fmt.Sprintf("%v", array.data)
}

// Range provides a way to iterate over all elements in the Array.
// It returns a read-only channel that yields each element in the Array
// in order, from index 0 to the last index. The channel is closed
// after all elements have been sent. This is similar to range-based
// for loops in C++ or Go's range keyword for slices.
//
// Returns:
//   - <-chan T: A read-only channel yielding the Array's elements.
//
// Example:
//
//	arr := NewArray[int](3, []int{10, 20, 30})
//	for val := range arr.Range() {
//		fmt.Println(val) // Prints 10, 20, 30
//	}
//
// Time Complexity:
//   - O(n), where n is the number of elements in the Array, as each element
//     is sent to the channel exactly once.
//
// Note:
//   - The channel is closed automatically after all elements are yielded.
//   - If the Array is empty, the channel is closed immediately.
//   - Concurrent modification of the Array during iteration is not safe and
//     may lead to undefined behavior.
func (array *Array[T]) Range() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, elem := range array.data {
			ch <- elem
		}
	}()
	return ch
}
