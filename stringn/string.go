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
	"slices"
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
	rs := s.runes()
	if len(rs) == 0 {
		return s.zero()
	}
	return rs[len(rs)-1]
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
	updatedRs, err := s.insertRune(i, r, rs)
	if err != nil {
		s.err = err
		return s.zero()
	}
	s.updateBuffer(updatedRs)
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

func (s *String) Clear() {
	if s.err != nil {
		return
	}
	s.b.Reset()
}

func (s *String) Prepend(strs ...string) []string {
	if s.err != nil {
		return []string{}
	}
	currentStr := s.b.String()
	var prependedStr string
	for _, str := range strs {
		prependedStr += str
	}
	s.b.Reset()
	s.b.WriteString(prependedStr + currentStr)
	return []string{s.b.String()}
}

func (s *String) Remove(start, end int) rune {
	if s.err != nil {
		return rune(0)
	}
	rs := s.runes()
	if start < 0 || end > len(rs) || start > end {
		s.err = ErrInvalidIndexForRemoval
		return s.zero()
	}
	removedRune := rs[start]
	s.updateBuffer(append(rs[:start], rs[end:]...))
	return removedRune
}

func (s *String) Replace(old, new string) {
	if s.err != nil {
		return
	}
	newStr := strings.ReplaceAll(s.b.String(), old, new)
	s.b.Reset()
	s.b.WriteString(newStr)
}

func (s *String) Reverse() {
	rs := s.runes()
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	s.updateBuffer(rs)
}

func (s *String) ToLower() string {
	if s.err != nil {
		return ""
	}
	return strings.ToLower(s.b.String())
}

func (s *String) ToUpper() string {
	if s.err != nil {
		return ""
	}
	return strings.ToUpper(s.b.String())
}

func (s *String) Compare(str string) int {
	if s.err != nil {
		return 0
	}
	return strings.Compare(s.b.String(), str)
}

func (s *String) Equals(str string) bool {
	if s.err != nil {
		return false
	}
	return s.b.String() == str
}

func (s *String) Split(sep string) []string {
	return strings.Split(s.b.String(), sep)
}

func (s *String) Substring(start, end int) string {
	if s.err != nil {
		return ""
	}
	rs := s.runes()
	if start < 0 {
		start = 0
	}
	if end > len(rs) {
		end = len(rs)
	}
	if start > end {
		return ""
	}
	return string(rs[start:end])
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
	return utf8.RuneError
}

func (*String) insertRune(i int, elem rune, s []rune) ([]rune, error) {
	if i < 0 || i > len(s) {
		zero := rune(0)
		return []rune{zero}, ErrIndexOutOfBounds
	}
	return slices.Insert(s, i, elem), nil
}
