package cmdinit

import (
	"github.com/tucnak/climax"

	"github.com/pieterclaerhout/go-james/internal"
	"github.com/pieterclaerhout/go-log"
)

var Cmd = climax.Command{
	Name:  "init",
	Brief: "Create a new Go app or library in an existing directory",
	Help:  "Create a new Go app or library in an existing directory",
	Handle: func(ctx climax.Context) int {

		project := internal.NewProject("")
		if err := project.DoInit(); err != nil {
			if log.DebugMode {
				log.Error(err)
			}
			return 1
		}

		return 0

	},
}
