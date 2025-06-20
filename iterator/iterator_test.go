package iterator

import (
	"testing"
)

type testSlice struct {
	data []int
}

func (t *testSlice) Begin() *Iterator[int] {
	return New(0, t.data)
}

func (t *testSlice) End() *Iterator[int] {
	return New(len(t.data)-1, t.data)
}

func TestIterable_BeginEnd(t *testing.T) {
	ts := &testSlice{data: []int{10, 20, 30}}
	var it Iterable[int] = ts

	begin := it.Begin()
	if begin == nil || begin.Value() != 10 {
		t.Errorf("Begin() = %v, want value 10", begin)
	}

	end := it.End()
	if end == nil || end.Value() != 30 {
		t.Errorf("End() = %v, want value 30", end)
	}
}
