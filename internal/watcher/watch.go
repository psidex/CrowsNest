package watcher

import (
	"log"
	"sync"

	"github.com/psidex/CrowsNest/internal/config"
)

// TODO: Is log safe for goroutines?
// TODO: Maybe method should be in config file instead of flag?

func Watch(id int, wg *sync.WaitGroup, runOnce bool, repoName string, repoConfig *config.RepositoryConfig) {
	defer wg.Done()

	if !runOnce {
		log.Printf("Running watcher (id %d) for %s", id, repoName)
	} else {
		log.Printf("Running watcher once (id %d) for %s", id, repoName)
	}

	log.Println(repoConfig)

	// Steps:
	//  - For each directory:
	//    - Spawn a goroutine and add to pile
	//    - Git pull (if checkpull, check remote then pull if needed)
	//    - Check if pull changed anything
	//  - Wait for goroutines to finish
	//  - Sleep for interval or exit (depending on loop)
}
