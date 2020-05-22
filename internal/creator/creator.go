package creator

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
	"github.com/pieterclaerhout/go-log"
	"github.com/pkg/errors"
)

// Mode is used to define the mode in which we run (init or new)
type Mode int

const (
	// InitProject initializes a project in an existing directory
	InitProject Mode = iota + 1

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
	common.Logging

	Mode          Mode
	Path          string
	Package       string
	Name          string
	Description   string
	Copyright     string
	Overwrite     bool
	CreateGitRepo bool
}

// Execute executes the command
func (creator Creator) Execute(project common.Project, cfg config.Config) error {

	if creator.Path == "" && creator.Name != "" {
		wd, _ := os.Getwd()
		creator.Path = filepath.Join(wd, creator.Name)
	}

	if creator.Path == "" && creator.Package != "" {
		creator.Path = filepath.Base(creator.Package)
	}

	if creator.Package == "" && creator.Name != "" {
		creator.Package = creator.Name
	}

	if creator.Mode == InitProject {

		creator.Path, _ = os.Getwd()

		if packageName, err := project.Package(); err == nil {
			creator.Package = packageName
		}

		if creator.Package == "" {
			creator.Package = common.PackageNameToShort(creator.Path)
		}

	}

	if creator.Path == "" {
		return errors.New("Path not specified")
	}

	if creator.Package == "" {
		return errors.New("Package not specified")
	}

	if creator.Mode == NewProject && creator.PathExists(creator.Path) {
		if creator.Overwrite {
			log.Warn("!!! Overwriting:", creator.Path, "!!!")
			if err := os.RemoveAll(creator.Path); err != nil {
				return err
			}
		} else {
			return errors.New("The destination path exists already")
		}
	}

	project = common.NewProject(creator.Path, creator.Package)

	log.Info("Creating package:", creator.Package)
	log.Info("Project path:", project.Path)

	type creationStep func(project common.Project, cfg config.Config) error

	var steps = []creationStep{
		creator.createConfig,
		creator.createTasks,
		creator.createLaunchConfig,
		creator.createLicense,
		creator.createGitIgnore,
		creator.createDockerIgnore,
		creator.createDockerFile,
		creator.createReadme,
		creator.createSourceFiles,
		creator.createGoMod,
	}

	if creator.CreateGitRepo {
		steps = append(steps, creator.createGitRepo)
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
			Version:     "1.0.0",
			Description: creator.Description,
			Copyright:   creator.Copyright,
			MainPackage: creator.Package + "/cmd/" + creator.Name,
		},
		Build: config.BuildConfig{
			OutputPath:     filepath.Join(common.BuildDirName + "/"),
			LDFlags:        []string{"-s", "-w"},
			LDFlagsWindows: []string{},
			LDFlagsDarwin:  []string{},
			LDFlagsLinux:   []string{},
			ExtraArgs:      []string{"-trimpath"},
		},
		Run: config.RunConfig{
			Environ: map[string]string{},
		},
		Package: config.PackageConfig{
			IncludeReadme: true,
		},
		Test: config.TestConfig{
			ExtraArgs: []string{},
		},
	}

	return creator.WriteJSONFileIfNotExists(configPath, cfg)

}

func (creator Creator) createTasks(project common.Project, cfg config.Config) error {

	tasks := newVisualStudioCodeTaskList(cfg, creator.CreateGitRepo)

	tasksPath := project.RelPath(visualStudioDirName, visualStudioCodeTasksFileName)
	return creator.WriteJSONFileIfNotExists(tasksPath, tasks)

}

func (creator Creator) createLaunchConfig(project common.Project, cfg config.Config) error {

	config := newVisualStudioCodeLaunchConfigs(cfg)

	launchConfigPath := project.RelPath(visualStudioDirName, visualStudioCodeLaunchFileName)
	return creator.WriteJSONFileIfNotExists(launchConfigPath, config)

}

func (creator Creator) createLicense(project common.Project, cfg config.Config) error {

	path := project.RelPath(common.LicenseFileName)
	return creator.WriteTextFileIfNotExists(path, apacheLicense)

}

func (creator Creator) createGitIgnore(project common.Project, cfg config.Config) error {

	gitIgnore := newGitIgnore(cfg)

	path := project.RelPath(common.GitIgnoreFileName)
	return creator.WriteTextFileIfNotExists(path, gitIgnore.string())

}

func (creator Creator) createDockerIgnore(project common.Project, cfg config.Config) error {

	gitIgnore := newDockerIgnore(cfg)

	path := project.RelPath(common.DockerIgnoreFileName)
	return creator.WriteTextFileIfNotExists(path, gitIgnore.string())

}

func (creator Creator) createDockerFile(project common.Project, cfg config.Config) error {

	gitIgnore := newDockerFile(cfg)

	path := project.RelPath(common.DockerfileFileName)
	return creator.WriteTextFileIfNotExists(path, gitIgnore.string())

}

func (creator Creator) createReadme(project common.Project, cfg config.Config) error {

	readme := newReadme(project, cfg)

	path := project.RelPath(common.ReadmeFileName)
	return creator.WriteTextFileIfNotExists(path, readme.markdownString())

}

func (creator Creator) createSourceFiles(project common.Project, cfg config.Config) error {

	packageName, err := project.Package()
	if err != nil {
		return err
	}

	filesToCreate := map[string]string{
		project.RelPath("library.go"):                                                                       mainLibTemplate,
		project.RelPath("library_test.go"):                                                                  mainLibTestingTemplate,
		project.RelPath(common.CmdDirName, filepath.Base(packageName), "main.go"):                           mainCmdTemplate,
		project.RelPath(common.VersionInfoPackage, common.VersionInfoFileName):                              versionInfoTemplate,
		project.RelPath(common.ScriptDirName, common.ScriptPreBuild, common.ScriptPreBuild+".example.go"):   preBuildScript,
		project.RelPath(common.ScriptDirName, common.ScriptPostBuild, common.ScriptPostBuild+".example.go"): postBuildScript,
	}

	for path, template := range filesToCreate {
		if err := creator.WriteTextTemplateIfNotExists(path, template, map[string]interface{}{
			"Project": project,
			"Config":  cfg,
		}); err != nil {
			return err
		}
	}

	return nil

}

func (creator Creator) createGoMod(project common.Project, cfg config.Config) error {

	packageName, err := project.Package()
	if err != nil {
		return err
	}

	goModPath := project.RelPath(common.GoModFileName)
	if creator.FileExists(goModPath) {
		return nil
	}

	env := map[string]string{"GO111MODULE": "on"}

	creator.LogPathCreation("Writing:", goModPath)
	cmd := []string{"go", "mod", "init", packageName}
	if output, err := creator.RunReturnOutput(cmd, project.Path, env); err != nil {
		creator.LogError(output)
		return err
	}

	return nil

}

func (creator Creator) createGitRepo(project common.Project, cfg config.Config) error {

	gitRepoPath := project.RelPath(common.GitRepoFileName)
	if creator.DirExists(gitRepoPath) {
		return nil
	}

	type command struct {
		description string
		cmdLine     []string
	}

	commandsToRun := []command{
		{
			description: "Creating git repo",
			cmdLine:     []string{"git", "init"},
		},
		{
			description: "Adding items to git repo",
			cmdLine:     []string{"git", "add", "."},
		},
		{
			description: "Committing git repo",
			cmdLine:     []string{"git", "commit", "-m", "Initial commit"},
		},
	}

	if strings.HasPrefix(creator.Package, "github.com/") {
		remoteRepository := "https://" + creator.Package + ".git"
		commandsToRun = append(commandsToRun,
			command{
				description: "Adding git remote: origin > " + remoteRepository,
				cmdLine:     []string{"git", "remote", "add", "origin", remoteRepository},
			},
		)
	}

	for _, cmd := range commandsToRun {
		creator.LogInfo(cmd.description)
		if output, err := creator.RunReturnOutput(cmd.cmdLine, project.Path, map[string]string{}); err != nil {
			creator.LogError(output)
			return err
		}
	}

	return nil

}
