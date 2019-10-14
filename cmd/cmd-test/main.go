package cmdtest

import (
	"github.com/tucnak/climax"

	"github.com/pieterclaerhout/go-james/internal"
)

var Cmd = climax.Command{
	Name:  "test",
	Brief: "Run the tests",
	Help:  "Run the tests",
	Handle: func(ctx climax.Context) (exitcode int) {
		executor := internal.NewExecutor("")
		return executor.DoTest()
	},
}
