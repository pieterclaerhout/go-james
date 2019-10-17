package installer

import (
	"github.com/pieterclaerhout/go-james/internal/builder"
	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

// Installer implements the "install" command
type Installer struct {
	common.FileSystem
	common.Golang

	Verbose bool
}

// Execute executes the command
func (installer Installer) Execute(project common.Project, cfg config.Config) error {

	dstPath := installer.GoBin(cfg.Project.Name)

	installer.LogPathCreation(dstPath)

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
