package james

import (
	"github.com/pieterclaerhout/go-log"
)

func (project Project) DoRun() error {

	log.Debug("Running: run")
	log.Debug("Project path:", project.Path)

	return nil

}
