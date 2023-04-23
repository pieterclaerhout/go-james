package cleaner

import (
	"os"
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
	"github.com/pieterclaerhout/go-log"
	"github.com/pkg/errors"
)

// Cleaner implements the "clean" command
type Cleaner struct {
	common.FileSystem
}

// Execute executes the command
func (cleaner Cleaner) Execute(project common.Project, cfg config.Config) error {

	if cfg.Build.OutputPath == "" {
		return errors.New("Config setting build.output_path shouldn't be empty")
	}

	buildPath := project.RelPath(filepath.Dir(cfg.Build.OutputPath))
	rootPath := project.RelPath()

	if cleaner.DirExists(buildPath) && buildPath != rootPath {
		log.Info("Removing:", buildPath)
		if err := os.RemoveAll(buildPath); err != nil {
			return err
		}
	}

	return nil

}

// RequiresBuild indicates if a build is required before running the command
func (cleaner Cleaner) RequiresBuild() bool {
	return false
}
