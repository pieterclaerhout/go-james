package cmdrun

import (
	"github.com/tucnak/climax"

	"github.com/pieterclaerhout/go-james"
	"github.com/pieterclaerhout/go-log"
)

var Cmd = climax.Command{
	Name:  "run",
	Brief: "Run a binary or example of the local package",
	Help:  "Run a binary or example of the local package",
	Handle: func(ctx climax.Context) int {

		project := james.NewProject("")
		if err := project.DoRun(ctx.Args); err != nil {
			if log.DebugMode {
				log.Error(err)
			}
			return 1
		}

		return 0

	},
}
