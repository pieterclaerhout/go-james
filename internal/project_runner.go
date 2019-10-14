package internal

import (
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal/config"
)

type projectRunner struct {
	Path    string
	project Project
	config  config.Config
}

func (runner projectRunner) execute(args []string) error {

	config := runner.config
	project := runner.project

	if err := project.DoBuild(false); err != nil {
		return err
	}

	runCmd := []string{filepath.Join(project.Path, config.Build.OuputName)}

	return project.runCommandToStdout(runCmd)

}
