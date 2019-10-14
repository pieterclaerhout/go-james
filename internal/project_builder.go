package internal

import (
	"strings"

	"github.com/kballard/go-shellquote"

	"github.com/pieterclaerhout/go-james/internal/config"
	"github.com/pieterclaerhout/go-log"
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

	versionInfo := map[string]string{
		"AppName":  config.Project.Name,
		"Revision": builder.determineRevision(),
		"Branch":   builder.determineBranch(),
	}

	buildCmd := []string{"go", "build"}

	if builder.verbose {
		buildCmd = append(buildCmd, "-v")
	}

	if config.Build.OuputName != "" {
		buildCmd = append(buildCmd, "-o", config.Build.OuputName)
	}

	ldFlags := config.Build.LDFlags

	for key, val := range versionInfo {
		ldFlags = append(ldFlags, builder.ldFlagForVersionInfo(key, val)...)
	}

	if len(ldFlags) > 0 {
		buildCmd = append(buildCmd, "-ldflags", shellquote.Join(ldFlags...))
	}

	if len(config.Build.ExtraArgs) > 0 {
		buildCmd = append(buildCmd, config.Build.ExtraArgs...)
	}

	buildCmd = append(buildCmd, config.Project.MainPackage)
	return project.runCommandToStdout(buildCmd)

}

func (builder projectBuilder) determineRevision() string {

	project := builder.project

	cmdLine := []string{"git", "rev-parse", "--short", "HEAD"}

	command, err := project.createCommand(cmdLine)
	if err != nil {
		if log.DebugMode {
			log.Error(err)
		}
		return ""
	}

	log.Debug("Executing:", shellquote.Join(cmdLine...))
	output, err := command.CombinedOutput()
	if err != nil {
		if log.DebugMode {
			log.Error(err)
		}
	}

	return strings.TrimSpace(string(output))

}

func (builder projectBuilder) determineBranch() string {

	project := builder.project

	cmdLine := []string{"git", "rev-parse", "--abbrev-ref", "HEAD"}

	command, err := project.createCommand(cmdLine)
	if err != nil {
		if log.DebugMode {
			log.Error(err)
		}
		return ""
	}

	log.Debug("Executing:", shellquote.Join(cmdLine...))
	output, err := command.CombinedOutput()
	if err != nil {
		if log.DebugMode {
			log.Error(err)
		}
	}

	return strings.TrimSpace(string(output))

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
