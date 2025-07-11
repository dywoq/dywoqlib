package console

import (
	"os"
	"os/exec"
	"runtime"
)

// Clear clears the current console screen. If it's not supported platform,
// it returns a error, mainly ErrNotSupportedPlatform.
func Clear() error {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		return nil
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		return nil
	default:
		return ErrNotSupportedPlatform
	}
}
