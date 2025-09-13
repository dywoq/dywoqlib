// Copyright 2025 dywoq
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package iterator

import (
	"errors"
	"testing"
)

func TestReverse(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		it := NewReverse([]int{})
		if it.Next() {
			t.Fatal("expected false, got true")
		}
		if it.Position() != -1 {
			t.Fatalf("expected position -1, got %d", it.Position())
		}
		if it.Value() != 0 {
			t.Fatalf("expected value 0, got %d", it.Value())
		}
		if !errors.Is(it.Error(), ErrOutOfBounds) {
			t.Fatalf("expected err %v, got %v", ErrOutOfBounds, it.Error())
		}
		if it.Length() != 0 {
			t.Fatalf("expected length 0, got %d", it.Length())
		}
	})

	t.Run("one element", func(t *testing.T) {
		it := NewReverse([]int{42})
		if !it.Next() {
			t.Fatal("expected true, got false")
		}
		if it.Position() != 0 {
			t.Fatalf("expected position 0, got %d", it.Position())
		}
		if it.Value() != 42 {
			t.Fatalf("expected value 42, got %d", it.Value())
		}
		if it.Error() != nil {
			t.Fatalf("expected nil error, got %v", it.Error())
		}
		if it.Length() != 1 {
			t.Fatalf("expected length 1, got %d", it.Length())
		}

		if it.Next() {
			t.Fatal("expected false, got true")
		}
		if it.Position() != -1 {
			t.Fatalf("expected position -1, got %d", it.Position())
		}
		if it.Value() != 0 {
			t.Fatalf("expected value 0, got %d", it.Value())
		}
		if !errors.Is(it.Error(), ErrOutOfBounds) {
			t.Fatalf("expected err %v, got %v", ErrOutOfBounds, it.Error())
		}
	})

	t.Run("multiple elements", func(t *testing.T) {
		data := []int{1, 2, 3}
		it := NewReverse(data)

		if !it.Next() {
			t.Fatal("expected true, got false")
		}
		if it.Position() != 2 {
			t.Fatalf("expected position 2, got %d", it.Position())
		}
		if it.Value() != 3 {
			t.Fatalf("expected value 3, got %d", it.Value())
		}
		if it.Error() != nil {
			t.Fatalf("expected nil error, got %v", it.Error())
		}
		if it.Length() != 3 {
			t.Fatalf("expected length 3, got %d", it.Length())
		}

		if !it.Next() {
			t.Fatal("expected true, got false")
		}
		if it.Position() != 1 {
			t.Fatalf("expected position 1, got %d", it.Position())
		}
		if it.Value() != 2 {
			t.Fatalf("expected value 2, got %d", it.Value())
		}
		if it.Error() != nil {
			t.Fatalf("expected nil error, got %v", it.Error())
		}

		if !it.Next() {
			t.Fatal("expected true, got false")
		}
		if it.Position() != 0 {
			t.Fatalf("expected position 0, got %d", it.Position())
		}
		if it.Value() != 1 {
			t.Fatalf("expected value 1, got %d", it.Value())
		}
		if it.Error() != nil {
			t.Fatalf("expected nil error, got %v", it.Error())
		}

		if it.Next() {
			t.Fatal("expected false, got true")
		}
		if it.Position() != -1 {
			t.Fatalf("expected position -1, got %d", it.Position())
		}
		if it.Value() != 0 {
			t.Fatalf("expected value 0, got %d", it.Value())
		}
		if !errors.Is(it.Error(), ErrOutOfBounds) {
			t.Fatalf("expected err %v, got %v", ErrOutOfBounds, it.Error())
		}
	})

	t.Run("reset", func(t *testing.T) {
		data := []int{1, 2, 3}
		it := NewReverse(data)

		if !it.Next() {
			t.Fatal("expected true, got false")
		}
		if it.Value() != 3 {
			t.Fatalf("expected value 3, got %d", it.Value())
		}

		it.Reset()
		if it.Position() != 3 {
			t.Fatalf("expected position 3, got %d", it.Position())
		}
		if it.Error() != nil {
			t.Fatalf("expected nil error, got %v", it.Error())
		}

		if !it.Next() {
			t.Fatal("expected true, got false")
		}
		if it.Value() != 3 {
			t.Fatalf("expected value 3, got %d", it.Value())
		}
	})

	t.Run("value ptr", func(t *testing.T) {
		data := []int{1, 2, 3}
		it := NewReverse(data)

		if !it.Next() {
			t.Fatal("expected true, got false")
		}
		ptr := it.ValuePtr()
		if ptr == nil {
			t.Fatal("expected not nil, got nil")
		}
		if *ptr != 3 {
			t.Fatalf("expected value 3, got %d", *ptr)
		}

		*ptr = 42
		if data[2] != 42 {
			t.Fatalf("expected data[2] to be 42, got %d", data[2])
		}
	})

	t.Run("value ptr out of bounds", func(t *testing.T) {
		it := NewReverse([]int{})
		if it.Next() {
			t.Fatal("expected false, got true")
		}
		if it.ValuePtr() != nil {
			t.Fatal("expected nil, got not nil")
		}
		if !errors.Is(it.Error(), ErrOutOfBounds) {
			t.Fatalf("expected err %v, got %v", ErrOutOfBounds, it.Error())
		}
	})
}
