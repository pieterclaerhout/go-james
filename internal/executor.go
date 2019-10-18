package internal

import (
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal/builder"
	"github.com/pieterclaerhout/go-james/internal/cleaner"
	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
	"github.com/pieterclaerhout/go-james/internal/creator"
	"github.com/pieterclaerhout/go-james/internal/debugger"
	"github.com/pieterclaerhout/go-james/internal/installer"
	"github.com/pieterclaerhout/go-james/internal/packager"
	"github.com/pieterclaerhout/go-james/internal/runner"
	"github.com/pieterclaerhout/go-james/internal/tester"
	"github.com/pieterclaerhout/go-james/internal/uninstaller"
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

// DoBuild performs a build of the project
func (executor Executor) DoBuild(outputPath string, goos string, goarch string, verbose bool) int {
	return executor.runSubcommand(builder.Builder{
		OutputPath: outputPath,
		GOOS:       goos,
		GOARCH:     goarch,
		Verbose:    verbose,
	}, true)
}

// DoPackage performs a package of the project
func (executor Executor) DoPackage(verbose bool, concurrency int) int {
	return executor.runSubcommand(packager.Packager{
		Verbose:     verbose,
		Concurrency: concurrency,
	}, true)
}

// DoClean performs a clean of the project
func (executor Executor) DoClean() int {
	return executor.runSubcommand(cleaner.Cleaner{}, true)
}

// DoTest performs the tests defined in the project
func (executor Executor) DoTest() int {
	return executor.runSubcommand(tester.Tester{}, true)
}

// DoInstall builds the executable and installs it in $GOPATH/bin
func (executor Executor) DoInstall(verbose bool) int {
	return executor.runSubcommand(installer.Installer{
		Verbose: verbose,
	}, true)
}

// DoUninstall removes the executable from $GOPATH/bin
func (executor Executor) DoUninstall() int {
	return executor.runSubcommand(uninstaller.Uninstaller{}, true)
}

// DoRun runs the project and passes the arguments to the command
func (executor Executor) DoRun(args []string) int {
	return executor.runSubcommand(runner.Runner{
		Args: args,
	}, true)
}

// DoDebug debugs the project and passes the arguments to the command
func (executor Executor) DoDebug(args []string) int {
	return executor.runSubcommand(debugger.Debugger{
		Args: args,
	}, true)
}

// DoInit initializes a project in an existing directory
func (executor Executor) DoInit() int {
	return executor.runSubcommand(creator.Creator{
		Mode: creator.InitProject,
	}, false)
}

// DoNew initializes a project in an existing directory
func (executor Executor) DoNew(path string, packageName string, name string, description string, overwrite bool, createGitRepo bool) int {
	return executor.runSubcommand(creator.Creator{
		Mode:          creator.NewProject,
		Path:          path,
		Package:       packageName,
		Name:          name,
		Description:   description,
		Overwrite:     overwrite,
		CreateGitRepo: createGitRepo,
	}, false)
}

func (executor Executor) runSubcommand(subcommand Subcommand, parseConfig bool) int {

	var cfg config.Config
	var err error

	if parseConfig {
		cfg, err = config.NewConfigFromDir(executor.Path)
		if err != nil {
			executor.LogErrorInDebugMode(err)
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
