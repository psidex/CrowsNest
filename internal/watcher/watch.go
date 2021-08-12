package watcher

import (
	"log"
	"sync"

	"github.com/psidex/CrowsNest/internal/config"
	"github.com/psidex/CrowsNest/internal/git"
)

// TODO: Is log safe for goroutines? Can we have diff logger for each watcher? Maybe nice 3rd party pkg?
// TODO: If pull ran over the next update check, have config opt to pull immediately?
//       E.g. interval is 1 minute but pull took 2 minutes - after this happens what do we do with the interval

func Watch(id int, wg *sync.WaitGroup, runOnce bool, repoName string, repoConfig *config.RepositoryConfig) {
	defer wg.Done()

	if !runOnce {
		log.Printf("[%d] Running watcher for %s", id, repoName)
	} else {
		log.Printf("[%d] Running watcher once for %s", id, repoName)
	}

	log.Printf("[%d] %+v", id, repoConfig)

	res, err := git.Pull(repoConfig.Directory, false, repoConfig.GitFlags)
	if err != nil {
		log.Printf("[%d] Failed to git pull %s: %s", id, repoName, err)
	} else {
		log.Printf("[%d] %s", id, res)
	}

	// Steps:
	//  - For each directory:
	//    - Spawn a goroutine and add to pile
	//    - Git pull (if checkpull, check remote then pull if needed)
	//    - Check if pull changed anything
	//  - Wait for goroutines to finish
	//  - Sleep for interval or exit (depending on loop)
}
