package runner

import (
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

	runCmd := []string{project.RelPath(cfg.Build.OutputPath, cfg.Project.Name)}
	runCmd = append(runCmd, runner.Args...)

	return runner.RunToStdout(runCmd, project.Path, cfg.Run.Environ)

}

// RequiresBuild indicates if a build is required before running the command
func (runner Runner) RequiresBuild() bool {
	return true
}
