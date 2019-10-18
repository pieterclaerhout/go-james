package internal

import (
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal/builder"
	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
	"github.com/pieterclaerhout/go-log"
)

// Executor is used to execute the subcommands
type Executor struct {
	common.Logging

	Path string
}

// NewExecutor returns a new Executor instance
//
// The path to the project is changed to the absolute path is it exists
func NewExecutor(path string) Executor {

	if absPath, err := filepath.Abs(path); err == nil {
		path = absPath
	}

	return Executor{
		Path: path,
	}

}

// RunTool runs the given tool
//
// parseConfig is used to indicate if the config file should be parsed or not
func (executor Executor) RunTool(subcommand Subcommand, parseConfig bool) int {

	var cfg config.Config
	var err error

	if parseConfig {
		cfg, err = config.NewConfigFromDir(executor.Path)
		if err != nil {
			executor.LogError(err)
			return 1
		}
	}

	project := common.NewProject(executor.Path)

	if subcommand.RequiresBuild() {

		b := builder.Builder{
			Verbose: false,
		}

		if err := b.Execute(project, cfg); err != nil {
			executor.LogErrorInDebugMode(err)
			return 1
		}

	}

	log.DebugDump(subcommand, "subcommand")
	if err := subcommand.Execute(project, cfg); err != nil {
		log.Error(err)
		return 1
	}

	return 0

}
