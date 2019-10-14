package internal

import (
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal/config"
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

	config, err := config.NewConfigFromDir(project.Path)
	if err != nil {
		return err
	}

	builder := projectBuilder{
		Path:    project.Path,
		config:  config,
		project: project,
		verbose: verbose,
	}

	return builder.execute()

}

// DoTest performs the tests defined in the project
func (project Project) DoTest() error {

	config, err := config.NewConfigFromDir(project.Path)
	if err != nil {
		return err
	}

	tester := projectTester{
		Path:    project.Path,
		config:  config,
		project: project,
	}

	return tester.execute()

}

// DoRun runs the project and passes the arguments to the command
func (project Project) DoRun(args []string) error {

	config, err := config.NewConfigFromDir(project.Path)
	if err != nil {
		return err
	}

	runner := projectRunner{
		Path:    project.Path,
		config:  config,
		project: project,
	}

	return runner.execute(args)

}
