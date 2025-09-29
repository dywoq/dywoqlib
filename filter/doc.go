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
