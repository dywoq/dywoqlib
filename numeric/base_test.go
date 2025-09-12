package numeric

import (
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	"math"
	go_testing "testing"

	"github.com/dywoq/dywoqlib/numeric/base"
)

func TestBaseGet(t *go_testing.T) {
	num := Int(32)
	got := num.Get()
	want := 32
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBaseAdd(t *go_testing.T) {
	num := Int(32)
	got := num.Add(4)
	want := 36
	if got.Get() != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBaseMinus(t *go_testing.T) {
	num := Int(32)
	got := num.Minus(4)
	want := 28
	if got.Get() != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBaseDivide(t *go_testing.T) {
	num := Int(32)
	got := num.Divide(16)
	want := 2
	if got.Get() != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBaseMultiply(t *go_testing.T) {
	num := Int(32)
	got := num.Multiply(2)
	want := 64
	if got.Get() != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBaseLimits(t *go_testing.T) {
	num := UInt8(0)
	gotmin, gotmax := num.Limits()
	wantmin, wantmax := 0, uint(math.MaxUint16)
	if gotmin != uint8(wantmin) {
		t.Errorf("gotmin %v, wantmin %v", gotmin, wantmin)
	}
	if gotmax != uint8(wantmax) {
		t.Errorf("gotmax %v, wantmax %v", gotmax, wantmax)
	}
}

func TestBaseZero(t *go_testing.T) {
	num := UInt8(0)
	got := num.Zero()
	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBaseNegative(t *go_testing.T) {
	num := Int(-1)
	got := num.Negative()
	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBasePositive(t *go_testing.T) {
	num := Int(1)
	got := num.Positive()
	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBaseEqual(t *go_testing.T) {
	numA := Int(-1)
	numB := Int(-1)
	got := numA.Equal(numB.Get())
	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBaseCompare(t *go_testing.T) {
	tests := []struct {
		name       string
		num        Base[int]
		compareNum int
		want       base.Compare
	}{
		{"greater", Int(32), 12, base.CompareGreater},
		{"less", Int(32), 64, base.CompareLess},
		{"greater", Int(32), 32, base.CompareEqual},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.Compare(test.compareNum)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestBaseOdd(t *go_testing.T) {
	num := Int(0)
	got := num.Odd()
	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBaseEven(t *go_testing.T) {
	num := Int(7)
	got := num.Even()
	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBaseAbsolute(t *go_testing.T) {
	num := Int(7)
	got := num.Absolute()
	want := 7
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBaseString(t *go_testing.T) {
	num := Int(7)
	got := num.String()
	want := "7"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBaseAnd(t *go_testing.T) {
	num := Int(5)
	got := num.And(1)
	want := 1
	if got.Get() != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBaseOr(t *go_testing.T) {
	num := Int(5)
	got := num.Or(1)
	want := 5
	if got.Get() != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBaseXor(t *go_testing.T) {
	num := Int(5)
	got := num.Xor(1)
	want := 4
	if got.Get() != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBaseShiftLeft(t *go_testing.T) {
	num := Int(5)
	got := num.ShiftLeft(1)
	want := 10
	if got.Get() != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBaseShiftRight(t *go_testing.T) {
	num := Int(5)
	got := num.ShiftRight(1)
	want := 2
	if got.Get() != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBaseSign(t *go_testing.T) {
	tests := []struct {
		name string
		num  Base[int]
		want base.Sign
	}{
		{"greater", Int(42), base.SignPositive},
		{"less", Int(-32), base.SignNegative},
		{"greater", Int(0), base.SignZero},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.Sign()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestBasePrime(t *go_testing.T) {
	num := Int(0)
	got := num.Prime()
	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func BenchmarkBaseSet(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.Set(0)
	}
}

func BenchmarkBaseAdd(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.Add(0)
	}
}

func BenchmarkBaseMinus(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.Minus(0)
	}
}

func BenchmarkBaseDivide(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.Divide(10)
	}
}

func BenchmarkBaseMultiply(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.Multiply(5)
	}
}

func BenchmarkBaseLimits(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_, _ = num.Limits()
	}
}

func BenchmarkBaseZero(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.Zero()
	}
}

func BenchmarkBaseNegative(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.Negative()
	}
}

func BenchmarkBasePositive(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.Positive()
	}
}

func BenchmarkBaseEqual(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.Equal(32)
	}
}

func BenchmarkBaseCompare(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.Compare(32)
	}
}

func BenchmarkBaseOdd(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.Odd()
	}
}

func BenchmarkBaseEven(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.Even()
	}
}

func BenchmarkBaseAbsolute(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.Absolute()
	}
}

func BenchmarkBaseString(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.String()
	}
}

func BenchmarkBaseAnd(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.And(2)
	}
}

func BenchmarkBaseOr(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.Or(2)
	}
}

func BenchmarkBaseXor(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.Xor(2)
	}
}

func BenchmarkBaseShiftLeft(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.ShiftLeft(2)
	}
}

func BenchmarkBaseShiftRight(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.ShiftRight(2)
	}
}

func BenchmarkBaseSign(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.Sign()
	}
}

func BenchmarkBasePrime(b *go_testing.B) {
	internal_testing.SetBase().Benchmark(b)
	num := Int(32)
	for b.Loop() {
		_ = num.Prime()
	}
}
