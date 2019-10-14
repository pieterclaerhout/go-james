package version

import (
	"fmt"

	"github.com/pieterclaerhout/go-james/versioninfo"
	"github.com/tucnak/climax"
)

// VersionCmd defines the version command
var VersionCmd = climax.Command{
	Name:  "version",
	Brief: "Print version info and exit",
	Help:  "Print version info and exit",
	Handle: func(ctx climax.Context) int {
		fmt.Println(versioninfo.AppName + " " + versioninfo.Revision + " (" + versioninfo.Branch + ")")
		return 0
	},
}
