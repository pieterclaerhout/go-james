package james

import (
	"github.com/pieterclaerhout/go-log"
)

func (project Project) DoInit() error {

	log.Debug("Running: init")
	log.Debug("Project path:", project.Path)

	return nil

}
