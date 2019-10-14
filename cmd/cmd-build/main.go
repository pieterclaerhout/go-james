package cmdbuild

import (
	"github.com/tucnak/climax"

	"github.com/pieterclaerhout/go-james/internal"
)

var Cmd = climax.Command{
	Name:  "build",
	Brief: "Compile the current package",
	Help:  "Compile the current package",
	Flags: []climax.Flag{
		{
			Name:     "verbose",
			Short:    "v",
			Usage:    `--verbose`,
			Help:     `print the names of packages as they are compiled.`,
			Variable: false,
		},
	},
	Handle: func(ctx climax.Context) int {

		verbose := ctx.Is("verbose")

		executor := internal.NewExecutor("")
		return executor.DoBuild(verbose)

	},
}
