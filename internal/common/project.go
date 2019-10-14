package common

import (
	"path/filepath"
)

// Project defines a Go project based on Go modules
type Project struct {
	Path string
}

// NewProject returns a new Project instance
func NewProject(path string) Project {

	if absPath, err := filepath.Abs(path); err == nil {
		path = absPath
	}

	return Project{
		Path: path,
	}

}
