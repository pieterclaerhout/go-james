package updater

import (
	"github.com/pieterclaerhout/go-james/internal"
	"github.com/pieterclaerhout/go-james/internal/updater"
	"github.com/pieterclaerhout/go-james/versioninfo"
	"github.com/tucnak/climax"
)

// UpdateCmd defines the update command
var UpdateCmd = climax.Command{
	Name:  "update",
	Brief: "Updates " + versioninfo.ProjectName + " to the latest available release",
	Help:  "Updates " + versioninfo.ProjectName + " to the latest available release",
	Handle: func(ctx climax.Context) int {

		tool := updater.Updater{}

		executor := internal.NewExecutor("")
		return executor.RunTool(tool, false)

	},
}
