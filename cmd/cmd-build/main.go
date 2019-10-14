package cmdbuild

import (
	"github.com/tucnak/climax"

	"github.com/pieterclaerhout/go-james"
	"github.com/pieterclaerhout/go-log"
)

var Cmd = climax.Command{
	Name:  "build",
	Brief: "Compile the current package",
	Help:  "Compile the current package",
	Handle: func(ctx climax.Context) int {

		project := james.NewProject("")
		if err := project.DoBuild(); err != nil {
			if log.DebugMode {
				log.Error(err)
			}
			return 1
		}

		return 0

	},
}
