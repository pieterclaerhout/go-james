package james

import (
	"strings"

	"github.com/pieterclaerhout/go-log"
)

func (project Project) DoBuild() error {

	log.Debug("Running: build")
	log.Debug("Project path:", project.Path)

	config, err := NewConfigFromDir(project.Path)
	if err != nil {
		return err
	}

	buildCmd := []string{"go", "build"}

	if config.Build.OuputName != "" {
		buildCmd = append(buildCmd, "-o", config.Build.OuputName)
	}

	if len(config.Build.LDFlags) > 0 {
		buildCmd = append(buildCmd, "-ldflags", strings.Join(config.Build.LDFlags, " "))
	}

	buildCmd = append(buildCmd, config.Project.Entrypoint)

	log.Info("Building:", config.Build.OuputName)

	if err := project.runCommand(buildCmd); err != nil {
		return err
	}

	return nil

}
