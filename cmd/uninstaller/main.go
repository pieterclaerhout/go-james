package uninstaller

import (
	"github.com/pieterclaerhout/go-james/internal"
	"github.com/tucnak/climax"
)

// UninstallCmd implements the uninstall command
var UninstallCmd = climax.Command{
	Name:  "uninstall",
	Brief: "Removes the executable from $GOPATh/bin",
	Help:  "Removes the executable from $GOPATh/bin",
	Handle: func(ctx climax.Context) int {
		executor := internal.NewExecutor("")
		return executor.DoUninstall()
	},
}
