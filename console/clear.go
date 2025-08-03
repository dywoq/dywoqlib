package console

import (
	"runtime"
)

// Clear clears the console screen.
func Clear() (err error) {
	goos := runtime.GOOS
	switch goos {
	case "windows":
		_, err = Run("cls")
	case "darwin", "unix":
		_, err = Run("clear")
	}
	return err
}
