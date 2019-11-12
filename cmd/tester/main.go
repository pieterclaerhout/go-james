package tester

import (
	"github.com/pieterclaerhout/go-james/internal"
	"github.com/pieterclaerhout/go-james/internal/tester"
	"github.com/tucnak/climax"
)

// TestCmd implements the test command
var TestCmd = climax.Command{
	Name:  "test",
	Brief: "Run the tests",
	Help:  "Run the tests",
	Handle: func(_ climax.Context) (exitcode int) {

		tool := tester.Tester{}

		executor := internal.NewExecutor("")
		return executor.RunTool(tool, true)

	},
}
