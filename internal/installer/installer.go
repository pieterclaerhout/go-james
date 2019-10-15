package installer

import (
	"os"
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal/builder"
	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
	"github.com/pieterclaerhout/go-log"
)

// Installer implements the "install" command
type Installer struct {
	Verbose bool
}

// Execute executes the command
func (installer Installer) Execute(project common.Project, cfg config.Config) error {

	dstPath := filepath.Join(os.Getenv("GOPATH"), "bin", filepath.Base(cfg.Build.OutputPath))

	log.Info("Creating:", dstPath)

	b := builder.Builder{
		OutputPath: dstPath,
		Verbose:    installer.Verbose,
	}
	return b.Execute(project, cfg)

}

// RequiresBuild indicates if a build is required before running the command
func (installer Installer) RequiresBuild() bool {
	return false
}
