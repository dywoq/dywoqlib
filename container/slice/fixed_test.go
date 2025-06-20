package slice

import (
	"errors"
	"testing"

	"github.com/dywoq/dywoqlib/iterator"
	"github.com/dywoq/dywoqlib/sliceutil"
)

func TestFixed_Find(t *testing.T) {
	type testCase struct {
		name      string
		size      int
		data      []int
		find      int
		wantValue int
		wantErr   error
		setupErr  error
	}

	var sentinelErr = errors.New("some error")

	tests := []testCase{
		{
			name:      "element present within fixed size",
			size:      4,
			data:      []int{1, 2, 3, 4},
			find:      3,
			wantValue: 3,
			wantErr:   nil,
		},
		{
			name:      "element not present",
			size:      4,
			data:      []int{1, 2, 3, 4},
			find:      5,
			wantValue: 0,
			wantErr:   sliceutil.ErrNotFound,
		},
		{
			name:      "over fixed size",
			size:      2,
			data:      []int{1, 2, 3},
			find:      2,
			wantValue: 0,
			wantErr:   ErrOverFixedSize,
		},
		{
			name:      "previous error exists",
			size:      4,
			data:      []int{1, 2, 3},
			find:      2,
			wantValue: 0,
			wantErr:   sentinelErr,
			setupErr:  sentinelErr,
		},
		{
			name:      "empty slice",
			size:      4,
			data:      []int{},
			find:      1,
			wantValue: 0,
			wantErr:   iterator.ErrInvalidPosition,
		},
		{
			name:      "negative fixed size",
			size:      -1,
			data:      []int{1, 2},
			find:      1,
			wantValue: 0,
			wantErr:   ErrOverFixedSize,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := &Fixed[int]{s: tc.data, err: tc.setupErr, fixedSize: tc.size}
			got := f.Find(tc.find)

			if got != tc.wantValue {
				t.Errorf("Find(%v) = %v, want %v", tc.find, got, tc.wantValue)
			}
			if tc.wantErr == nil {
				if f.Err() != nil {
					t.Errorf("Err() = %v, want nil", f.Err())
				}
			} else {
				if f.Err() == nil || !errors.Is(f.Err(), tc.wantErr) {
					t.Errorf("Err() = %v, want %v", f.Err(), tc.wantErr)
				}
			}
		})
	}
}
