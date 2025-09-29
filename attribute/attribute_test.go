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

package attribute_test

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/dywoq/dywoqlib/attribute"
)

func TestAttribute(t *testing.T) {
	testCases := []struct {
		name      string
		attribute func(func())
		expected  string
	}{
		{"Deprecated", attribute.Deprecated, "attribute.Deprecated: github.com/dywoq/dywoqlib/attribute_test.TestAttribute"},
		{"Experimental", attribute.Experimental, "attribute.Experimental: github.com/dywoq/dywoqlib/attribute_test.TestAttribute"},
		{"Removed", attribute.Removed, "attribute.Removed: github.com/dywoq/dywoqlib/attribute_test.TestAttribute"},
		{"Todo", attribute.Todo, "attribute.Todo: todo in github.com/dywoq/dywoqlib/attribute_test.TestAttribute"},
		{"Unsafe", attribute.Unsafe, "attribute.Unsafe: github.com/dywoq/dywoqlib/attribute_test.TestAttribute"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test with nil event (capture output)
			t.Run("nil event", func(t *testing.T) {
				old := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				log.SetOutput(w)

				tc.attribute(nil)

				w.Close()
				os.Stdout = old
				log.SetOutput(os.Stderr)

				var buf bytes.Buffer
				buf.ReadFrom(r)
				output := strings.TrimSpace(buf.String())

				if !strings.HasPrefix(output, tc.expected) {
					t.Errorf("expected prefix %q, got %q", tc.expected, output)
				}
			})

			// Test with non-nil event
			t.Run("non-nil event", func(t *testing.T) {
				called := false
				event := func() {
					called = true
				}
				tc.attribute(event)
				if !called {
					t.Error("event function was not called")
				}
			})
		})
	}
}
