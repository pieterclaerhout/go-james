package main

import (
	"os"
	"path/filepath"

	"github.com/tucnak/climax"

	"github.com/pieterclaerhout/go-james/cmd/cmd-build"
	"github.com/pieterclaerhout/go-james/cmd/cmd-init"
	"github.com/pieterclaerhout/go-james/cmd/cmd-new"
	"github.com/pieterclaerhout/go-james/cmd/cmd-run"
	"github.com/pieterclaerhout/go-james/cmd/cmd-test"
	"github.com/pieterclaerhout/go-james/version"
	"github.com/pieterclaerhout/go-log"
)

func main() {

	if log.DebugMode {
		log.PrintTimestamp = true
	}

	exePath, _ := os.Executable()
	exeName := filepath.Base(exePath)

	app := climax.New(exeName)
	app.Brief = "James is your butler and helps you to create, build, test and run your Go projects"
	app.Version = version.Revision

	app.AddCommand(cmdbuild.Cmd)
	app.AddCommand(cmdinit.Cmd)
	app.AddCommand(cmdnew.Cmd)
	app.AddCommand(cmdrun.Cmd)
	app.AddCommand(cmdtest.Cmd)

	result := app.Run()
	os.Exit(result)

}
