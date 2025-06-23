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

package mapn

import (
	"errors"
	"testing"
)

func TestMap_At(t *testing.T) {
	type testCase struct {
		name      string
		data      map[int]string
		key       int
		wantValue string
		wantErr   error
		setupErr  error
	}

	tests := []testCase{
		{
			name:      "existing key",
			data:      map[int]string{1: "one", 2: "two"},
			key:       2,
			wantValue: "two",
			wantErr:   nil,
		},
		{
			name:      "missing key",
			data:      map[int]string{1: "one"},
			key:       3,
			wantValue: "",
			wantErr:   ErrNotFound,
		},
		{
			name:      "empty map",
			data:      map[int]string{},
			key:       1,
			wantValue: "",
			wantErr:   ErrEmpty,
		},
		{
			name:      "previous error exists",
			data:      map[int]string{1: "one"},
			key:       1,
			wantValue: "",
			wantErr:   errors.New("some error"),
			setupErr:  errors.New("some error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &Map[int, string]{data: tc.data, err: tc.setupErr}
			got := m.At(tc.key)

			if got != tc.wantValue {
				t.Errorf("At(%v) = %v, want %v", tc.key, got, tc.wantValue)
			}
			if tc.wantErr == nil {
				if m.Err() != nil {
					t.Errorf("Err() = %v, want nil", m.Err())
				}
			} else {
				if m.Err() == nil || m.Err().Error() != tc.wantErr.Error() {
					t.Errorf("Err() = %v, want %v", m.Err(), tc.wantErr)
				}
			}
		})
	}
}