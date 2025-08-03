package ansi

import (
	"fmt"
)

// Color represents ANSI color type, always equivalent to int8.
type Color int8

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
func FgFrom(c Color) string {
	return fmt.Sprintf("\x1b[3%dm", c)
}

// FgFrom returns the ANSI escape sequence of background with the given color
func BgFrom(c Color) string {
	return fmt.Sprintf("\x1b[4%dm", c)
}

// ApplyFg returns the string wrapped around ANSI escape sequences
// of foreground with given color.
func ApplyFg(value string, c Color) string {
	return fmt.Sprintf("%s%s%s", FgFrom(c), value, resetCode)
}

// ApplyBg returns the string wrapped around ANSI escape sequences
// of background with given color.
func ApplyBg(value string, c Color) string {
	return fmt.Sprintf("%s%s%s", BgFrom(c), value, resetCode)
}

// ApplyBoth returns the string wrapped around ANSI escape sequences
// of backgroud and foreground with the given colors.
func ApplyBoth(value string, textColor, bgColor Color) string {
	return fmt.Sprintf("%s%s%s%s", FgFrom(textColor), BgFrom(bgColor), value, resetCode)
}
