package creator

import (
	"github.com/pieterclaerhout/go-james/internal"
	"github.com/pieterclaerhout/go-james/internal/creator"
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
		{
			Name:     "copyright",
			Short:    "",
			Usage:    `--copyright=<copyright>`,
			Help:     `The copyright of the project`,
			Variable: true,
		},
		{
			Name:     "overwrite",
			Short:    "",
			Usage:    `--overwrite`,
			Help:     `Overwrite the destination path if it exists already`,
			Variable: false,
		},
		{
			Name:     "with-git",
			Short:    "",
			Usage:    `--with-git`,
			Help:     `Create a local git repository for this project`,
			Variable: false,
		},
		{
			Name:     "with-docker",
			Short:    "",
			Usage:    `--with-docker`,
			Help:     `Create a sample Dockerfile`,
			Variable: false,
		},
		{
			Name:     "with-github-action",
			Short:    "",
			Usage:    `--with-github-action`,
			Help:     `Create a sample Github Action workflow`,
			Variable: false,
		},
		{
			Name:     "with-gitlab-ci",
			Short:    "",
			Usage:    `--with-gitlab-ci`,
			Help:     `Create a sample Gitlab-CI workflow`,
			Variable: false,
		},
	},
	Handle: func(ctx climax.Context) int {

		path, _ := ctx.Get("path")
		packageName, _ := ctx.Get("package")
		name, _ := ctx.Get("name")
		description, _ := ctx.Get("description")
		copyright, _ := ctx.Get("copyright")
		overwrite := ctx.Is("overwrite")
		withGit := ctx.Is("with-git")
		withDocker := ctx.Is("with-docker")
		withGithubAction := ctx.Is("with-github-action")
		withGitlabCI := ctx.Is("with-gitlab-ci")

		tool := creator.Creator{
			Mode:             creator.NewProject,
			Path:             path,
			Package:          packageName,
			Name:             name,
			Description:      description,
			Copyright:        copyright,
			Overwrite:        overwrite,
			WithGit:          withGit,
			WithDocker:       withDocker,
			WithGithubAction: withGithubAction,
			WithGitlabCI:     withGitlabCI,
		}

		executor := internal.NewExecutor("")
		return executor.RunTool(tool, false)

	},
}

// InitCmd defines the init command
var InitCmd = climax.Command{
	Name:  "init",
	Brief: "Create a new Go app or library in an existing directory",
	Help:  "Create a new Go app or library in an existing directory",
	Flags: []climax.Flag{
		{
			Name:     "with-git",
			Short:    "",
			Usage:    `--with-git`,
			Help:     `Create a local git repository for this project`,
			Variable: false,
		},
		{
			Name:     "with-docker",
			Short:    "",
			Usage:    `--with-docker`,
			Help:     `Create a sample Dockerfile`,
			Variable: false,
		},
		{
			Name:     "with-github-action",
			Short:    "",
			Usage:    `--with-github-action`,
			Help:     `Create a sample Github Action workflow`,
			Variable: false,
		},
	},
	Handle: func(ctx climax.Context) int {

		withGit := ctx.Is("with-git")
		withDocker := ctx.Is("with-docker")
		withGithubAction := ctx.Is("with-github-action")
		withGitlabCI := ctx.Is("with-gitlab-ci")

		tool := creator.Creator{
			Mode:             creator.InitProject,
			WithGit:          withGit,
			WithDocker:       withDocker,
			WithGithubAction: withGithubAction,
			WithGitlabCI:     withGitlabCI,
		}

		executor := internal.NewExecutor("")
		return executor.RunTool(tool, false)

	},
}
