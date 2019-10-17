package debugger

import (
	"github.com/pieterclaerhout/go-james/internal"
	"github.com/tucnak/climax"
)

// DebugCmd implements the debug command
var DebugCmd = climax.Command{
	Name:  "debug",
	Brief: "Debug a binary or example using delve",
	Help:  "Debug a binary or example using delve",
	Handle: func(ctx climax.Context) int {
		executor := internal.NewExecutor("")
		return executor.DoDebug(ctx.Args)
	},
}
