package git

import (
	"fmt"
	"os/exec"

	"github.com/psidex/CrowsNest/internal/cli"
)

// TODO: What to do about priv / auth required?

func BinaryExists() bool {
	_, err := exec.LookPath("git")
	return err == nil
}

// Pull runs `git pull` in the given directory, pass "" for current working dir.
func Pull(flags []string, dir string) (string, error) {
	flags = append([]string{"pull"}, flags...)
	output, exitcode, err := cli.RunCmd("git", flags, dir)
	if exitcode != 0 {
		return output, fmt.Errorf("git exited with a non-zero exit code: %d", exitcode)
	}
	return output, err
}
