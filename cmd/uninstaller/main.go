package uninstaller

import (
	"github.com/pieterclaerhout/go-james/internal"
	"github.com/pieterclaerhout/go-james/internal/uninstaller"
	"github.com/tucnak/climax"
)

// UninstallCmd implements the uninstall command
var UninstallCmd = climax.Command{
	Name:  "uninstall",
	Brief: "Removes the executable from $GOPATh/bin",
	Help:  "Removes the executable from $GOPATh/bin",
	Handle: func(_ climax.Context) int {

		tool := uninstaller.Uninstaller{}

		executor := internal.NewExecutor("")
		return executor.RunTool(tool, true)

	},
}
