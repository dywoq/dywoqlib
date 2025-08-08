package asm

import "unsafe"

// String simplifies creating a Go string from a pointer and a given length in Go assembly.
// This function is a wrapper for a Go assembly implementation that constructs
// a string header from the provided pointer and length without allocating new memory.
// The caller must ensure that the backing memory for the pointer remains valid
// for the entire lifetime of the returned string to prevent runtime errors.
//
//go:inline
func String(ptr unsafe.Pointer, len int) string
