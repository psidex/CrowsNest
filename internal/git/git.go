package git

import (
	"bytes"
	"io"
	"os"
	"os/exec"
)

func BinaryExists() bool {
	_, err := exec.LookPath("git")
	return err == nil
}

// Pull runs `git pull` in the given directory, pass "" for current working dir.
func Pull(dir string, toStdout bool, flags []string) (string, error) {
	flags = append([]string{"pull"}, flags...)
	cmd := exec.Command("git", flags...)

	var stdBuffer bytes.Buffer
	var writer io.Writer

	if toStdout {
		writer = io.MultiWriter(os.Stdout, &stdBuffer)
	} else {
		writer = &stdBuffer
	}

	cmd.Stdout = writer
	cmd.Stderr = writer

	if dir != "" {
		cmd.Dir = dir
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		return "", err
	}

	return stdBuffer.String(), nil
}
