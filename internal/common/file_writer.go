package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/pieterclaerhout/go-log"
)

// FileWriter is what can be injected into a subcommand when you need to write files
type FileWriter struct{}

// WriteTextFile creates a text with with data as its contents
func (fileWriter FileWriter) WriteTextFile(project Project, path string, data string) error {
	path = filepath.Join(project.Path, path)
	fileWriter.logPathCreation(project, path)
	return ioutil.WriteFile(path, []byte(data), 0666)
}

// WriteJSONFile marshals the data into a JSON file
func (fileWriter FileWriter) WriteJSONFile(project Project, path string, data interface{}) error {

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	b = fileWriter.formatBytes(b)

	path = filepath.Join(project.Path, path)
	fileWriter.logPathCreation(project, path)
	return ioutil.WriteFile(path, b, 0777)

}

// logPathCreation logs the creation of a file path
func (fileWriter FileWriter) logPathCreation(project Project, path string) {
	relPath, err := filepath.Rel(project.Path, path)
	if err != nil {
		relPath = path
	}
	log.Info("Creating:", relPath)
}

// formatBytes is a helper used to format JSON byte slices
func (fileWriter FileWriter) formatBytes(data []byte) []byte {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "    ")
	if err == nil {
		return out.Bytes()
	}
	return data
}
