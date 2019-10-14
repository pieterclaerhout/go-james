package james

import (
	"github.com/pieterclaerhout/go-log"
)

func (project Project) DoNew() error {

	log.Debug("Running: new")
	log.Debug("Project path:", project.Path)

	return nil

}
