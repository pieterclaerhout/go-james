package james

import (
	"github.com/kballard/go-shellquote"
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
	ldFlags = append(ldFlags, config.ldFlagForVersionInfo("AppName", config.Project.Name)...)
	if revision := project.determineRevision(); revision != "" {
		ldFlags = append(ldFlags, config.ldFlagForVersionInfo("Revision", revision)...)
	}
	if branch := project.determineBranch(); branch != "" {
		ldFlags = append(ldFlags, config.ldFlagForVersionInfo("Branch", branch)...)
	}

	buildCmd := []string{"go", "build"}

	if config.Build.OuputName != "" {
		buildCmd = append(buildCmd, "-o", config.Build.OuputName)
	}

	if len(ldFlags) > 0 {
		buildCmd = append(buildCmd, "-ldflags", shellquote.Join(ldFlags...))
	}

	buildCmd = append(buildCmd, config.Project.Entrypoint)

	return project.runCommandToStdout(buildCmd)

}
