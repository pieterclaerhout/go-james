package debugger

import (
	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

const delvePackagePath = "github.com/go-delve/delve/cmd/dlv"

// Debugger implements the "debug" command
type Debugger struct {
	common.CommandRunner
	common.Logging
	common.FileSystem
	common.Golang

	Args []string
}

// Execute executes the command
func (debugger Debugger) Execute(project common.Project, cfg config.Config) error {

	debugCmdPath := debugger.GoBin("dlv")

	if !debugger.FileExists(debugCmdPath) {
		debugger.LogPathCreation("Installing:", debugCmdPath)
	}

	env := map[string]string{}
	installCmd := []string{"go", "get", "-u", delvePackagePath}
	if output, err := debugger.RunReturnOutput(installCmd, project.Path, env); err != nil {
		debugger.LogError(output)
		return err
	}

	debugCmd := []string{debugCmdPath, "debug", cfg.Project.MainPackage}
	debugCmd = append(debugCmd, debugger.Args...)

	return debugger.RunInteractive(debugCmd, project.Path, map[string]string{})

}

// RequiresBuild indicates if a build is required before running the command
func (debugger Debugger) RequiresBuild() bool {
	return false
}
