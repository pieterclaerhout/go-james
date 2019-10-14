package runner

import (
	"github.com/tucnak/climax"

	"github.com/pieterclaerhout/go-james/internal"
)

// RunCmd implements the run command
var RunCmd = climax.Command{
	Name:  "run",
	Brief: "Run a binary or example of the local package",
	Help:  "Run a binary or example of the local package",
	Handle: func(ctx climax.Context) int {
		executor := internal.NewExecutor("")
		return executor.DoRun(ctx.Args)
	},
}
