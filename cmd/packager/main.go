package packager

import (
	"github.com/pieterclaerhout/go-james/internal"
	"github.com/tucnak/climax"
)

// PackageCmd defines the package command
var PackageCmd = climax.Command{
	Name:  "package",
	Brief: "Build the main executable of your project for windows / linux / mac",
	Help:  "Build the main executable of your project for windows / linux / mac",
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
		return executor.DoPackage(verbose)

	},
}
