package internal

import (
	"strings"

	"github.com/kballard/go-shellquote"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

type projectBuilder struct {
	common.CommandRunner
	config  config.Config
	verbose bool
}

func (builder projectBuilder) Execute(project Project, cfg config.Config) error {

	versionInfo := map[string]string{
		"AppName":  cfg.Project.Name,
		"Revision": builder.determineRevision(project),
		"Branch":   builder.determineBranch(project),
	}

	buildCmd := []string{"go", "build"}

	if builder.verbose {
		buildCmd = append(buildCmd, "-v")
	}

	if cfg.Build.OuputName != "" {
		buildCmd = append(buildCmd, "-o", cfg.Build.OuputName)
	}

	ldFlags := cfg.Build.LDFlags

	for key, val := range versionInfo {
		ldFlags = append(ldFlags, builder.ldFlagForVersionInfo(key, val)...)
	}

	if len(ldFlags) > 0 {
		buildCmd = append(buildCmd, "-ldflags", shellquote.Join(ldFlags...))
	}

	if len(cfg.Build.ExtraArgs) > 0 {
		buildCmd = append(buildCmd, cfg.Build.ExtraArgs...)
	}

	buildCmd = append(buildCmd, cfg.Project.MainPackage)
	return builder.RunToStdout(buildCmd, project.Path)

}

func (builder projectBuilder) determineRevision(project Project) string {

	cmdLine := []string{"git", "rev-parse", "--short", "HEAD"}

	result, _ := builder.RunReturnOutput(cmdLine, project.Path)
	return strings.TrimSpace(result)

}

func (builder projectBuilder) determineBranch(project Project) string {

	cmdLine := []string{"git", "rev-parse", "--abbrev-ref", "HEAD"}

	result, _ := builder.RunReturnOutput(cmdLine, project.Path)
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
