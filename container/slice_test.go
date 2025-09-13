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

package container_test

import (
	"testing"

	"github.com/dywoq/dywoqlib/container"
)

func TestGrowableSlice_Grow(t *testing.T) {
	t.Run("negative capacity should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		s := make(container.GrowableSlice[int], 0)
		s.Grow(-1)
	})

	t.Run("capacity less than i", func(t *testing.T) {
		s := make(container.GrowableSlice[int], 0, 2)
		s.Grow(5)
		if cap(s) != 5 {
			t.Errorf("expected capacity 5, got %d", cap(s))
		}
	})

	t.Run("capacity greater than i", func(t *testing.T) {
		s := make(container.GrowableSlice[int], 0, 10)
		s.Grow(5)
		if cap(s) != 10 {
			t.Errorf("expected capacity 10, got %d", cap(s))
		}
	})
}
