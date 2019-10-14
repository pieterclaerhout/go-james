package common

import (
	"os"
	"os/exec"

	"github.com/pkg/errors"

	"github.com/kballard/go-shellquote"
	"github.com/pieterclaerhout/go-log"
)

var (
	// ErrEmptyCommand is returned when the command is empty
	ErrEmptyCommand = errors.New("No command specified")
)

// CommandRunner is what can be injected into a subcommand when you need to run system commands
type CommandRunner struct{}

func (commandRunner CommandRunner) createCommand(cmdLine []string, workdir string) (*exec.Cmd, error) {

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
	command.Dir = workdir

	return command, nil

}

// RunToStdout runs the command and outputs the result to stdout/stderr
func (commandRunner CommandRunner) RunToStdout(cmdLine []string, workdir string) error {

	command, err := commandRunner.createCommand(cmdLine, workdir)
	if err != nil {
		return err
	}

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	log.Debug("Executing:", shellquote.Join(cmdLine...))
	if err := command.Start(); err != nil {
		return err
	}

	return command.Wait()

}

// RunReturnOutput runs the command and returns the result as a string
func (commandRunner CommandRunner) RunReturnOutput(cmdLine []string, workdir string) (string, error) {

	command, err := commandRunner.createCommand(cmdLine, workdir)
	if err != nil {
		return "", err
	}

	log.Debug("Executing:", shellquote.Join(cmdLine...))
	output, err := command.CombinedOutput()
	if err != nil {
		if log.DebugMode {
			log.Error(err)
		}
	}

	return string(output), err

}
