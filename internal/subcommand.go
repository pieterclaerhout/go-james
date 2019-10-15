package internal

import (
	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

// Subcommand defines the interface a subcommand need to implement
type Subcommand interface {
	Execute(project common.Project, cfg config.Config) error // Execute executes the subcommand with the given project and config
	RequiresBuild() bool                                     // Return true if a build is required before running the command
}
