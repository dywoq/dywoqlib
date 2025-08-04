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

func Int(val ...int) Optional[int]                      { return conversion(val...) }
func Int8(val ...int8) Optional[int8]                   { return conversion(val...) }
func Int16(val ...int16) Optional[int16]                { return conversion(val...) }
func Int32(val ...int32) Optional[int32]                { return conversion(val...) }
func Int64(val ...int64) Optional[int64]                { return conversion(val...) }
func UInt(val ...uint) Optional[uint]                   { return conversion(val...) }
func UInt8(val ...uint8) Optional[uint8]                { return conversion(val...) }
func UInt16(val ...uint16) Optional[uint16]             { return conversion(val...) }
func UInt32(val ...uint32) Optional[uint32]             { return conversion(val...) }
func UInt64(val ...uint64) Optional[uint64]             { return conversion(val...) }
func String(val ...string) Optional[string]             { return conversion(val...) }
func Bool(val ...bool) Optional[bool]                   { return conversion(val...) }
func Float32(val ...float32) Optional[float32]          { return conversion(val...) }
func Float64(val ...float64) Optional[float64]          { return conversion(val...) }
func UIntptr(val ...uintptr) Optional[uintptr]          { return conversion(val...) }
func Rune(val ...rune) Optional[rune]                   { return conversion(val...) }
func Complex64(val ...complex64) Optional[complex64]    { return conversion(val...) }
func Complex128(val ...complex128) Optional[complex128] { return conversion(val...) }
func Error(val ...error) Optional[error]                { return conversion(val...) }

func conversion[T any](val ...T) Optional[T] {
	if len(val) == 0 {
		return None[T]()
	}
	return New(val[0])
}
