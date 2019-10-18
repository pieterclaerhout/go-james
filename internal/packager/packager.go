package packager

import (
	"context"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pieterclaerhout/go-james/internal/builder"
	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
	"github.com/pieterclaerhout/go-waitgroup"
)

// Packager implements the "package" command
type Packager struct {
	common.CommandRunner
	common.Compressor
	common.FileSystem
	common.Logging
	common.Timer

	Concurrency int
	Verbose     bool
}

// Execute executes the command
func (packager Packager) Execute(project common.Project, cfg config.Config) error {

	packager.StartTimer()
	defer packager.PrintElapsed("Package time:")

	if packager.Concurrency < 1 {
		packager.Concurrency = runtime.NumCPU()
	}

	packager.LogInfo("Using a concurrency of:", packager.Concurrency)

	distributions, err := packager.allDistributionsToBuild()
	if err != nil {
		return err
	}

	ctx := context.Background()
	wg, ctx := waitgroup.NewErrorGroup(ctx, packager.Concurrency)
	if err != nil {
		return err
	}

	for _, distribution := range distributions {

		localDistribution := distribution

		wg.Add(func() error {
			return packager.buildDistribution(project, cfg, localDistribution)
		})

	}

	return wg.Wait()

}

// RequiresBuild indicates if a build is required before running the command
func (packager Packager) RequiresBuild() bool {
	return false
}

func (packager Packager) buildDistribution(project common.Project, cfg config.Config, d distribution) error {

	buildOutputPath := packager.buildOutputPathForDistribution(cfg, d)

	b := builder.Builder{
		OutputPath: buildOutputPath,
		Verbose:    packager.Verbose,
	}
	if err := b.Execute(project, cfg); err != nil {
		return err
	}

	archivePath := packager.archiveOutputPathForDistribution(cfg, d)

	compressor := packager.CreateTgz(archivePath)
	if filepath.Ext(archivePath) == ".zip" {
		compressor = packager.CreateZip(archivePath)
	}

	compressor.AddFile("", buildOutputPath)

	readmePath := project.RelPath("README.md")
	if cfg.Package.IncludeReadme && packager.FileExists(readmePath) {
		compressor.AddFile("", readmePath)
	}

	if err := compressor.Close(); err != nil {
		return err
	}

	if err := os.RemoveAll(filepath.Dir(buildOutputPath)); err != nil {
		return err
	}

	return nil

}

func (packager Packager) buildOutputPathForDistribution(cfg config.Config, d distribution) string {
	path := filepath.Join(cfg.Build.OutputPath, d.DirName(), cfg.Project.Name)
	if d.GOOS == "windows" {
		path += ".exe"
	}
	return path
}

func (packager Packager) archiveOutputPathForDistribution(cfg config.Config, d distribution) string {
	path := filepath.Join(cfg.Build.OutputPath, cfg.Project.Name+"_"+d.DirName())
	if d.GOOS == "windows" {
		path += ".exe.zip"
	} else {
		path += ".tgz"
	}
	return path
}
