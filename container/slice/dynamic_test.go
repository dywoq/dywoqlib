package slice

import (
	"errors"
	"testing"

	"github.com/dywoq/dywoqlib/iterator"
	"github.com/dywoq/dywoqlib/sliceutil"
)

func TestDynamic_Find(t *testing.T) {
	type testCase struct {
		name      string
		data      []int
		find      int
		wantValue int
		wantErr   error
		setupErr  error
	}

	var sentinelErr = errors.New("some error")

	tests := []testCase{
		{
			name:      "element present",
			data:      []int{1, 2, 3, 4},
			find:      3,
			wantValue: 3,
			wantErr:   nil,
		},
		{
			name:      "element not present",
			data:      []int{1, 2, 3, 4},
			find:      5,
			wantValue: 0,
			wantErr:   sliceutil.ErrNotFound,
		},
		{
			name:      "empty slice",
			data:      []int{},
			find:      1,
			wantValue: 0,
			wantErr:   iterator.ErrInvalidPosition,
		},
		{
			name:      "previous error exists",
			data:      []int{1, 2, 3},
			find:      2,
			wantValue: 0,
			wantErr:   sentinelErr,
			setupErr:  sentinelErr,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			d := &Dynamic[int]{s: tc.data, err: tc.setupErr}
			got := d.Find(tc.find)

			if got != tc.wantValue {
				t.Errorf("Find(%v) = %v, want %v", tc.find, got, tc.wantValue)
			}
			if tc.wantErr == nil {
				if d.Err() != nil {
					t.Errorf("Err() = %v, want nil", d.Err())
				}
			} else {
				if d.Err() == nil || !errors.Is(d.Err(), tc.wantErr) {
					t.Errorf("Err() = %v, want %v", d.Err(), tc.wantErr)
				}
			}
		})
	}
}
