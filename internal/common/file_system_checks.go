package common

import (
	"os"
	"path/filepath"
)

// FileSystemChecks is what can be injected into a subcommand when you need to check the existence of files and folder
type FileSystemChecks struct{}

// FileExists is a helper function to check if a file exists or not
func (fileSystemChecks FileSystemChecks) FileExists(project Project, path string) bool {
	path = filepath.Join(project.Path, path)
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// FolderExists is a helper function to check if a folder exists or not
func (fileSystemChecks FileSystemChecks) FolderExists(project Project, path string) bool {
	path = filepath.Join(project.Path, path)
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
