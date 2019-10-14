package james

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pieterclaerhout/go-log"
)

const configFileName = "go.json"

type Config struct {
	Project ProjectConfig `json:"project"`
	Build   BuildConfig   `json:"build"`
}

type ProjectConfig struct {
	Name       string `json:"name"`
	Package    string `json:"package"`
	Repository string `json:"repository"`
}

type BuildConfig struct {
	OuputName string   `json:"ouput_name"`
	LDFlags   []string `json:"ld_flags"`
	ExtraArgs []string `json:"extra_args"`
}

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

func NewConfigFromDir(path string) (Config, error) {
	configPath := filepath.Join(path, configFileName)
	return NewConfigFromPath(configPath)
}
