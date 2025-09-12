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

package ansi

import (
	"fmt"

	"github.com/dywoq/dywoqlib/attribute"
)

// Color represents ANSI color type, always equivalent to int8.
//
// DEPRECATED, MAY BE REMOVED IN THE FUTURE
type Color int8

// DEPRECATED, MAY BE REMOVED IN THE FUTURE
const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	Reset Color = 9
	None  Color = 10
)

const resetCode = "\x1b[0m"

// FgFrom returns the ANSI escape sequence of foreground with the given color.
//
// DEPRECATED, MAY BE REMOVED IN THE FUTURE
func FgFrom(c Color) string {
	attribute.Deprecated(nil)
	return fmt.Sprintf("\x1b[3%dm", c)
}

// FgFrom returns the ANSI escape sequence of background with the given color
//
// DEPRECATED, MAY BE REMOVED IN THE FUTURE
func BgFrom(c Color) string {
	attribute.Deprecated(nil)
	return fmt.Sprintf("\x1b[4%dm", c)
}

// ApplyFg returns the string wrapped around ANSI escape sequences
// of foreground with given color.
//
// DEPRECATED, MAY BE REMOVED IN THE FUTURE
func ApplyFg(value string, c Color) string {
	attribute.Deprecated(nil)
	return fmt.Sprintf("%s%s%s", FgFrom(c), value, resetCode)
}

// ApplyBg returns the string wrapped around ANSI escape sequences
// of background with given color.
//
// DEPRECATED, MAY BE REMOVED IN THE FUTURE
func ApplyBg(value string, c Color) string {
	attribute.Deprecated(nil)
	return fmt.Sprintf("%s%s%s", BgFrom(c), value, resetCode)
}

// ApplyBoth returns the string wrapped around ANSI escape sequences
// of backgroud and foreground with the given colors.
//
// DEPRECATED, MAY BE REMOVED IN THE FUTURE
func ApplyBoth(value string, textColor, bgColor Color) string {
	attribute.Deprecated(nil)
	return fmt.Sprintf("%s%s%s%s", FgFrom(textColor), BgFrom(bgColor), value, resetCode)
}
