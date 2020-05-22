package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pieterclaerhout/go-james/cmd/dockerimage"
	"github.com/pieterclaerhout/go-james/cmd/staticcheck"

	"github.com/pieterclaerhout/go-james/cmd/builder"
	"github.com/pieterclaerhout/go-james/cmd/cleaner"
	"github.com/pieterclaerhout/go-james/cmd/creator"
	"github.com/pieterclaerhout/go-james/cmd/debugger"
	"github.com/pieterclaerhout/go-james/cmd/installer"
	"github.com/pieterclaerhout/go-james/cmd/packager"
	"github.com/pieterclaerhout/go-james/cmd/runner"
	"github.com/pieterclaerhout/go-james/cmd/tester"
	"github.com/pieterclaerhout/go-james/cmd/uninstaller"
	"github.com/pieterclaerhout/go-james/cmd/version"
	"github.com/pieterclaerhout/go-james/internal"
	rawrunner "github.com/pieterclaerhout/go-james/internal/runner"
	"github.com/pieterclaerhout/go-james/versioninfo"
	"github.com/pieterclaerhout/go-log"
	"github.com/tucnak/climax"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	log.PrintColors = true
	log.DebugMode = false

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

		tool := rawrunner.Runner{
			Args: args,
		}

		executor := internal.NewExecutor("")
		result = executor.RunTool(tool, true)

	} else {

		exePath, _ := os.Executable()
		exeName := filepath.Base(exePath)

		app := climax.New(exeName)
		app.Name = versioninfo.ProjectName
		app.Brief = versioninfo.ProjectDescription
		app.Version = versioninfo.Revision

		app.AddCommand(builder.BuildCmd)
		app.AddCommand(cleaner.CleanCmd)
		app.AddCommand(creator.InitCmd)
		app.AddCommand(creator.NewCmd)
		app.AddCommand(debugger.DebugCmd)
		app.AddCommand(dockerimage.DockerImageCmd)
		app.AddCommand(installer.InstallCmd)
		app.AddCommand(packager.PackageCmd)
		app.AddCommand(runner.RunCmd)
		app.AddCommand(staticcheck.StaticcheckCmd)
		app.AddCommand(tester.TestCmd)
		app.AddCommand(uninstaller.UninstallCmd)
		app.AddCommand(version.VersionCmd)

		result = app.Run()

	}

	os.Exit(result)

}
