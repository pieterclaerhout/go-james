package tester

import (
	"github.com/pieterclaerhout/go-james/internal"
	"github.com/tucnak/climax"
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
