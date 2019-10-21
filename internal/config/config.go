package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pieterclaerhout/go-log"
	"github.com/pkg/errors"
)

// FileName is the name of the config file
const FileName = "go-james.json"

// Config defines the project configuration
type Config struct {
	Project ProjectConfig `json:"project"` // Project contains the general project variables
	Build   BuildConfig   `json:"build"`   // Build contains the build specific project variables
	Package PackageConfig `json:"package"` // Build contains the package specific project variables
	Test    TestConfig    `json:"test"`    // Build contains the test specific project variables
}

// ProjectConfig contains the general project variables
type ProjectConfig struct {
	Name        string `json:"name"`        // The name of the project
	Version     string `json:"version"`     // The version of the project
	Description string `json:"description"` // The description of the project
	Copyright   string `json:"copyright"`   // The copyright statement for the project
	// Package     string `json:"package"`      // The top-level package for the project
	MainPackage string `json:"main_package"` // The package path to the main entry point (the package containing main)
}

// BuildConfig contains the build specific configuration settings
type BuildConfig struct {
	OutputPath string   `json:"ouput_path"` // The path of the executable to generate
	LDFlags    []string `json:"ld_flags"`   // The ldflags to pass to the build command
	ExtraArgs  []string `json:"extra_args"` // The extra arguments to pass to the build command
}

// PackageConfig contains the build specific configuration settings
type PackageConfig struct {
	IncludeReadme bool `json:"include_readme"` // Include the readme when packaging or not
}

// TestConfig contains the test specific configuration settings
type TestConfig struct {
	ExtraArgs []string `json:"extra_args"` // The extra arguments to pass to the test command
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

	if config.Project.Name == "" {
		return config, errors.New("Config setting Project.Name should not be empty")
	}

	if config.Project.Version == "" {
		return config, errors.New("Config setting Project.Version should not be empty")
	}

	// if config.Project.Package == "" {
	// 	return config, errors.New("Config setting Project.Package should not be empty")
	// }

	if config.Project.MainPackage == "" {
		return config, errors.New("Config setting Project.MainPackage should not be empty")
	}

	log.DebugDump(config, "Parsed:")

	return config, nil

}

// NewConfigFromDir parses the config from a file called "go-james.json" which should be present in the specified path
func NewConfigFromDir(path string) (Config, error) {
	configPath := filepath.Join(path, FileName)
	return NewConfigFromPath(configPath)
}
