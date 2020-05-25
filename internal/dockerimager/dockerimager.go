package dockerimager

import (
	"errors"
	"strings"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

const dockerFileName = "Dockerfile"

// DockerImager implements the "docker-image" command
type DockerImager struct {
	common.CommandRunner
	common.FileSystem
	common.Logging
	common.Version
}

// Execute executes the command
func (dockerImager DockerImager) Execute(project common.Project, cfg config.Config) error {

	dockerFilePath := project.RelPath(dockerFileName)

	if !dockerImager.FileExists(dockerFilePath) {
		return errors.New(dockerFileName + "\" was not found in the project root")
	}

	imageName := cfg.DockerImage.Name
	if imageName == "" {
		imageName = cfg.Project.Name
	}

	if imageName == "" {
		return errors.New("docker-image name and project name should not both be empty")
	}

	tag := cfg.Project.Version
	if strings.EqualFold(cfg.DockerImage.Tag, "revision") {
		tag = dockerImager.Revision(project)
	}

	imageName += ":" + tag

	dockerImager.LogInfo("Building Docker Image:", imageName)

	runCmd := []string{"docker", "build", "-t", imageName, "."}
	if err := dockerImager.RunToStdout(runCmd, project.Path, cfg.Run.Environ); err != nil {
		return err
	}

	repository := cfg.DockerImage.Repository
	if repository != "" {

		dockerImager.LogInfo("Tagging Docker Image:", imageName, "as:", repository+":"+tag)

		tagCmd := []string{"docker", "tag", imageName, repository + ":" + tag}
		if err := dockerImager.RunToStdout(tagCmd, project.Path, cfg.Run.Environ); err != nil {
			return err
		}

		dockerImager.LogInfo("Pushing Docker Image:", repository+":"+tag)

		pushCmd := []string{"docker", "push", repository + ":" + tag}
		if err := dockerImager.RunToStdout(pushCmd, project.Path, cfg.Run.Environ); err != nil {
			return err
		}

		// docker tag geoip-server pieterclaerhout/geoip-server:$(REVISION)
		// docker push pieterclaerhout/geoip-server:$(REVISION)

	}

	return nil

}

// RequiresBuild indicates if a build is required before running the command
func (dockerImager DockerImager) RequiresBuild() bool {
	return false
}
