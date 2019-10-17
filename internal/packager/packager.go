package packager

import (
	"os"
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal/builder"
	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

// Packager implements the "package" command
type Packager struct {
	common.CommandRunner
	common.Compressor
	common.FileSystem
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

		buildOutputPath := packager.buildOutputPathForDistribution(cfg, distribution)

		b := builder.Builder{
			OutputPath: buildOutputPath,
			Verbose:    packager.Verbose,
		}
		if err := b.Execute(project, cfg); err != nil {
			return err
		}

		archivePath := packager.archiveOutputPathForDistribution(cfg, distribution)

		filesToArchive := []string{buildOutputPath}

		if cfg.Package.IncludeReadme {
			readmePath := project.RelPath("README.md")
			if packager.FileExists(readmePath) {
				filesToArchive = append(filesToArchive, readmePath)
			}
		}

		if filepath.Ext(archivePath) == ".tgz" {
			if err := packager.CreateTgz(filesToArchive, archivePath); err != nil {
				return err
			}
		}

		if filepath.Ext(archivePath) == ".zip" {
			if err := packager.CreateZip(filesToArchive, archivePath); err != nil {
				return err
			}
		}

		if err := os.RemoveAll(filepath.Dir(buildOutputPath)); err != nil {
			return err
		}

	}

	return nil

}

// RequiresBuild indicates if a build is required before running the command
func (packager Packager) RequiresBuild() bool {
	return false
}

func (packager Packager) buildOutputPathForDistribution(cfg config.Config, d distribution) string {
	path := filepath.Join(cfg.Build.OutputPath, d.DirName(), cfg.Project.Name)
	if d.GOOS == "windows" {
		path += ".exe"
	}
	return path
}

func (packager Packager) archiveOutputPathForDistribution(cfg config.Config, d distribution) string {
	path := filepath.Join(cfg.Build.OutputPath, cfg.Project.Name+"_"+d.DirName()+"_v"+cfg.Project.Version)
	if d.GOOS == "windows" {
		path += ".zip"
	} else {
		path += ".tgz"
	}
	return path
}
