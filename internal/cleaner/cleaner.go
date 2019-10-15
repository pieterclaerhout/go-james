package cleaner

import (
	"os"
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
	"github.com/pieterclaerhout/go-log"
)

// Cleaner implements the "clean" command
type Cleaner struct {
	common.FileSystem
}

// Execute executes the command
func (cleaner Cleaner) Execute(project common.Project, cfg config.Config) error {

	buildPath := project.RelPath(filepath.Dir(cfg.Build.OutputPath))

	if cleaner.DirExists(buildPath) {
		log.Info("Removing:", buildPath)
		if err := os.RemoveAll(buildPath); err != nil {
			return err
		}
	}

	return nil

}
