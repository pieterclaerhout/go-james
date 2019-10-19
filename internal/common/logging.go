package common

import (
	"os"
	"path/filepath"

	"github.com/pieterclaerhout/go-log"
)

// Logging is what can be injected into a subcommand when you need logging
type Logging struct{}

// LogPathCreation logs the creation of a file path
func (logging Logging) LogPathCreation(prefix string, path string) {
	if wd, err := os.Getwd(); err == nil {
		if relPath, err := filepath.Rel(wd, path); err == nil {
			path = relPath
		}
	}
	log.Info(prefix, path)
}

// LogInfo logs an info message
func (logging Logging) LogInfo(args ...interface{}) {
	log.Info(args...)
}

// LogError logs an error
func (logging Logging) LogError(args ...interface{}) {
	log.Error(args...)
}

// LogErrorInDebugMode logs an error when running in debug mode
func (logging Logging) LogErrorInDebugMode(args ...interface{}) {
	if log.DebugMode {
		log.Error(args...)
	}
}
