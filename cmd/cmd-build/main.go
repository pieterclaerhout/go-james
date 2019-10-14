package cmdbuild

import (
	"github.com/tucnak/climax"

	"github.com/pieterclaerhout/go-james/internal"
	"github.com/pieterclaerhout/go-log"
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

		project := internal.NewProject("")
		if err := project.DoBuild(verbose); err != nil {
			if log.DebugMode {
				log.Error(err)
			}
			return 1
		}

		return 0

	},
}
