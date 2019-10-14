package internal

import (
	"strings"

	"github.com/kballard/go-shellquote"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

type projectBuilder struct {
	common.CommandRunner
	Path    string
	project Project
	config  config.Config
	verbose bool
}

func (builder projectBuilder) execute() error {

	config := builder.config

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
	return builder.RunToStdout(buildCmd, builder.Path)

}

func (builder projectBuilder) determineRevision() string {

	cmdLine := []string{"git", "rev-parse", "--short", "HEAD"}

	result, _ := builder.RunReturnOutput(cmdLine, builder.Path)
	return strings.TrimSpace(result)

}

func (builder projectBuilder) determineBranch() string {

	cmdLine := []string{"git", "rev-parse", "--abbrev-ref", "HEAD"}

	result, _ := builder.RunReturnOutput(cmdLine, builder.Path)
	return strings.TrimSpace(result)

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
