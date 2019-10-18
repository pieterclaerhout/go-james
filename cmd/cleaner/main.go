package cleaner

import (
	"github.com/pieterclaerhout/go-james/internal"
	"github.com/pieterclaerhout/go-james/internal/cleaner"
	"github.com/tucnak/climax"
)

// CleanCmd defines the clean command
var CleanCmd = climax.Command{
	Name:  "clean",
	Brief: "Removes the build artifacts",
	Help:  "Removes the build artifacts",
	Handle: func(ctx climax.Context) int {

		tool := cleaner.Cleaner{}

		executor := internal.NewExecutor("")
		return executor.RunTool(tool, true)

	},
}
