package cleaner

import (
	"github.com/tucnak/climax"

	"github.com/pieterclaerhout/go-james/internal"
)

// CleanCmd defines the clean command
var CleanCmd = climax.Command{
	Name:  "clean",
	Brief: "Removes the build artifacts",
	Help:  "Removes the build artifacts",
	Handle: func(ctx climax.Context) int {
		executor := internal.NewExecutor("")
		return executor.DoClean()
	},
}
