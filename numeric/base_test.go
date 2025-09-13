package numeric

import (
	internal_testing "github.com/dywoq/dywoqlib/internal/testing"
	"math"
	go_testing "testing"

	"github.com/dywoq/dywoqlib/numeric/base"
)

func TestBaseGet(t *go_testing.T) {
	tests := []struct {
		name string
		num  Base[int]
		want int
	}{
		{"positive", Int(32), 32},
		{"negative", Int(-10), -10},
		{"zero", Int(0), 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.Get()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
func TestBaseAdd(t *go_testing.T) {
	tests := []struct {
		name string
		num  Base[int]
		add  int
		want int
	}{
		{"positive", Int(32), 4, 36},
		{"add negative", Int(10), -5, 5},
		{"add zero", Int(5), 0, 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.Add(test.add)
			want := Int(test.want)
			if !got.Equal(want.Get()) {
				t.Errorf("got %v, want %v", got.Get(), want.Get())
			}
		})
	}
}

func TestBaseMinus(t *go_testing.T) {
	tests := []struct {
		name string
		num  Base[int]
		sub  int
		want int
	}{
		{"positive", Int(32), 4, 28},
		{"subtract negative", Int(10), -5, 15},
		{"from zero", Int(0), 5, -5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.Minus(test.sub)
			want := Int(test.want)
			if !got.Equal(want.Get()) {
				t.Errorf("got %v, want %v", got.Get(), want.Get())
			}
		})
	}
}

func TestBaseDivide(t *go_testing.T) {
	tests := []struct {
		name      string
		num       Base[int]
		div       int
		want      int
		wantError bool
	}{
		{"divides evenly", Int(32), 16, 2, false},
		{"odd result", Int(10), 3, 3, false},
		{"negative result", Int(10), -5, -2, false},
		{"divide by zero", Int(10), 0, 10, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.Divide(test.div)
			if got.Error().Nil() && test.wantError {
				t.Error("wanted error, got nil")
			}
			if !got.Error().Nil() && !test.wantError {
				t.Errorf("wanted no error, got %v", got.Error())
			}
			if !test.wantError && !got.Equal(test.want) {
				t.Errorf("got %v, want %v", got.Get(), test.want)
			}
		})
	}
}

func TestBaseMultiply(t *go_testing.T) {
	tests := []struct {
		name string
		num  Base[int]
		mul  int
		want int
	}{
		{"positive", Int(32), 2, 64},
		{"multiply by zero", Int(10), 0, 0},
		{"multiply by one", Int(10), 1, 10},
		{"multiply by negative", Int(5), -2, -10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.Multiply(test.mul)
			want := Int(test.want)
			if !got.Equal(want.Get()) {
				t.Errorf("got %v, want %v", got.Get(), want.Get())
			}
		})
	}
}

func TestBaseLimits(t *go_testing.T) {
	num := UInt8(0)
	gotmin, gotmax := num.Limits()
	wantmin, wantmax := uint8(0), uint8(math.MaxUint8)
	if gotmin != wantmin {
		t.Errorf("gotmin %v, wantmin %v", gotmin, wantmin)
	}
	if gotmax != wantmax {
		t.Errorf("gotmax %v, wantmax %v", gotmax, wantmax)
	}
}

func TestBaseZero(t *go_testing.T) {
	tests := []struct {
		name string
		num  Base[uint8]
		want bool
	}{
		{"is zero", UInt8(0), true},
		{"is not zero", UInt8(5), false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.Zero()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestBaseNegative(t *go_testing.T) {
	tests := []struct {
		name string
		num  Base[int]
		want bool
	}{
		{"is negative", Int(-1), true},
		{"is positive", Int(1), false},
		{"is zero", Int(0), false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.Negative()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestBasePositive(t *go_testing.T) {
	tests := []struct {
		name string
		num  Base[int]
		want bool
	}{
		{"is positive", Int(1), true},
		{"is negative", Int(-1), false},
		{"is zero", Int(0), false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.Positive()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
func TestBaseEqual(t *go_testing.T) {
	tests := []struct {
		name string
		numA Base[int]
		numB Base[int]
		want bool
	}{
		{"equal", Int(-1), Int(-1), true},
		{"not equal", Int(5), Int(10), false},
		{"different signs", Int(-5), Int(5), false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.numA.Equal(test.numB.Get())
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
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
	tests := []struct {
		name string
		num  Base[int]
		want bool
	}{
		{"is odd", Int(7), true},
		{"is not odd (even)", Int(0), false},
		{"negative odd", Int(-3), true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.Odd()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestBaseEven(t *go_testing.T) {
	tests := []struct {
		name string
		num  Base[int]
		want bool
	}{
		{"is even", Int(8), true},
		{"is not even (odd)", Int(7), false},
		{"negative even", Int(-4), true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.Even()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestBaseAbsolute(t *go_testing.T) {
	tests := []struct {
		name string
		num  Base[int]
		want int
	}{
		{"positive", Int(7), 7},
		{"negative", Int(-7), 7},
		{"zero", Int(0), 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.Absolute()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestBaseString(t *go_testing.T) {
	tests := []struct {
		name string
		num  Base[int]
		want string
	}{
		{"positive", Int(7), "7"},
		{"negative", Int(-10), "-10"},
		{"zero", Int(0), "0"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.String()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func TestBaseAnd(t *go_testing.T) {
	tests := []struct {
		name string
		num  Base[int]
		and  int
		want int
	}{
		{"positive", Int(5), 1, 1},
		{"zero result", Int(4), 3, 0},
		{"negative number", Int(-5), 3, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.And(test.and)
			want := Int(test.want)
			if !got.Equal(want.Get()) {
				t.Errorf("got %v, want %v", got.Get(), want.Get())
			}
		})
	}
}

func TestBaseOr(t *go_testing.T) {
	tests := []struct {
		name string
		num  Base[int]
		or   int
		want int
	}{
		{"positive", Int(5), 1, 5},
		{"or results in original", Int(5), 4, 5},
		{"negative number", Int(-5), 2, -5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.Or(test.or)
			want := Int(test.want)
			if !got.Equal(want.Get()) {
				t.Errorf("got %v, want %v", got.Get(), want.Get())
			}
		})
	}
}

func TestBaseXor(t *go_testing.T) {
	tests := []struct {
		name string
		num  Base[int]
		xor  int
		want int
	}{
		{"positive", Int(5), 1, 4},
		{"with zero", Int(5), 0, 5},
		{"negative number", Int(-5), 2, -7},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.Xor(test.xor)
			want := Int(test.want)
			if !got.Equal(want.Get()) {
				t.Errorf("got %v, want %v", got.Get(), want.Get())
			}
		})
	}
}

func TestBaseShiftLeft(t *go_testing.T) {
	tests := []struct {
		name string
		num  Base[int]
		bits int
		want int
	}{
		{"positive", Int(5), 1, 10},
		{"shift of zero", Int(5), 0, 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.ShiftLeft(test.bits)
			want := Int(test.want)
			if !got.Equal(want.Get()) {
				t.Errorf("got %v, want %v", got.Get(), want.Get())
			}
		})
	}
}

func TestBaseShiftRight(t *go_testing.T) {
	tests := []struct {
		name string
		num  Base[int]
		bits int
		want int
	}{
		{"positive", Int(5), 1, 2},
		{"shift of zero", Int(5), 0, 5},
		{"large shift", Int(100), 10, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.ShiftRight(test.bits)
			want := Int(test.want)
			if !got.Equal(want.Get()) {
				t.Errorf("got %v, want %v", got.Get(), want.Get())
			}
		})
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

func TestBaseError(t *go_testing.T) {
	num := Int(10).Divide(0)
	if num.Error().Nil() {
		t.Fatal("error should not be nil")
	}

	t.Run("TestErrorAdd", func(t *go_testing.T) {
		if num.Add(5).Error().Nil() {
			t.Error("error should be propagated")
		}
	})

	t.Run("TestErrorMinus", func(t *go_testing.T) {
		if num.Minus(5).Error().Nil() {
			t.Error("error should be propagated")
		}
	})

	t.Run("TestErrorMultiply", func(t *go_testing.T) {
		if num.Multiply(5).Error().Nil() {
			t.Error("error should be propagated")
		}
	})

	t.Run("TestErrorDivide", func(t *go_testing.T) {
		if num.Divide(5).Error().Nil() {
			t.Error("error should be propagated")
		}
	})

	t.Run("TestErrorAnd", func(t *go_testing.T) {
		if num.And(5).Error().Nil() {
			t.Error("error should be propagated")
		}
	})

	t.Run("TestErrorOr", func(t *go_testing.T) {
		if num.Or(5).Error().Nil() {
			t.Error("error should be propagated")
		}
	})

	t.Run("TestErrorXor", func(t *go_testing.T) {
		if num.Xor(5).Error().Nil() {
			t.Error("error should be propagated")
		}
	})

	t.Run("TestErrorShiftLeft", func(t *go_testing.T) {
		if num.ShiftLeft(1).Error().Nil() {
			t.Error("error should be propagated")
		}
	})

	t.Run("TestErrorShiftRight", func(t *go_testing.T) {
		if num.ShiftRight(1).Error().Nil() {
			t.Error("error should be propagated")
		}
	})

	t.Run("TestErrorSet", func(t *go_testing.T) {
		if num.Set(1).Error().Nil() {
			t.Error("error should be propagated")
		}
	})

	t.Run("TestErrorPrime", func(t *go_testing.T) {
		if num.Prime() {
			t.Error("prime should be false")
		}
	})

	t.Run("TestErrorAbsolute", func(t *go_testing.T) {
		if num.Absolute() != 0 {
			t.Error("absolute should be 0")
		}
	})

	t.Run("TestErrorEven", func(t *go_testing.T) {
		if num.Even() {
			t.Error("even should be false")
		}
	})

	t.Run("TestErrorOdd", func(t *go_testing.T) {
		if num.Odd() {
			t.Error("odd should be false")
		}
	})

	t.Run("TestErrorCompare", func(t *go_testing.T) {
		if num.Compare(5) != 0 {
			t.Error("compare should be 0")
		}
	})

	t.Run("TestErrorLimits", func(t *go_testing.T) {
		min, max := num.Limits()
		if min != 0 || max != 0 {
			t.Error("limits should be 0,0")
		}
	})

	t.Run("TestErrorSign", func(t *go_testing.T) {
		if num.Sign() != 0 {
			t.Error("sign should be 0")
		}
	})

}

func TestBasePrime(t *go_testing.T) {
	tests := []struct {
		name string
		num  Base[int]
		want bool
	}{
		{"is prime", Int(2), true},
		{"is prime 1", Int(3), true},
		{"is not prime (even)", Int(4), false},
		{"is not prime (negative)", Int(-5), false},
		{"is not prime (1)", Int(1), false},
		{"is not prime (large)", Int(15), false},
		{"large prime", Int(101), true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *go_testing.T) {
			got := test.num.Prime()
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
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
