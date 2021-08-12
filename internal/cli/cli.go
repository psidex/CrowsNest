package cli

import (
	"bytes"
	"os/exec"
)

// RunCmd runs the given name with the given flags in the given dir, returns the output and exit code.
func RunCmd(name string, flags []string, dir string) (string, int, error) {
	cmd := exec.Command(name, flags...)

	var stdBuffer bytes.Buffer
	writer := &stdBuffer

	cmd.Stdout = writer
	cmd.Stderr = writer

	if dir != "" {
		cmd.Dir = dir
	}

	if err := cmd.Start(); err != nil {
		return "", 1, err
	}

	if err := cmd.Wait(); err != nil {
		return "", cmd.ProcessState.ExitCode(), err
	}

	return stdBuffer.String(), cmd.ProcessState.ExitCode(), nil
}
