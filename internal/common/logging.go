package common

import (
	"github.com/pieterclaerhout/go-log"
)

// Logging is what can be injected into a subcommand when you need logging
type Logging struct{}

// LogPathCreation logs the creation of a file path
func (logging Logging) LogPathCreation(path string) {
	log.Info("Creating:", path)
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
