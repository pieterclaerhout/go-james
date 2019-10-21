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

// RelPath returns a relative path inside the project
func (project Project) RelPath(subpath ...string) string {
	fullpath := []string{}
	fullpath = append(fullpath, project.Path)
	fullpath = append(fullpath, subpath...)
	return filepath.Join(fullpath...)
}

// Package gets the main package of the project from the go.mod file
func (project Project) Package() (string, error) {

	return "", nil

}
