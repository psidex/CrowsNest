package watcher

import (
	"fmt"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/psidex/CrowsNest/internal/cli"
	"github.com/psidex/CrowsNest/internal/config"
	"github.com/psidex/CrowsNest/internal/git"
	"github.com/psidex/CrowsNest/internal/log"
)

// runExternal uses cli.RunCmd to run a user provided binary.
func runExternal(logger log.WatcherLogger, cnFlags config.Flags, cmd config.CliBinOpts, name string) error {
	if cmd.BinaryPath != "" {
		logger.Info("%s running", name)
		output, exitcode, err := cli.RunCmd(
			cmd.BinaryPath,
			cmd.Flags,
			cmd.WorkingDirectory,
		)
		if output != "" {
			logger.InfoV("%s output: %s", name, output)
		}
		if exitcode != 0 {
			errStr := fmt.Sprintf("returned non-zero exit code: %d", exitcode)
			return fmt.Errorf(errStr)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// Watch watches the given repo and runs the pull process on it every interval.
func Watch(id int, wg *sync.WaitGroup, cnFlags config.Flags, repoName string, repoConfig *config.RepositoryConfig) {
	defer wg.Done()

	// NOTE: Each watcher goroutine runs its own logger with its own config, but all
	//  instances of log.WatcherLogger are linked to one mutex for writing to stdout/file.
	logger := log.NewWatcher(id, repoName, cnFlags)

	firstRun := true
	sleepTime := time.Duration(repoConfig.Interval) * time.Second

	if !cnFlags.RunOnce {
		logger.Info("Running watcher")
	} else {
		logger.Info("Running watcher once")
	}

	logger.InfoV("Loaded configuration: %s", spew.Sdump(repoConfig))

	for {
		if !firstRun {
			if cnFlags.RunOnce {
				return
			}
			time.Sleep(sleepTime)
		} else {
			firstRun = false
		}

		err := runExternal(logger, cnFlags, repoConfig.PrePullCmd, "PrePullCmd")
		if err != nil {
			logger.Info("PrePullCmd error: %s", err)
			continue
		}

		logger.InfoV("Performing git pull")

		gitOutput, err := git.Pull(repoConfig.GitPullFlags, repoConfig.Directory)
		logger.InfoV("Git pull output: %s", gitOutput)
		if err != nil {
			logger.Info("Failed to git pull: %s", err)
		}

		if git.PullPulled(gitOutput) {
			logger.Info("Git pull resulted in local change")
		}

		err = runExternal(logger, cnFlags, repoConfig.PostPullCmd, "PostPullCmd")
		if err != nil {
			logger.Info("PostPullCmd error: %s", err)
		}
	}
}
