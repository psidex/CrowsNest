package watcher

import (
	"log"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/psidex/CrowsNest/internal/cli"
	"github.com/psidex/CrowsNest/internal/config"
	"github.com/psidex/CrowsNest/internal/git"
)

// TODO: Is log safe for goroutines? Can we have diff logger for each watcher? Maybe nice 3rd party pkg?
// TODO: If pull ran over the next update check, have config opt to pull immediately?
//       E.g. interval is 1 minute but pull took 2 minutes - after this happens what do we do with the interval

// Watch watches the given repo and runs the pull process on it every interval.
func Watch(id int, wg *sync.WaitGroup, cnFlags config.Flags, repoName string, repoConfig *config.RepositoryConfig) {
	defer wg.Done()

	firstRun := true
	sleepTime := time.Duration(repoConfig.Interval) * time.Second

	if !cnFlags.RunOnce {
		log.Printf("[%d] Running watcher for %s", id, repoName)
	} else {
		log.Printf("[%d] Running watcher once for %s", id, repoName)
	}
	if cnFlags.Verbose {
		log.Printf("[%d] Loaded configuration: %s", id, spew.Sdump(repoConfig))
	}

	for {
		if !firstRun {
			if cnFlags.RunOnce {
				return
			}
			time.Sleep(sleepTime)
		} else {
			firstRun = false
		}

		// TODO: This block is repeated twice, maybe move into function.
		if repoConfig.PrePullCmd.BinaryPath != "" {
			log.Printf("[%d] Running PrePullCmd", id)
			output, exitcode, err := cli.RunCmd(
				repoConfig.PrePullCmd.BinaryPath,
				repoConfig.PrePullCmd.Flags,
				repoConfig.PrePullCmd.WorkingDirectory,
			)
			if output != "" && cnFlags.Verbose {
				log.Printf("[%d] PrePullCmd output: %s", id, output)
			}
			if err != nil {
				log.Printf("[%d] PrePullCmd err: %s", id, err)
				continue
			}
			if exitcode != 0 {
				log.Printf("[%d] PrePullCmd returned non-zero exit code: %d", id, exitcode)
				continue
			}
		}

		gitOutput, err := git.Pull(repoConfig.GitPullFlags, repoConfig.Directory)
		log.Printf("[%d] git pull output: %s", id, gitOutput)
		if err != nil {
			log.Printf("[%d] Failed to git pull %s: %s", id, repoName, err)
		}

		if repoConfig.PostPullCmd.BinaryPath != "" {
			log.Printf("[%d] Running PostPullCmd", id)
			output, exitcode, err := cli.RunCmd(
				repoConfig.PostPullCmd.BinaryPath,
				repoConfig.PostPullCmd.Flags,
				repoConfig.PostPullCmd.WorkingDirectory,
			)
			if output != "" && cnFlags.Verbose {
				log.Printf("[%d] PostPullCmd output: %s", id, output)
			}
			if err != nil {
				log.Printf("[%d] PostPullCmd err: %s", id, err)
			}
			if exitcode != 0 {
				log.Printf("[%d] PostPullCmd returned non-zero exit code: %d", id, exitcode)
			}
		}

		// TODO: Check and log if anything changed?
	}
}
