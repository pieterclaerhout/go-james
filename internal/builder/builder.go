package builder

import (
	"path/filepath"
	"runtime"
	"strings"

	"github.com/kballard/go-shellquote"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

// Builder implements the "build" command
type Builder struct {
	common.CommandRunner
	OutputPath string
	GOOS       string
	GOARCH     string
	Verbose    bool
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

	outputPath := builder.outputPath(cfg)
	if outputPath != "" {
		buildCmd = append(buildCmd, "-o", outputPath)
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
	return builder.RunToStdout(buildCmd, project.Path, map[string]string{
		"GO111MODULE": "on",
		"GOOS":        builder.GOOS,
		"GOARCH":      builder.GOARCH,
	})

}

// RequiresBuild indicates if a build is required before running the command
func (builder Builder) RequiresBuild() bool {
	return false
}

func (builder Builder) determineRevision(project common.Project) string {

	cmdLine := []string{"git", "rev-parse", "--short", "HEAD"}

	result, _ := builder.RunReturnOutput(cmdLine, project.Path, map[string]string{})
	return strings.TrimSpace(result)

}

func (builder Builder) determineBranch(project common.Project) string {

	cmdLine := []string{"git", "rev-parse", "--abbrev-ref", "HEAD"}

	result, _ := builder.RunReturnOutput(cmdLine, project.Path, map[string]string{})
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

func (builder Builder) outputPath(cfg config.Config) string {

	outputPath := builder.OutputPath
	if outputPath == "" {
		outputPath = cfg.Build.OutputPath
	}

	if outputPath != "" {
		if runtime.GOOS == "win" {
			if filepath.Ext(outputPath) != ".exe" {
				outputPath = outputPath + ".exe"
			}
		}
	}

	return outputPath

}
