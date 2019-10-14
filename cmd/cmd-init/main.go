package cmdinit

import (
	"github.com/tucnak/climax"
)

var Cmd = climax.Command{
	Name:  "init",
	Brief: "Create a new Go app or library in an existing directory",
	Help:  "Create a new Go app or library in an existing directory",
	Handle: func(ctx climax.Context) int {

		// project := internal.NetExecutor("")
		// if err := project.DoInit(); err != nil {
		// 	if log.DebugMode {
		// 		log.Error(err)
		// 	}
		// 	return 1
		// }

		return 0

	},
}
