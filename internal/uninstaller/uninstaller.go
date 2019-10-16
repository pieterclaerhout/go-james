package uninstaller

import (
	"os"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
	"github.com/pieterclaerhout/go-log"
)

// Uninstaller implements the "uninstall" command
type Uninstaller struct {
	common.FileSystem
	common.Golang
}

// Execute executes the command
func (uninstaller Uninstaller) Execute(project common.Project, cfg config.Config) error {

	dstPath := uninstaller.GoBin(cfg.Project.Name)

	if uninstaller.FileExists(dstPath) {
		log.Info("Deleting:", dstPath)
		if err := os.Remove(dstPath); err != nil {
			return nil
		}
	}

	return nil

}

// RequiresBuild indicates if a build is required before running the command
func (uninstaller Uninstaller) RequiresBuild() bool {
	return false
}
