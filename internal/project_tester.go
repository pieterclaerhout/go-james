package internal

import (
	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

type projectTester struct {
	common.CommandRunner
	config config.Config
}

func (tester projectTester) Execute(project Project, cfg config.Config) error {

	testCmd := []string{"go", "test", "-cover", "./..."}

	return tester.RunToStdout(testCmd, project.Path)

}
