package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pieterclaerhout/go-log"
)

const configFileName = "go.json"

// Config defines the project configuration
type Config struct {
	Project ProjectConfig `json:"project"` // Project contains the general project variables
	Build   BuildConfig   `json:"build"`   // Build contains the build specific project variables
}

// ProjectConfig contains the general project variables
type ProjectConfig struct {
	Name        string `json:"name"`         // The name of the project
	Package     string `json:"package"`      // The top-level package for the project
	MainPackage string `json:"main_package"` // The package path to the main entry point (the package containing main)
	Repository  string `json:"repository"`   // The path to the the Git repository of the project
}

// BuildConfig contains the build specific configuration settings
type BuildConfig struct {
	OuputName string   `json:"ouput_name"` // The name of the executable to generate
	LDFlags   []string `json:"ld_flags"`   // The ldflags to pass to the build command
	ExtraArgs []string `json:"extra_args"` // The extra arguments to pass to the build command
}

// NewConfigFromPath reads the configuration file from the specified path and returns the configuration settings
func NewConfigFromPath(path string) (Config, error) {

	var config Config

	log.Debug("Parsing:", path)

	jsonFile, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer jsonFile.Close()

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return config, err
	}

	if err := json.Unmarshal(jsonBytes, &config); err != nil {
		return config, err
	}

	log.DebugDump(config, "Parsed:")

	return config, nil

}

// NewConfigFromDir parses the config from a file called "go.json" which should be present in the specified path
func NewConfigFromDir(path string) (Config, error) {
	configPath := filepath.Join(path, configFileName)
	return NewConfigFromPath(configPath)
}
