package cmdbuild

import (
	"fmt"

	"github.com/pieterclaerhout/go-james/version"
	"github.com/tucnak/climax"
)

var Cmd = climax.Command{
	Name:  "version",
	Brief: "Print version info and exit",
	Help:  "Print version info and exit",
	Handle: func(ctx climax.Context) int {
		fmt.Println(version.AppName + " " + version.Revision + " (" + version.Branch + ")")
		return 0

	},
}
