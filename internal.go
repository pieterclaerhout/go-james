package james

import (
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"

	"github.com/kballard/go-shellquote"
	"github.com/pieterclaerhout/go-log"
)

var (
	// ErrEmptyCommand is returned when the command is empty
	ErrEmptyCommand = errors.New("No command specified")
)

func (project Project) createCommand(cmdLine []string) (*exec.Cmd, error) {

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
	command.Dir = project.Path

	return command, nil

}

func (project Project) runCommandToStdout(cmdLine []string) error {

	command, err := project.createCommand(cmdLine)
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

func (project Project) determineRevision() string {

	cmdLine := []string{"git", "rev-parse", "--short", "HEAD"}

	command, err := project.createCommand(cmdLine)
	if err != nil {
		if log.DebugMode {
			log.Error(err)
		}
		return ""
	}

	log.Debug("Executing:", shellquote.Join(cmdLine...))
	output, err := command.CombinedOutput()
	if err != nil {
		if log.DebugMode {
			log.Error(err)
		}
	}

	return strings.TrimSpace(string(output))

}

func (project Project) determineBranch() string {

	cmdLine := []string{"git", "rev-parse", "--abbrev-ref", "HEAD"}

	command, err := project.createCommand(cmdLine)
	if err != nil {
		if log.DebugMode {
			log.Error(err)
		}
		return ""
	}

	log.Debug("Executing:", shellquote.Join(cmdLine...))
	output, err := command.CombinedOutput()
	if err != nil {
		if log.DebugMode {
			log.Error(err)
		}
	}

	return strings.TrimSpace(string(output))

}
