package tester

import (
	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

// Tester implements the "test" command
type Tester struct {
	common.CommandRunner
}

// Execute executes the command
func (tester Tester) Execute(project common.Project, cfg config.Config) error {

	testCmd := []string{"go", "test", "-cover"}

	if len(cfg.Build.ExtraArgs) > 0 {
		testCmd = append(testCmd, cfg.Build.ExtraArgs...)
	}

	testCmd = append(testCmd, "./...")

	return tester.RunToStdout(testCmd, project.Path, map[string]string{})

}
