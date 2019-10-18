package creator

import (
	"os"
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal"
	"github.com/tucnak/climax"
)

// NewCmd defines the new command
var NewCmd = climax.Command{
	Name:  "new",
	Brief: "Create a new Go app or library",
	Help:  "Create a new Go app or library",
	Flags: []climax.Flag{
		{
			Name:     "path",
			Short:    "",
			Usage:    `--path=<target-path>`,
			Help:     `The path where the command should be created`,
			Variable: true,
		},
		{
			Name:     "package",
			Short:    "",
			Usage:    `--package=<package>`,
			Help:     `The package for the project`,
			Variable: true,
		},
		{
			Name:     "name",
			Short:    "",
			Usage:    `--name=<name>`,
			Help:     `The name of the project`,
			Variable: true,
		},
		{
			Name:     "description",
			Short:    "",
			Usage:    `--description=<description>`,
			Help:     `The description of the project`,
			Variable: true,
		},
	},
	Handle: func(ctx climax.Context) int {

		path, _ := ctx.Get("path")
		packageName, _ := ctx.Get("package")
		name, _ := ctx.Get("name")
		description, _ := ctx.Get("description")

		if path == "" && name != "" {
			wd, _ := os.Getwd()
			path = filepath.Join(wd, name)
		}

		executor := internal.NewExecutor("")
		return executor.DoNew(path, packageName, name, description)

	},
}

// InitCmd defines the init command
var InitCmd = climax.Command{
	Name:  "init",
	Brief: "Create a new Go app or library in an existing directory",
	Help:  "Create a new Go app or library in an existing directory",
	Handle: func(ctx climax.Context) int {
		executor := internal.NewExecutor("")
		return executor.DoInit()
	},
}
