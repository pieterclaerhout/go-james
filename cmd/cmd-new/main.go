package cmdnew

import (
	"github.com/tucnak/climax"

	"github.com/pieterclaerhout/go-james"
	"github.com/pieterclaerhout/go-log"
)

var Cmd = climax.Command{
	Name:  "new",
	Brief: "Create a new Go app or library",
	Help:  "Create a new Go app or library",
	Handle: func(ctx climax.Context) int {

		project := james.NewProject("")
		if err := project.DoNew(); err != nil {
			if log.DebugMode {
				log.Error(err)
			}
			return 1
		}

		return 0

	},
}
