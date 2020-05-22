package creator

import (
	"path/filepath"

	"github.com/pieterclaerhout/go-james/internal/config"
)

const visualStudioCodeSettingsFileName = "settings.json"

type visualStudioCodeSettings struct {
	FilesExclude        map[string]bool `json:"files.exclude"`
	FilesWatcherExclude map[string]bool `json:"files.watcherExclude"`
	SearchExclude       map[string]bool `json:"search.exclude"`
}

func newVisualStudioCodeSettings(cfg config.Config) *visualStudioCodeSettings {

	buildFolderPath := filepath.Join(cfg.Build.OutputPath, cfg.Project.Name)
	buildFolderPath = filepath.Dir(buildFolderPath) // Trick to get the path with the trailing slash

	return &visualStudioCodeSettings{
		FilesExclude: map[string]bool{
			"**/.git":               true,
			"**/.svn":               true,
			"**/.hg":                true,
			"**/CVS":                true,
			"**/" + buildFolderPath: true,
			"**/.DS_Store":          true,
		},
		FilesWatcherExclude: map[string]bool{
			"**/" + buildFolderPath + "**": true,
		},
		SearchExclude: map[string]bool{
			"**/" + buildFolderPath + "": true,
		},
	}

}
