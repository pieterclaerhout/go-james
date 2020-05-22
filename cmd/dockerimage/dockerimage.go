package dockerimage

import (
	"github.com/pieterclaerhout/go-james/internal"
	"github.com/pieterclaerhout/go-james/internal/dockerimager"
	"github.com/tucnak/climax"
)

// DockerImageCmd defines the docker-image command
var DockerImageCmd = climax.Command{
	Name:  "docker-image",
	Brief: "Builds the project as a docker image",
	Help:  "Builds the project as a docker image",
	Flags: []climax.Flag{},
	Handle: func(ctx climax.Context) int {

		tool := dockerimager.DockerImager{}

		executor := internal.NewExecutor("")
		return executor.RunTool(tool, true)

	},
}
