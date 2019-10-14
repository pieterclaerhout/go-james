package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/pieterclaerhout/go-log"
)

// FileWriter is what can be injected into a subcommand when you need to write files
type FileWriter struct{}

// WriteTextFile creates a text with with data as its contents
func (fileWriter FileWriter) WriteTextFile(path string, data string) error {
	fileWriter.logPathCreation(path)
	return ioutil.WriteFile(path, []byte(data), 0666)
}

// WriteJSONFile marshals the data into a JSON file
func (fileWriter FileWriter) WriteJSONFile(path string, data interface{}) error {

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	b = fileWriter.formatBytes(b)

	fileWriter.logPathCreation(path)
	return ioutil.WriteFile(path, b, 0777)

}

// logPathCreation logs the creation of a file path
func (fileWriter FileWriter) logPathCreation(path string) {
	log.Info("Creating:", path)
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
