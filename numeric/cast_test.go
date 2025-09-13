package numeric

import (
	"math"
	"testing"
)

func TestCast(t *testing.T) {
	t.Run("int to int32 success", func(t *testing.T) {
		num := Int(100)
		res, err := Cast[int, int32](num)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if res.Get() != 100 {
			t.Fatalf("expected 100, got %d", res.Get())
		}
	})

	t.Run("int32 to int8 success", func(t *testing.T) {
		num := Int32(127)
		res, err := Cast[int32, int8](num)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if res.Get() != 127 {
			t.Fatalf("expected 127, got %d", res.Get())
		}
	})

	t.Run("int32 to int8 overflow", func(t *testing.T) {
		num := Int32(128)
		_, err := Cast[int32, int8](num)
		if err == nil {
			t.Fatal("expected an error, got nil")
		}
		if err.Error() != ErrOverflow {
			t.Fatalf("expected ErrOverflow, got %v", err.Error())
		}
	})

	t.Run("int to uint success", func(t *testing.T) {
		num := Int(100)
		res, err := Cast[int, uint](num)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if res.Get() != 100 {
			t.Fatalf("expected 100, got %d", res.Get())
		}
	})

	t.Run("negative int to uint overflow", func(t *testing.T) {
		num := Int(-1)
		_, err := Cast[int, uint](num)
		if err == nil {
			t.Fatal("expected an error, got nil")
		}
		if err.Error() != ErrOverflow {
			t.Fatalf("expected ErrOverflow, got %v", err.Error())
		}
	})

	t.Run("uint to int success", func(t *testing.T) {
		num := UInt(100)
		res, err := Cast[uint, int](num)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if res.Get() != 100 {
			t.Fatalf("expected 100, got %d", res.Get())
		}
	})

	t.Run("large uint to int overflow", func(t *testing.T) {
		num := UInt(math.MaxUint)
		_, err := Cast[uint, int](num)
		if err == nil {
			t.Fatal("expected an error, got nil")
		}
		if err.Error() != ErrOverflow {
			t.Fatalf("expected ErrOverflow, got %v", err.Error())
		}
	})

	t.Run("uint to uint8 success", func(t *testing.T) {
		num := UInt(255)
		res, err := Cast[uint, uint8](num)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if res.Get() != 255 {
			t.Fatalf("expected 255, got %d", res.Get())
		}
	})

	t.Run("uint to uint8 overflow", func(t *testing.T) {
		num := UInt(256)
		_, err := Cast[uint, uint8](num)
		if err == nil {
			t.Fatal("expected an error, got nil")
		}
		if err.Error() != ErrOverflow {
			t.Fatalf("expected ErrOverflow, got %v", err.Error())
		}
	})
}
