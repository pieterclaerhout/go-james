package creator

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

// Mode is used to define the mode in which we run (init or new)
type Mode int

const (
	// InitProject initializes a project in an existing directory
	InitProject Mode = iota

	// NewProject creates a new project in the given path
	NewProject
)

// String translates the creator mode to a string
func (c Mode) String() string {
	return [...]string{"Init Project", "New Project"}[c]
}

// Creator implements the "init" and "new" commands
type Creator struct {
	common.CommandRunner
	common.FileSystem
	common.Template
	Mode        Mode
	Path        string
	Package     string
	Name        string
	Description string
}

// Execute executes the command
func (creator Creator) Execute(project common.Project, cfg config.Config) error {

	if creator.Mode == InitProject {
		creator.Path, _ = os.Getwd()
		creator.Package = filepath.Base(creator.Path)
	}

	if creator.Path == "" {
		return errors.New("Path not specified")
	}

	if creator.Package == "" {
		return errors.New("Package not specified")
	}

	project = common.Project{
		Path: creator.Path,
	}

	type creationStep func(project common.Project, cfg config.Config) error

	var steps = []creationStep{
		creator.createConfig,
		creator.createTasks,
		creator.createLicense,
		creator.createGitIgnore,
		creator.createReadme,
		creator.createSourceFiles,
		creator.createGoMod,
	}

	for _, step := range steps {

		var err error

		if err = step(project, cfg); err != nil {
			return err
		}

		if cfg, err = config.NewConfigFromDir(project.Path); err != nil {
			return err
		}

	}

	return nil

}

// RequiresBuild indicates if a build is required before running the command
func (creator Creator) RequiresBuild() bool {
	return false
}

func (creator Creator) createConfig(project common.Project, cfg config.Config) error {

	configPath := project.RelPath(config.FileName)

	if creator.Name == "" {
		creator.Name = filepath.Base(creator.Path)
	}

	cfg = config.Config{
		Project: config.ProjectConfig{
			Name:        creator.Name,
			Description: creator.Description,
			Package:     creator.Package,
			MainPackage: creator.Package + "/cmd/" + creator.Name,
		},
		Build: config.BuildConfig{
			OutputPath: filepath.Join("build", creator.Name),
			LDFlags:    []string{"-s", "-w"},
			ExtraArgs:  []string{"-trimpath"},
		},
	}

	return creator.WriteJSONFileIfNotExists(configPath, cfg)

}

func (creator Creator) createTasks(project common.Project, cfg config.Config) error {

	tasks := newVisualStudioCodeTaskList()

	tasksPath := project.RelPath(visualStudioDirName, visualStudioCodeTasksFileName)
	return creator.WriteJSONFileIfNotExists(tasksPath, tasks)

}

func (creator Creator) createLicense(project common.Project, cfg config.Config) error {

	path := project.RelPath(licenseFileName)
	return creator.WriteTextFileIfNotExists(path, apacheLicense)

}

func (creator Creator) createGitIgnore(project common.Project, cfg config.Config) error {

	gitIgnore := newGitIgnore(cfg)

	path := project.RelPath(gitIgnoreFileName)
	return creator.WriteTextFileIfNotExists(path, gitIgnore.string())

}

func (creator Creator) createReadme(project common.Project, cfg config.Config) error {

	readme := newReadme(cfg)

	path := project.RelPath(readmeFileName)
	return creator.WriteTextFileIfNotExists(path, readme.markdownString())

}

func (creator Creator) createGoMod(project common.Project, cfg config.Config) error {

	goModPath := project.RelPath(goModFileName)
	if creator.FileExists(goModPath) {
		return nil
	}

	cmd := []string{"go", "mod", "init", cfg.Project.Package}
	return creator.RunToStdout(cmd, project.Path, map[string]string{
		"GO111MODULE": "on",
	})

}

func (creator Creator) createSourceFiles(project common.Project, cfg config.Config) error {

	mainLibPath := project.RelPath("library.go")
	if err := creator.WriteTextTemplateIfNotExists(mainLibPath, mainLibTemplate, cfg); err != nil {
		return err
	}

	mainCmdPath := project.RelPath("cmd", filepath.Base(cfg.Project.Package), "main.go")
	if err := creator.WriteTextTemplateIfNotExists(mainCmdPath, mainCmdTemplate, cfg); err != nil {
		return err
	}

	return nil

}
