package runner

import (
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal/builder"
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

	b := builder.Builder{
		Verbose: false,
	}

	if err := b.Execute(project, cfg); err != nil {
		return err
	}

	runCmd := []string{filepath.Join(project.Path, cfg.Build.OutputName)}
	runCmd = append(runCmd, runner.Args...)

	return runner.RunToStdout(runCmd, project.Path)

}
