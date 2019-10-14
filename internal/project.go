package internal

import (
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal/config"
	"github.com/pieterclaerhout/go-log"
)

// Project defines a Go project based on Go modules
type Project struct {
	Path string
}

// NewProject returns a new project instance
//
// The path to the project is changed to the absolute path is it exists
func NewProject(path string) Project {

	if absPath, err := filepath.Abs(path); err == nil {
		path = absPath
	}

	return Project{
		Path: path,
	}

}

// DoBuild performs a build of the project
func (project Project) DoBuild(verbose bool) error {
	return project.runSubcommand(projectBuilder{
		verbose: verbose,
	})
}

// DoTest performs the tests defined in the project
func (project Project) DoTest() error {
	return project.runSubcommand(projectTester{})
}

// DoRun runs the project and passes the arguments to the command
func (project Project) DoRun(args []string) error {
	return project.runSubcommand(projectRunner{
		args: args,
	})
}

func (project Project) runSubcommand(subcommand Subcommand) error {

	cfg, err := config.NewConfigFromDir(project.Path)
	if err != nil {
		return err
	}

	log.DebugDump(subcommand, "subcommand")
	return subcommand.Execute(project, cfg)

}
