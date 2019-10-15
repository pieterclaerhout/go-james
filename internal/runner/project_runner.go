package runner

import (
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

// Runner implements the "run" command
type Runner struct {
	common.CommandRunner
	Args []string
}

// Execute executes the command
func (runner Runner) Execute(project common.Project, cfg config.Config) error {

	runCmd := []string{filepath.Join(project.Path, cfg.Build.OutputPath)}
	runCmd = append(runCmd, runner.Args...)

	return runner.RunToStdout(runCmd, project.Path, map[string]string{})

}

// RequiresBuild indicates if a build is required before running the command
func (runner Runner) RequiresBuild() bool {
	return true
}
