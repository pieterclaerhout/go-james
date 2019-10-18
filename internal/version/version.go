package version

import (
	"fmt"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
	"github.com/pieterclaerhout/go-james/versioninfo"
)

// Version implements the "version" command
type Version struct {
}

// Execute executes the command
func (version Version) Execute(project common.Project, cfg config.Config) error {
	fmt.Println(versioninfo.ProjectName + " " + versioninfo.Version + " (" + versioninfo.Revision + ", " + versioninfo.Branch + ")")
	if versioninfo.ProjectDescription != "" {
		fmt.Println(versioninfo.ProjectDescription)
	}
	if versioninfo.ProjectCopyright != "" {
		fmt.Println("\n" + versioninfo.ProjectCopyright)
	}
	return nil
}

// RequiresBuild indicates if a build is required before running the command
func (version Version) RequiresBuild() bool {
	return false
}
