package packager

import (
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal/builder"
	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

// Packager implements the "package" command
type Packager struct {
	common.CommandRunner
	common.Logging

	Verbose bool
}

// Execute executes the command
func (packager Packager) Execute(project common.Project, cfg config.Config) error {

	allDistributions, err := packager.allPossibleDistributions()
	if err != nil {
		return err
	}

	for _, distribution := range allDistributions {

		if !distribution.ShouldBuild() {
			continue
		}

		dstPath := packager.outputPathForDistribution(cfg, distribution)

		packager.LogPathCreation(dstPath)

		b := builder.Builder{
			OutputPath: dstPath,
			Verbose:    packager.Verbose,
		}
		if err := b.Execute(project, cfg); err != nil {
			return err
		}

	}

	return nil

}

// RequiresBuild indicates if a build is required before running the command
func (packager Packager) RequiresBuild() bool {
	return false
}

func (packager Packager) outputPathForDistribution(cfg config.Config, d distribution) string {
	path := filepath.Join(cfg.Build.OutputPath, d.DirName(), cfg.Project.Name)
	if d.GOOS == "windows" {
		path += ".exe"
	}
	return path
}
