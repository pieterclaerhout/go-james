package james

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

func (project Project) runCommand(cmdLine []string) error {

	var cmdPath string
	var cmdArgs []string

	switch len(cmdLine) {
	case 0:
		return ErrEmptyCommand
	case 1:
		cmdPath = cmdLine[0]
		cmdArgs = make([]string, 0)
	default:
		cmdPath = cmdLine[0]
		cmdArgs = cmdLine[1:]
	}

	log.Debug("Executing:", shellquote.Join(cmdLine...))

	command := exec.Command(cmdPath, cmdArgs...)
	command.Env = os.Environ()

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Start(); err != nil {
		return err
	}

	return command.Wait()

}
