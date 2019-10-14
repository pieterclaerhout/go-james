package internal

import "github.com/pieterclaerhout/go-james/internal/config"

// Subcommand defines the interface a subcommand need to implement
type Subcommand interface {
	Execute(project Project, cfg config.Config) error // Execute executes the subcommand with the given project and config
}
