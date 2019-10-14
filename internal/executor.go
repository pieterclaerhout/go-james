package internal

import (
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal/builder"
	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
	"github.com/pieterclaerhout/go-james/internal/creator"
	"github.com/pieterclaerhout/go-james/internal/runner"
	"github.com/pieterclaerhout/go-james/internal/tester"
	"github.com/pieterclaerhout/go-log"
)

// Executor is used to execute the subcommands
type Executor struct {
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

// DoBuild performs a build of the project
func (executor Executor) DoBuild(verbose bool) int {
	return executor.runSubcommand(builder.Builder{
		Verbose: verbose,
	})
}

// DoTest performs the tests defined in the project
func (executor Executor) DoTest() int {
	return executor.runSubcommand(tester.Tester{})
}

// DoRun runs the project and passes the arguments to the command
func (executor Executor) DoRun(args []string) int {
	return executor.runSubcommand(runner.Runner{
		Args: args,
	})
}

// DoInit initializes a project in an existing folder
func (executor Executor) DoInit() int {
	return executor.runSubcommand(creator.Creator{
		Mode: creator.InitProject,
	})
}

// DoNew initializes a project in an existing folder
func (executor Executor) DoNew() int {
	return executor.runSubcommand(creator.Creator{
		Mode: creator.NewProject,
	})
}

func (executor Executor) runSubcommand(subcommand Subcommand) int {

	cfg, err := config.NewConfigFromDir(executor.Path)
	if err != nil {
		if log.DebugMode {
			log.Error(err)
		}
		return 1
	}

	project := common.NewProject(executor.Path)

	log.DebugDump(subcommand, "subcommand")
	if err := subcommand.Execute(project, cfg); err != nil {
		if log.DebugMode {
			log.Error(err)
		}
		return 1
	}

	return 0

}
