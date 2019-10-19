package builder

import (
	"github.com/pieterclaerhout/go-james/internal"
	"github.com/pieterclaerhout/go-james/internal/builder"
	"github.com/tucnak/climax"
)

// BuildCmd defines the build command
var BuildCmd = climax.Command{
	Name:  "build",
	Brief: "Build the main executable of your project",
	Help:  "Build the main executable of your project",
	Flags: []climax.Flag{
		{
			Name:     "output",
			Short:    "o",
			Usage:    `--ouput=<output-path>`,
			Help:     `The path where the executable should be stored`,
			Variable: true,
		},
		{
			Name:     "goos",
			Short:    "",
			Usage:    `--goos=<os>`,
			Help:     `The goos to build for`,
			Variable: true,
		},
		{
			Name:     "goarch",
			Short:    "",
			Usage:    `--goarch=<arch>`,
			Help:     `The goarch to build for`,
			Variable: true,
		},
		{
			Name:     "verbose",
			Short:    "v",
			Usage:    `--verbose`,
			Help:     `print the names of packages as they are compiled.`,
			Variable: false,
		},
	},
	Handle: func(ctx climax.Context) int {

		output, _ := ctx.Get("output")
		goos, _ := ctx.Get("goos")
		goarch, _ := ctx.Get("goarch")
		verbose := ctx.Is("verbose")

		tool := builder.Builder{
			OutputPath: output,
			GOOS:       goos,
			GOARCH:     goarch,
			Verbose:    verbose,
		}

		executor := internal.NewExecutor("")
		return executor.RunTool(tool, true)

	},
}
