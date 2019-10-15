package builder

import (
	"github.com/tucnak/climax"

	"github.com/pieterclaerhout/go-james/internal"
)

// BuildCmd defines the build command
var BuildCmd = climax.Command{
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
		{
			Name:     "output",
			Short:    "o",
			Usage:    `--ouput=<output-path>`,
			Help:     `The path where the executable should be stored`,
			Variable: true,
		},
	},
	Handle: func(ctx climax.Context) int {

		output, _ := ctx.Get("output")
		verbose := ctx.Is("verbose")

		executor := internal.NewExecutor("")
		return executor.DoBuild(output, verbose)

	},
}
