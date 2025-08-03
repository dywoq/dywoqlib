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

package optional

type ConversionFunc[T any] func(...T) Optional[T]

func Int(val ...int) Optional[int] {
	if len(val) == 0 {
		return None[int]()
	}
	return New(val[0])
}

func Int8(val ...int8) Optional[int8] {
	if len(val) == 0 {
		return None[int8]()
	}
	return New(val[0])
}

func Int16(val ...int16) Optional[int16] {
	if len(val) == 0 {
		return None[int16]()
	}
	return New(val[0])
}

func Int32(val ...int32) Optional[int32] {
	if len(val) == 0 {
		return None[int32]()
	}
	return New(val[0])
}

func Int64(val ...int64) Optional[int64] {
	if len(val) == 0 {
		return None[int64]()
	}
	return New(val[0])
}

func UInt(val ...uint) Optional[uint] {
	if len(val) == 0 {
		return None[uint]()
	}
	return New(val[0])
}

func UInt8(val ...uint8) Optional[uint8] {
	if len(val) == 0 {
		return None[uint8]()
	}
	return New(val[0])
}

func UInt16(val ...uint16) Optional[uint16] {
	if len(val) == 0 {
		return None[uint16]()
	}
	return New(val[0])
}

func UInt32(val ...uint32) Optional[uint32] {
	if len(val) == 0 {
		return None[uint32]()
	}
	return New(val[0])
}

func UInt64(val ...uint64) Optional[uint64] {
	if len(val) == 0 {
		return None[uint64]()
	}
	return New(val[0])
}

func String(val ...string) Optional[string] {
	if len(val) == 0 {
		return None[string]()
	}
	return New(val[0])
}

func Bool(val ...bool) Optional[bool] {
	if len(val) == 0 {
		return None[bool]()
	}
	return New(val[0])
}

func Float32(val ...float32) Optional[float32] {
	if len(val) == 0 {
		return None[float32]()
	}
	return New(val[0])
}

func Float64(val ...float64) Optional[float64] {
	if len(val) == 0 {
		return None[float64]()
	}
	return New(val[0])
}

func UIntptr(val ...uintptr) Optional[uintptr] {
	if len(val) == 0 {
		return None[uintptr]()
	}
	return New(val[0])
}

func Rune(val ...rune) Optional[rune] {
	if len(val) == 0 {
		return None[rune]()
	}
	return New(val[0])
}

func Complex64(val ...complex64) Optional[complex64] {
	if len(val) == 0 {
		return None[complex64]()
	}
	return New(val[0])
}

func Complex128(val ...complex128) Optional[complex128] {
	if len(val) == 0 {
		return None[complex128]()
	}
	return New(val[0])
}

func Error(val ...error) Optional[error] {
	if len(val) == 0 {
		return None[error]()
	}
	return New(val[0])
}
