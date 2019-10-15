package config

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/pieterclaerhout/go-log"
)

// FileName is the name of the config file
const FileName = "go-james.json"

// Config defines the project configuration
type Config struct {
	Project ProjectConfig `json:"project"` // Project contains the general project variables
	Build   BuildConfig   `json:"build"`   // Build contains the build specific project variables
	Test    TestConfig    `json:"test"`    // Build contains the test specific project variables
}

// ProjectConfig contains the general project variables
type ProjectConfig struct {
	Name        string `json:"name"`         // The name of the project
	Description string `json:"description"`  // The description of the project
	Package     string `json:"package"`      // The top-level package for the project
	MainPackage string `json:"main_package"` // The package path to the main entry point (the package containing main)
}

// BuildConfig contains the build specific configuration settings
type BuildConfig struct {
	OutputPath string   `json:"ouput_path"` // The path of the executable to generate
	LDFlags    []string `json:"ld_flags"`   // The ldflags to pass to the build command
	ExtraArgs  []string `json:"extra_args"` // The extra arguments to pass to the build command
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

	log.DebugDump(config, "Parsed:")

	return config, nil

}

// NewConfigFromDir parses the config from a file called "go-james.json" which should be present in the specified path
func NewConfigFromDir(path string) (Config, error) {
	configPath := filepath.Join(path, FileName)
	return NewConfigFromPath(configPath)
}

// Badges returns the list of badges which should be shown for the project
func (config Config) Badges() []Badge {

	packageName := strings.TrimRight(config.Project.Package, "/")
	relativePackageName := strings.TrimPrefix(packageName, "github.com/")

	badges := []Badge{}

	goReportBadge := Badge{
		Title: "Go Report Card",
		Link:  "https://goreportcard.com/report/" + packageName,
		Image: "https://goreportcard.com/badge/" + packageName,
	}
	badges = append(badges, goReportBadge)

	docsBadge := Badge{
		Title: "Documentation",
		Link:  "http://godoc.org/" + packageName,
		Image: "https://godoc.org/" + packageName + "?status.svg",
	}
	badges = append(badges, docsBadge)

	if strings.HasPrefix(packageName, "github.com/") {

		licenseBadge := Badge{
			Title: "License",
			Link:  "https://" + packageName + "/raw/master/LICENSE",
			Image: "https://img.shields.io/badge/license-Apache%20v2-orange.svg",
		}
		badges = append(badges, licenseBadge)

		versionBadge := Badge{
			Title: "GitHub Version",
			Link:  "https://badge.fury.io/gh/" + url.PathEscape(relativePackageName),
			Image: "https://badge.fury.io/gh/" + url.PathEscape(relativePackageName) + ".svg",
		}
		badges = append(badges, versionBadge)

		issuesBadge := Badge{
			Title: "GitHub issues",
			Link:  "https://" + packageName + "/issues",
			Image: "https://img.shields.io/github/issues/" + relativePackageName + ".svg",
		}
		badges = append(badges, issuesBadge)

	}

	return badges

}

// ShortPackageName returns the package name for the main package
func (config Config) ShortPackageName() string {
	name := filepath.Base(config.Project.Package)
	name = strings.TrimPrefix(name, "go-")
	name = strings.TrimSuffix(name, "go-")
	name = strings.ReplaceAll(name, "-", "")
	name = strings.ToLower(name)
	return name
}
