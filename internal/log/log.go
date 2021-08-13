package log

import (
	"fmt"
	"log"
	"sync"
)

// WatcherLogger is for logging from within watcher.Watch goroutines.
type WatcherLogger struct {
	id       int
	repoName string
	mu       *sync.Mutex
}

// stdoutmu controls package level access to stdout.
var stdoutmu = sync.Mutex{}

// NewWatcher creates a new WatcherLogger.
func NewWatcher(id int, repoName string) WatcherLogger {
	return WatcherLogger{id, repoName, &stdoutmu}
}

// Info logs information.
func (l WatcherLogger) Info(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	l.mu.Lock()
	defer l.mu.Unlock()
	log.Printf("[Watcher %d] [Repo %s] %s", l.id, l.repoName, msg)
}
