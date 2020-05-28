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

	imageNameWithTag := imageName + ":" + tag

	dockerImager.LogInfo("Building Docker Image:", imageNameWithTag)

	runCmd := []string{"docker", "build", "-t", imageNameWithTag, "."}
	if err := dockerImager.RunToStdout(runCmd, project.Path, cfg.Run.Environ); err != nil {
		return err
	}

	repository := cfg.DockerImage.Repository
	if repository != "" {

		dockerImager.LogInfo("Tagging Docker Image:", repository+":"+tag)

		tagCmd := []string{"docker", "tag", imageNameWithTag, repository + ":" + tag}
		if err := dockerImager.RunToStdout(tagCmd, project.Path, cfg.Run.Environ); err != nil {
			return err
		}

		dockerImager.LogInfo("Pushing Docker Image:", repository+":"+tag)

		pushCmd := []string{"docker", "push", repository + ":" + tag}
		if err := dockerImager.RunToStdout(pushCmd, project.Path, cfg.Run.Environ); err != nil {
			return err
		}

		cleanupCmd := []string{"docker", "image", "rm", imageNameWithTag}
		if err := dockerImager.RunToStdout(cleanupCmd, project.Path, cfg.Run.Environ); err != nil {
			return err
		}

	}

	if cfg.DockerImage.PruneAfterBuild {

		dockerImager.LogInfo("Pruning Docker Images")

		pruneCmd := []string{"docker", "image", "prune", "-f"}
		if _, err := dockerImager.RunReturnOutput(pruneCmd, project.Path, cfg.Run.Environ); err != nil {
			return err
		}

	}

	dockerImager.LogInfo("Docker Images after building")

	listCmd := []string{"docker", "image", "list", repository}
	if err := dockerImager.RunToStdout(listCmd, project.Path, cfg.Run.Environ); err != nil {
		return err
	}

	return nil

}

// RequiresBuild indicates if a build is required before running the command
func (dockerImager DockerImager) RequiresBuild() bool {
	return false
}
