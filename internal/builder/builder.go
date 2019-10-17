package builder

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/kballard/go-shellquote"
	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
	"github.com/pkg/errors"
)

// Builder implements the "build" command
type Builder struct {
	common.CommandRunner
	common.FileSystem

	OutputPath string
	GOOS       string
	GOARCH     string
	Verbose    bool
}

// Execute executes the command
func (builder Builder) Execute(project common.Project, cfg config.Config) error {

	versionInfo := map[string]string{
		"ProjectName":        cfg.Project.Name,
		"ProjectDescription": cfg.Project.Description,
		"Version":            cfg.Project.Version,
		"Revision":           builder.determineRevision(project),
		"Branch":             builder.determineBranch(project),
	}

	buildCmd := []string{"go", "build"}

	if builder.Verbose {
		buildCmd = append(buildCmd, "-v")
	}

	outputPath, err := builder.outputPath(cfg)
	if err != nil {
		return err
	}
	if outputPath != "" {
		buildCmd = append(buildCmd, "-o", outputPath)
	}

	outputFolder := filepath.Dir(outputPath)
	if builder.DirExists(outputFolder) || builder.FileExists(outputFolder) {
		if err := os.RemoveAll(outputFolder); err != nil {
			return err
		}
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

func (builder Builder) outputPath(cfg config.Config) (string, error) {

	outputPath := builder.OutputPath
	if outputPath == "" {
		if builder.FileExists(cfg.Build.OutputPath) {
			return "", errors.New("build.output_path in config should point to a directory, not a file")
		}
		outputPath = filepath.Join(cfg.Build.OutputPath, cfg.Project.Name)
	}

	if outputPath != "" && runtime.GOOS == "windows" && filepath.Ext(outputPath) != ".exe" {
		outputPath = outputPath + ".exe"
	}

	return outputPath, nil

}
