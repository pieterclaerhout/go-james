package james

import (
	"github.com/pieterclaerhout/go-log"
)

func (project Project) DoTest() error {

	log.Debug("Running: test")
	log.Debug("Project path:", project.Path)

	testCmd := []string{"go", "test", "-cover", "./..."}

	return project.runCommandToStdout(testCmd)

}
