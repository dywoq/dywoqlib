// filter provides a functionality to filter slices and maps.
//
// # Purpose
//
// The purpose is to reduce rewriting the same logic code.
//
// The developers have to rewrite the same logic of filtering slices and maps every time,
// which is redundant and time-wasting. filter package solves this problem.
//
// # Example of using
//
//  package main
//
//  import (
//    "fmt"
//    "github.com/dywoq/dywoqlib/filter"
//  )
//
//  func main() {
//     s := filter.Slice([]int{1, 2, 3, 4, 5, 6}, func(elem int) bool { return elem%2 == 0 })
//     fmt.Println(s)
//  }
//
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

// Map returns a filtered map of keys that satisfy pred.
// Returns an empty map if len(m) is 0.
func MapNot[K comparable, V any](m map[K]V, pred func(K, V) bool) map[K]V {
	return Map(m, func(k K, v V) bool { return !pred(k, v) })
}
