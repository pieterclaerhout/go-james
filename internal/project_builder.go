package internal

import (
	"github.com/kballard/go-shellquote"

	"github.com/pieterclaerhout/go-james/internal/config"
)

type projectBuilder struct {
	Path    string
	project Project
	config  config.Config
	verbose bool
}

func (builder projectBuilder) execute() error {

	config := builder.config
	project := builder.project

	ldFlags := config.Build.LDFlags
	ldFlags = append(ldFlags, builder.ldFlagForVersionInfo("AppName", config.Project.Name)...)
	if revision := project.determineRevision(); revision != "" {
		ldFlags = append(ldFlags, builder.ldFlagForVersionInfo("Revision", revision)...)
	}
	if branch := project.determineBranch(); branch != "" {
		ldFlags = append(ldFlags, builder.ldFlagForVersionInfo("Branch", branch)...)
	}

	buildCmd := []string{"go", "build"}

	if builder.verbose {
		buildCmd = append(buildCmd, "-v")
	}

	if config.Build.OuputName != "" {
		buildCmd = append(buildCmd, "-o", config.Build.OuputName)
	}

	if len(ldFlags) > 0 {
		buildCmd = append(buildCmd, "-ldflags", shellquote.Join(ldFlags...))
	}

	buildCmd = append(buildCmd, config.Project.MainPackage)
	return project.runCommandToStdout(buildCmd)

}

func (builder projectBuilder) ldFlagForVersionInfo(name string, value string) []string {

	config := builder.config

	result := []string{}

	if name != "" && value != "" {
		result = append(
			result,
			"-X", config.Project.Package+"/version."+name+"="+value,
		)
	}

	return result

}
