package log

import (
	"fmt"
	"log"
	"sync"

	"github.com/psidex/CrowsNest/internal/config"
)

// WatcherLogger is for logging from within watcher.Watch goroutines.
type WatcherLogger struct {
	logString string
	mu        *sync.Mutex
}

// stdoutmu controls package level access to stdout.
var stdoutmu = sync.Mutex{}

// NewWatcher creates a new WatcherLogger.
func NewWatcher(id int, repoName string, cnFlags config.Flags) WatcherLogger {
	var out string

	if cnFlags.Verbose {
		out = fmt.Sprintf("[Watcher %d] [Repo %s]", id, repoName)
	} else {
		out = fmt.Sprintf("[Repo %s]", repoName)
	}

	// Where msg will go.
	out += " %s"

	return WatcherLogger{out, &stdoutmu}
}

// Info logs information.
func (l WatcherLogger) Info(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	l.mu.Lock()
	defer l.mu.Unlock()
	log.Printf(l.logString, msg)
}
