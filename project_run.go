package james

import (
	"path/filepath"

	"github.com/pieterclaerhout/go-log"
)

func (project Project) DoRun(args []string) error {

	log.Debug("Running: run")
	log.Debug("Project path:", project.Path)

	if err := project.DoBuild(); err != nil {
		return err
	}

	config, err := NewConfigFromDir(project.Path)
	if err != nil {
		return err
	}

	runCmd := []string{filepath.Join(project.Path, config.Build.OuputName)}

	return project.runCommandToStdout(runCmd)

}
