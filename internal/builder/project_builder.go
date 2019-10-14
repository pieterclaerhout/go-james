package builder

import (
	"strings"

	"github.com/kballard/go-shellquote"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

// Builder implements the "build" command
type Builder struct {
	common.CommandRunner
	Verbose bool
}

// Execute executes the command
func (builder Builder) Execute(project common.Project, cfg config.Config) error {

	versionInfo := map[string]string{
		"AppName":  cfg.Project.Name,
		"Revision": builder.determineRevision(project),
		"Branch":   builder.determineBranch(project),
	}

	buildCmd := []string{"go", "build"}

	if builder.Verbose {
		buildCmd = append(buildCmd, "-v")
	}

	if cfg.Build.OuputName != "" {
		buildCmd = append(buildCmd, "-o", cfg.Build.OuputName)
	}

	ldFlags := cfg.Build.LDFlags

	for key, val := range versionInfo {
		ldFlags = append(ldFlags, builder.ldFlagForVersionInfo(cfg, key, val)...)
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

func (builder Builder) determineRevision(project common.Project) string {

	cmdLine := []string{"git", "rev-parse", "--short", "HEAD"}

	result, _ := builder.RunReturnOutput(cmdLine, project.Path)
	return strings.TrimSpace(result)

}

func (builder Builder) determineBranch(project common.Project) string {

	cmdLine := []string{"git", "rev-parse", "--abbrev-ref", "HEAD"}

	result, _ := builder.RunReturnOutput(cmdLine, project.Path)
	return strings.TrimSpace(result)

}

func (builder Builder) ldFlagForVersionInfo(cfg config.Config, name string, value string) []string {

	result := []string{}

	if name != "" && value != "" {
		result = append(
			result,
			"-X", cfg.Project.Package+"/versioninfo."+name+"="+value,
		)
	}

	return result

}
