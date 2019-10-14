package creator

import (
	"os"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

// Mode is used to define the mode in which we run (init or new)
type Mode int

const (
	// InitProject initializes a project in an existing directory
	InitProject Mode = iota

	// NewProject creates a new project in the given path
	NewProject
)

// String translates the creator mode to a string
func (c Mode) String() string {
	return [...]string{"Init Project", "New Project"}[c]
}

// Creator implements the "init" and "new" commands
type Creator struct {
	common.CommandRunner
	common.FileWriter
	common.FileSystemChecks
	Mode Mode
}

// Execute executes the command
func (creator Creator) Execute(project common.Project, cfg config.Config) error {

	type creationStep func(project common.Project, cfg config.Config) error

	var steps = []creationStep{
		creator.createTasks,
		creator.createLicense,
	}

	for _, step := range steps {
		if err := step(project, cfg); err != nil {
			return err
		}
	}

	return nil

}

func (creator Creator) createTasks(project common.Project, cfg config.Config) error {

	tasksPath := project.RelPath(visualStudioFolderName, visualStudioCodeTasksFileName)
	if creator.FileExists(tasksPath) {
		return nil
	}

	tasks := newVisualStudioCodeTaskList()

	tasks.Add(visualStudioCodeTask{
		Label:          "build",
		Command:        "./build/go-james build",
		ProblemMatcher: []string{"$go"},
	})

	tasks.Add(visualStudioCodeTask{
		Label:          "tests",
		Command:        "./build/go-james test",
		ProblemMatcher: []string{"$go"},
	})

	tasks.Add(visualStudioCodeTask{
		Label:          "run",
		Command:        "./build/go-james run",
		ProblemMatcher: []string{"$go"},
	})

	os.MkdirAll(project.RelPath(visualStudioFolderName), 0755)

	return creator.WriteJSONFile(tasksPath, tasks)

}

func (creator Creator) createLicense(project common.Project, cfg config.Config) error {
	path := project.RelPath(licenseFileName)
	if creator.FileExists(path) {
		return nil
	}
	return creator.WriteTextFile(path, apacheLicense)
}
