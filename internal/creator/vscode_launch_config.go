package creator

import (
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

const visualStudioCodeLaunchFileName = "launch.json"

type visualStudioCodeLaunchConfig struct {
	Name    string            `json:"name"`
	Type    string            `json:"type"`
	Request string            `json:"request"`
	Mode    string            `json:"mode"`
	Program string            `json:"program"`
	Env     map[string]string `json:"env"`
	Args    []string          `json:"args"`
}

type visualStudioCodeLaunchConfigs struct {
	Version        string
	Configurations []visualStudioCodeLaunchConfig `json:"configurations"`
}

func newVisualStudioCodeLaunchConfigs(cfg config.Config) *visualStudioCodeLaunchConfigs {
	result := &visualStudioCodeLaunchConfigs{
		Version: "2.0.0",
		Configurations: []visualStudioCodeLaunchConfig{
			{
				Name:    "Debug Executable",
				Type:    "go",
				Request: "launch",
				Mode:    "auto",
				Program: filepath.Join("${workspaceFolder}", common.CmdDirName, "main"),
				Env:     map[string]string{},
				Args:    []string{},
			},
		},
	}
	return result
}
