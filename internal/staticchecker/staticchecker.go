package staticchecker

import (
	"strings"

	"github.com/kballard/go-shellquote"
	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

const staticcheckPackagePath = "honnef.co/go/tools/cmd/staticcheck"

// StaticChecker implements the "staticcheck" command
type StaticChecker struct {
	common.CommandRunner
	common.FileSystem
	common.Golang
}

// Execute executes the command
func (staticChecker StaticChecker) Execute(project common.Project, cfg config.Config) error {

	staticcheckCmdPath := staticChecker.GoBin("staticcheck")

	if !staticChecker.FileExists(staticcheckCmdPath) {
		staticChecker.LogPathCreation("Installing:", staticcheckCmdPath)
	}

	env := map[string]string{}
	installCmd := []string{"go", "install", staticcheckPackagePath}
	if output, err := staticChecker.RunReturnOutput(installCmd, project.Path, env); err != nil {
		staticChecker.LogError(output)
		return err
	}

	staticcheckCmd := []string{staticcheckCmdPath}
	if len(cfg.Staticcheck.Checks) > 0 {
		staticcheckCmd = append(staticcheckCmd, "-checks")
		staticcheckCmd = append(staticcheckCmd, strings.Join(cfg.Staticcheck.Checks, ","))
	}
	staticcheckCmd = append(staticcheckCmd, "./...")

	staticChecker.LogInfo("> Running: staticcheck", shellquote.Join(staticcheckCmd[1:]...))

	err := staticChecker.RunToStdout(staticcheckCmd, project.Path, map[string]string{})
	if err != nil && err.Error() != "exit status 1" {
		return err
	}

	return nil

}

// RequiresBuild indicates if a build is required before running the command
func (staticChecker StaticChecker) RequiresBuild() bool {
	return false
}
