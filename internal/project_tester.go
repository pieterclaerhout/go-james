package internal

import (
	"github.com/pieterclaerhout/go-log"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

type projectTester struct {
	common.CommandRunner
	Path    string
	project Project
	config  config.Config
}

func (tester projectTester) execute() error {

	project := tester.project

	log.Debug("Running: test")
	log.Debug("Project path:", project.Path)

	testCmd := []string{"go", "test", "-cover", "./..."}

	return tester.RunToStdout(testCmd, project.Path)

}
