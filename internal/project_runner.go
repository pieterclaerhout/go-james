package internal

import (
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

type projectRunner struct {
	common.CommandRunner
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

	return runner.RunToStdout(runCmd, project.Path)

}
