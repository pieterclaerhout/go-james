package staticcheck

import (
	"github.com/pieterclaerhout/go-james/internal"
	"github.com/pieterclaerhout/go-james/internal/staticchecker"
	"github.com/tucnak/climax"
)

// StaticcheckCmd defines the build command
var StaticcheckCmd = climax.Command{
	Name:  "staticcheck",
	Brief: "Perform a static analysis of the code using staticcheck",
	Help:  "Perform a static analysis of the code using staticcheck",
	Flags: []climax.Flag{},
	Handle: func(ctx climax.Context) int {

		tool := staticchecker.StaticChecker{}

		executor := internal.NewExecutor("")
		return executor.RunTool(tool, true)

	},
}
