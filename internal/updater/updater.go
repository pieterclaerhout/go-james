package updater

import (
	"github.com/blang/semver"
	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
	"github.com/pieterclaerhout/go-james/versioninfo"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const repoName = "pieterclaerhout/go-james"

// Updater implements the "updater" command
type Updater struct {
	common.Logging
}

// Execute executes the command
func (updater Updater) Execute(project common.Project, cfg config.Config) error {

	v := semver.MustParse(versioninfo.Version)

	latest, err := selfupdate.UpdateSelf(v, repoName)
	if err != nil {
		return err
	}

	if latest.Version.Equals(v) {
		updater.LogInfo("You are running the latest version", versioninfo.Version)
	} else {
		updater.LogInfo("Successfully updated to version", latest.Version)
		updater.LogInfo("\nRelease notes:\n", latest.ReleaseNotes)
	}

	return nil

}

// RequiresBuild indicates if a build is required before running the command
func (updater Updater) RequiresBuild() bool {
	return false
}
