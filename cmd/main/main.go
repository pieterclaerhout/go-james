package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/tucnak/climax"

	cmdbuild "github.com/pieterclaerhout/go-james/cmd/cmd-build"
	cmdinit "github.com/pieterclaerhout/go-james/cmd/cmd-init"
	cmdnew "github.com/pieterclaerhout/go-james/cmd/cmd-new"
	cmdrun "github.com/pieterclaerhout/go-james/cmd/cmd-run"
	cmdtest "github.com/pieterclaerhout/go-james/cmd/cmd-test"
	cmdversion "github.com/pieterclaerhout/go-james/cmd/cmd-version"
	"github.com/pieterclaerhout/go-james/internal"
	"github.com/pieterclaerhout/go-james/version"
	"github.com/pieterclaerhout/go-log"
)

func main() {

	if log.DebugMode {
		log.PrintTimestamp = true
	}

	var commandName string
	var result int

	if len(os.Args) > 1 {
		commandName = strings.TrimSpace(strings.ToLower(os.Args[1]))
	}

	if commandName == "run" {

		args := []string{}
		if len(os.Args) > 2 {
			args = os.Args[2:]
		}

		executor := internal.NewExecutor("")
		result = executor.DoRun(args)

	} else {

		exePath, _ := os.Executable()
		exeName := filepath.Base(exePath)

		app := climax.New(exeName)
		app.Name = version.AppName
		app.Brief = "James is your butler and helps you to create, build, test and run your Go projects"
		app.Version = version.Revision

		app.AddCommand(cmdbuild.Cmd)
		app.AddCommand(cmdinit.Cmd)
		app.AddCommand(cmdnew.Cmd)
		app.AddCommand(cmdrun.Cmd)
		app.AddCommand(cmdtest.Cmd)
		app.AddCommand(cmdversion.Cmd)

		result = app.Run()

	}

	os.Exit(result)

}
