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

type ConversionFunc[T any] func(...T) Maybe[T]

func Int(val ...int) Maybe[int]                      { return conversion(val...) }
func Int8(val ...int8) Maybe[int8]                   { return conversion(val...) }
func Int16(val ...int16) Maybe[int16]                { return conversion(val...) }
func Int32(val ...int32) Maybe[int32]                { return conversion(val...) }
func Int64(val ...int64) Maybe[int64]                { return conversion(val...) }
func UInt(val ...uint) Maybe[uint]                   { return conversion(val...) }
func UInt8(val ...uint8) Maybe[uint8]                { return conversion(val...) }
func UInt16(val ...uint16) Maybe[uint16]             { return conversion(val...) }
func UInt32(val ...uint32) Maybe[uint32]             { return conversion(val...) }
func UInt64(val ...uint64) Maybe[uint64]             { return conversion(val...) }
func String(val ...string) Maybe[string]             { return conversion(val...) }
func Bool(val ...bool) Maybe[bool]                   { return conversion(val...) }
func Float32(val ...float32) Maybe[float32]          { return conversion(val...) }
func Float64(val ...float64) Maybe[float64]          { return conversion(val...) }
func UIntptr(val ...uintptr) Maybe[uintptr]          { return conversion(val...) }
func Rune(val ...rune) Maybe[rune]                   { return conversion(val...) }
func Complex64(val ...complex64) Maybe[complex64]    { return conversion(val...) }
func Complex128(val ...complex128) Maybe[complex128] { return conversion(val...) }
func Error(val ...error) Maybe[error]                { return conversion(val...) }
func Byte(val ...byte) Maybe[byte]                   { return conversion(val...) }

func conversion[T any](val ...T) Maybe[T] {
	if len(val) == 0 {
		return None[T]()
	}
	return New(val[0])
}
