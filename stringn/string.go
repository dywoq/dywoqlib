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

package stringn

import (
	"bytes"
	"strings"
	"unicode/utf8"

	"github.com/dywoq/dywoqlib/iterator"
	"github.com/dywoq/dywoqlib/sliceutil"
)

type String struct {
	err error
	b   bytes.Buffer
}

func (s *String) Error() error {
	return s.err
}

func (s *String) Native() string {
	return s.b.String()
}

func (s *String) Length() int {
	return s.b.Len()
}

func (s *String) Grow(i int) {
	if s.err != nil {
		return
	}
	if s.b.Len() == 0 {
		s.b.Grow(i)
		return
	}
	var newStr bytes.Buffer
	newStr.Grow(i)
	newStr.WriteString(s.b.String())
	s.b = newStr
}

func (s *String) Iterating() *iterator.Combined[rune] {
	return iterator.NewCombined(s.runes())
}

func (s *String) Append(strs ...string) []string {
	if s.err != nil {
		return []string{}
	}
	for _, str := range strs {
		s.b.WriteString(str)
	}
	return strs
}

func (s *String) At(i int) rune {
	if s.err != nil {
		return s.zero()
	}

	strBytes := s.b.Bytes()
	strLen := len(strBytes)

	if i < 0 || i >= utf8.RuneCount(strBytes) {
		s.err = ErrIndexOutOfBounds
		return s.zero()
	}

	runeIdx := 0
	for byteOffset := 0; byteOffset < strLen; {
		r, size := utf8.DecodeRune(strBytes[byteOffset:])
		if runeIdx == i {
			return r
		}
		runeIdx++
		byteOffset += size
	}
	s.err = ErrRuneNotFound
	return s.zero()
}

func (s *String) Front() rune {
	if s.err != nil {
		return s.zero()
	}
	res := s.At(0)
	if s.err != nil {
		return s.zero()
	}
	return res
}

func (s *String) Back() rune {
	if s.err != nil {
		return s.zero()
	}
	res := s.At(s.b.Len() - 1)
	if s.err != nil {
		return s.zero()
	}
	return res
}

func (s *String) String() string {
	if s.err != nil {
		return ""
	}
	return s.b.String()
}

func (s *String) Empty() bool {
	return s.b.Len() == 0
}

func (s *String) HasRunePrefix(str rune) bool {
	if s.err != nil {
		return false
	}
	if s.Empty() {
		return false
	}
	return strings.HasPrefix(s.String(), string(str))
}

func (s *String) HasStringPrefix(str string) bool {
	if s.err != nil {
		return false
	}
	if s.Empty() {
		return false
	}
	return strings.HasPrefix(s.String(), str)
}

func (s *String) HasRuneSuffix(str rune) bool {
	if s.err != nil {
		return false
	}
	if s.Empty() {
		return false
	}
	return strings.HasSuffix(s.String(), string(str))
}

func (s *String) HasStringSuffix(str string) bool {
	if s.err != nil {
		return false
	}
	if s.Empty() {
		return false
	}
	return strings.HasSuffix(s.String(), str)
}

func (s *String) Insert(i int, r rune) rune {
	if s.err != nil {
		return s.zero()
	}
	rs := s.runes()
	_, err := sliceutil.Insert(i, r, rs)
	if err != nil {
		s.err = err
		return s.zero()
	}
	s.updateBuffer(rs)
	return r
}

func (s *String) Set(r rune, i int) rune {
	if s.err != nil {
		return s.zero()
	}
	rs := s.runes()
	newR, err := sliceutil.Set(r, i, rs)
	if err != nil {
		s.err = err
		return s.zero()
	}
	s.updateBuffer(rs)
	return newR
}

func (s *String) ContainsRune(r rune) bool {
	if s.err != nil {
		return false
	}
	if s.b.Len() == 0 {
		return false
	}
	return strings.Contains(s.b.String(), string(r))
}

func (s *String) ContainsString(str string) bool {
	if s.err != nil {
		return false
	}
	if s.b.Len() == 0 {
		return false
	}
	return strings.Contains(s.b.String(), str)
}

func (s *String) Write(p []byte) (int, error) {
	if s.err != nil {
		return len(p), s.err
	}
	return s.b.Write(p)
}

func (s *String) Read(p []byte) (n int, err error) {
	if s.err != nil {
		return 0, s.err
	}
	return s.b.Read(p)
}

func (s *String) runes() []rune {
	numRunes := utf8.RuneCountInString(s.b.String())
	runes := make([]rune, 0, numRunes)
	for _, r := range s.b.String() {
		runes = append(runes, r)
	}
	return runes
}

func (s *String) updateBuffer(rs []rune) {
	s.b.Reset()
	for _, r := range rs {
		s.b.WriteRune(r)
	}
}

func (s *String) zero() rune {
	return 0
}
