package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// FileSystem is what can be injected into a subcommand when you need file system access
type FileSystem struct {
	Logging
}

// WriteTextFileIfNotExists writes a text file if it doesn't exist yet
func (fileSystem FileSystem) WriteTextFileIfNotExists(path string, data string) error {
	if fileSystem.FileExists(path) {
		return nil
	}
	if err := fileSystem.createDirForPathIfNeeded(path); err != nil {
		return err
	}
	return fileSystem.WriteTextFile(path, data)
}

// WriteTextFile creates a text with with data as its contents
func (fileSystem FileSystem) WriteTextFile(path string, data string) error {
	fileSystem.LogPathCreation("Writing:", path)
	return ioutil.WriteFile(path, []byte(data), 0666)
}

// WriteJSONFileIfNotExists writes JSON into a file if it doesn't exist yet
func (fileSystem FileSystem) WriteJSONFileIfNotExists(path string, data interface{}) error {
	if fileSystem.FileExists(path) {
		return nil
	}
	if err := fileSystem.createDirForPathIfNeeded(path); err != nil {
		return err
	}
	return fileSystem.WriteJSONFile(path, data)
}

// WriteJSONFile marshals the data into a JSON file
func (fileSystem FileSystem) WriteJSONFile(path string, data interface{}) error {

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	b = fileSystem.formatBytes(b)

	fileSystem.LogPathCreation("Writing:", path)
	return ioutil.WriteFile(path, b, 0777)

}

// formatBytes is a helper used to format JSON byte slices
func (fileSystem FileSystem) formatBytes(data []byte) []byte {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "    ")
	if err == nil {
		return out.Bytes()
	}
	return data
}

// PathExists is a helper function to check if a path exists or not
func (fileSystem FileSystem) PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// FileExists is a helper function to check if a file exists or not
func (fileSystem FileSystem) FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirExists is a helper function to check if a directory exists or not
func (fileSystem FileSystem) DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func (fileSystem FileSystem) createDirForPathIfNeeded(path string) error {
	path = filepath.Dir(path)
	if fileSystem.DirExists(path) {
		return nil
	}
	if fileSystem.FileExists(path) {
		return errors.New("Cannot create dir, file with that name exists already: " + path)
	}
	fileSystem.LogPathCreation("Writing:", path)
	return os.MkdirAll(path, 0755)
}
