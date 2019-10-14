package cmdtest

import (
	"github.com/tucnak/climax"

	"github.com/pieterclaerhout/go-james"
	"github.com/pieterclaerhout/go-log"
)

var Cmd = climax.Command{
	Name:  "test",
	Brief: "Run the tests",
	Help:  "Run the tests",
	Handle: func(ctx climax.Context) (exitcode int) {

		project := james.NewProject("")
		if err := project.DoTest(); err != nil {
			if log.DebugMode {
				log.Error(err)
			}
			return 1
		}

		return 0

	},
}
