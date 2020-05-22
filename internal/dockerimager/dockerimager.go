package dockerimager

import (
	"errors"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

const dockerFileName = "Dockerfile"

// DockerImager implements the "docker-image" command
type DockerImager struct {
	common.CommandRunner
	common.FileSystem
}

// Execute executes the command
func (dockerImager DockerImager) Execute(project common.Project, cfg config.Config) error {

	dockerFilePath := project.RelPath(dockerFileName)

	if !dockerImager.FileExists(dockerFilePath) {
		return errors.New("No \"" + dockerFileName + "\" found in the project root")
	}

	runCmd := []string{"docker", "build", "-t", cfg.Project.Name, "."}
	return dockerImager.RunToStdout(runCmd, project.Path, cfg.Run.Environ)

}

// RequiresBuild indicates if a build is required before running the command
func (dockerImager DockerImager) RequiresBuild() bool {
	return false
}
