package version

import (
	"github.com/pieterclaerhout/go-james/internal"
	"github.com/pieterclaerhout/go-james/internal/version"
	"github.com/tucnak/climax"
)

// VersionCmd defines the version command
var VersionCmd = climax.Command{
	Name:  "version",
	Brief: "Print version info and exit",
	Help:  "Print version info and exit",
	Handle: func(ctx climax.Context) int {

		tool := version.Version{}

		executor := internal.NewExecutor("")
		return executor.RunTool(tool, false)

	},
}
