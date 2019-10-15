package installer

import (
	"github.com/tucnak/climax"

	"github.com/pieterclaerhout/go-james/internal"
)

// InstallCmd implements the install command
var InstallCmd = climax.Command{
	Name:  "install",
	Brief: "Build the executable and install it in $GOPATh/bin",
	Help:  "Build the executable and install it in $GOPATh/bin",
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
		return executor.DoInstall(verbose)

	},
}
