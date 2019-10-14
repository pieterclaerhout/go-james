package common

import (
	"os"
)

// FileSystemChecks is what can be injected into a subcommand when you need to check the existence of files and folder
type FileSystemChecks struct{}

// FileExists is a helper function to check if a file exists or not
func (fileSystemChecks FileSystemChecks) FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// FolderExists is a helper function to check if a folder exists or not
func (fileSystemChecks FileSystemChecks) FolderExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
