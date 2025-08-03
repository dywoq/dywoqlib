package ansi

import "fmt"

// Base represents ANSI message.
type Base interface {
	fmt.Stringer
	// BgColor returns the current background color.
	BgColor() Color
	// FgColor returns the current foreground color.
	FgColor() Color
	// SetBgColor sets a color value to the ANSI message background color.
	SetBgColor(Color) Base
	// SetFgColor sets a color value to the ANSI message foreground color.
	SetFgColor(Color) Base
}
