package internal

import (
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

type projectRunner struct {
	common.CommandRunner
	config config.Config
	args   []string
}

func (runner projectRunner) Execute(project Project, cfg config.Config) error {

	if err := project.DoBuild(false); err != nil {
		return err
	}

	runCmd := []string{filepath.Join(project.Path, cfg.Build.OuputName)}
	runCmd = append(runCmd, runner.args...)

	return runner.RunToStdout(runCmd, project.Path)

}
