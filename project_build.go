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

	ldFlags := config.Build.LDFlags
	if revision := project.determineRevision(); revision != "" {
		ldFlags = append(ldFlags, "-X", config.Project.Package+".version.Revision="+revision)
	}
	if branch := project.determineBranch(); branch != "" {
		ldFlags = append(ldFlags, "-X", config.Project.Package+".version.Branch="+branch)
	}

	buildCmd := []string{"go", "build"}

	if config.Build.OuputName != "" {
		buildCmd = append(buildCmd, "-o", config.Build.OuputName)
	}

	if len(ldFlags) > 0 {
		buildCmd = append(buildCmd, "-ldflags", strings.Join(ldFlags, " "))
	}

	buildCmd = append(buildCmd, config.Project.Entrypoint)

	if err := project.runCommandToStdout(buildCmd); err != nil {
		return err
	}

	return nil

}
