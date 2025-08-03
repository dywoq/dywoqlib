package console

import "os/exec"

// Run runs the command with arguments if they're provided.
// It returns an output of command and an error.
func Run(command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	return cmd.CombinedOutput()
}
