package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/tucnak/climax"

	"github.com/pieterclaerhout/go-james/cmd/builder"
	"github.com/pieterclaerhout/go-james/cmd/cleaner"
	"github.com/pieterclaerhout/go-james/cmd/creator"
	"github.com/pieterclaerhout/go-james/cmd/installer"
	"github.com/pieterclaerhout/go-james/cmd/runner"
	"github.com/pieterclaerhout/go-james/cmd/tester"
	"github.com/pieterclaerhout/go-james/cmd/version"
	"github.com/pieterclaerhout/go-james/internal"
	"github.com/pieterclaerhout/go-james/versioninfo"
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
		app.Name = versioninfo.AppName
		app.Brief = "James is your butler and helps you to create, build, test and run your Go projects"
		app.Version = versioninfo.Revision

		app.AddCommand(builder.BuildCmd)
		app.AddCommand(cleaner.CleanCmd)
		app.AddCommand(creator.InitCmd)
		app.AddCommand(creator.NewCmd)
		app.AddCommand(installer.InstallCmd)
		app.AddCommand(runner.RunCmd)
		app.AddCommand(tester.TestCmd)
		app.AddCommand(version.VersionCmd)

		result = app.Run()

	}

	os.Exit(result)

}
