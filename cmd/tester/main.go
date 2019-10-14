package tester

import (
	"github.com/tucnak/climax"

	"github.com/pieterclaerhout/go-james/internal"
)

// TestCmd implements the test command
var TestCmd = climax.Command{
	Name:  "test",
	Brief: "Run the tests",
	Help:  "Run the tests",
	Handle: func(ctx climax.Context) (exitcode int) {
		executor := internal.NewExecutor("")
		return executor.DoTest()
	},
}
