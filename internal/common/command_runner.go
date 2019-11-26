package common

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/kballard/go-shellquote"
	"github.com/pieterclaerhout/go-log"
	"github.com/pkg/errors"
)

var (
	// ErrEmptyCommand is returned when the command is empty
	ErrEmptyCommand = errors.New("No command specified")
)

// CommandRunner is what can be injected into a subcommand when you need to run system commands
type CommandRunner struct {
	FileSystem
	Encoding
}

// createCommand creates the command instance
func (commandRunner CommandRunner) createCommand(cmdLine []string, workdir string, env map[string]string) (*exec.Cmd, error) {

	var cmdPath string
	var cmdArgs []string

	switch len(cmdLine) {
	case 0:
		return nil, ErrEmptyCommand
	case 1:
		cmdPath = cmdLine[0]
		cmdArgs = make([]string, 0)
	default:
		cmdPath = cmdLine[0]
		cmdArgs = cmdLine[1:]
	}

	command := exec.Command(cmdPath, cmdArgs...)
	command.Env = os.Environ()
	if workdir != "" {
		command.Dir = workdir
	}

	for key, val := range env {
		command.Env = append(command.Env, fmt.Sprintf("%s=%s", key, val))
	}

	return command, nil

}

// RunInteractive and interactive command
func (commandRunner CommandRunner) RunInteractive(cmdLine []string, workdir string, env map[string]string) error {

	command, err := commandRunner.createCommand(cmdLine, workdir, env)
	if err != nil {
		return err
	}

	command.Stdout = os.Stdout
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr

	return command.Run()

}

// RunToStdout runs the command and outputs the result to stdout/stderr
func (commandRunner CommandRunner) RunToStdout(cmdLine []string, workdir string, env map[string]string) error {

	command, err := commandRunner.createCommand(cmdLine, workdir, env)
	if err != nil {
		return err
	}

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	commandRunner.logCommand(cmdLine, env)

	if err := command.Start(); err != nil {
		return err
	}

	return command.Wait()

}

// RunReturnOutput runs the command and returns the result as a string
func (commandRunner CommandRunner) RunReturnOutput(cmdLine []string, workdir string, env map[string]string) (string, error) {

	command, err := commandRunner.createCommand(cmdLine, workdir, env)
	if err != nil {
		return "", err
	}

	commandRunner.logCommand(cmdLine, env)

	output, err := command.CombinedOutput()
	if err != nil && log.DebugMode {
		log.Error(err)
	}

	return string(output), err

}

// RunProjectHook runs the given hook in the context of the project passing the args to ti
func (commandRunner CommandRunner) RunProjectHook(project Project, hookName string, args interface{}) error {
	return commandRunner.RunScriptIfExistsToStdout(
		project.RelPath(ScriptDirName, hookName, hookName+".go"), args, project.Path, map[string]string{},
	)
}

// RunScriptToStdout runs the given script using "go run" and outputs the result to stdout/stderr
func (commandRunner CommandRunner) RunScriptToStdout(scriptPath string, args interface{}, workdir string, env map[string]string) error {
	log.Debug("Running script:", scriptPath)
	cmdLine := commandRunner.cmdLineForScriptWithArgs(scriptPath, args)
	return commandRunner.RunToStdout(cmdLine, workdir, env)
}

// RunScriptIfExistsToStdout runs the given script if it exists using "go run" and outputs the result to stdout/stderr
func (commandRunner CommandRunner) RunScriptIfExistsToStdout(scriptPath string, args interface{}, workdir string, env map[string]string) error {
	if !commandRunner.FileExists(scriptPath) {
		return nil
	}
	return commandRunner.RunScriptToStdout(scriptPath, args, workdir, env)
}

// RunScriptReturnOutput runs the given script using "go run" and outputs the result to stdout/stderr
func (commandRunner CommandRunner) RunScriptReturnOutput(scriptPath string, args interface{}, workdir string, env map[string]string) (string, error) {
	log.Debug("Running script:", scriptPath)
	cmdLine := commandRunner.cmdLineForScriptWithArgs(scriptPath, args)
	return commandRunner.RunReturnOutput(cmdLine, workdir, env)
}

// RunScriptIfExistsReturnOutput runs the given script if it exists using "go run" and outputs the result to stdout/stderr
func (commandRunner CommandRunner) RunScriptIfExistsReturnOutput(scriptPath string, args interface{}, workdir string, env map[string]string) (string, error) {
	if !commandRunner.FileExists(scriptPath) {
		return "", nil
	}
	return commandRunner.RunScriptReturnOutput(scriptPath, args, workdir, env)
}

func (commandRunner CommandRunner) logCommand(cmdLine []string, env map[string]string) {
	log.Debug("Executing:", shellquote.Join(cmdLine...))
	if len(env) > 0 {
		log.DebugDump(env, "Environment:")
	}
}

func (commandRunner CommandRunner) cmdLineForScriptWithArgs(script string, args interface{}) []string {
	cmdLine := []string{"go", "run", script}
	cmdLine = append(cmdLine, commandRunner.ToJSONString(args))
	return cmdLine
}
