package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/tucnak/climax"

	"github.com/pieterclaerhout/go-james"
	cmdbuild "github.com/pieterclaerhout/go-james/cmd/cmd-build"
	cmdinit "github.com/pieterclaerhout/go-james/cmd/cmd-init"
	cmdnew "github.com/pieterclaerhout/go-james/cmd/cmd-new"
	cmdrun "github.com/pieterclaerhout/go-james/cmd/cmd-run"
	cmdtest "github.com/pieterclaerhout/go-james/cmd/cmd-test"
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

		project := james.NewProject("")
		if err := project.DoRun(args); err != nil {
			log.Error(err)
			result = 1
		}

	} else {

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

		result = app.Run()

	}

	os.Exit(result)

}
