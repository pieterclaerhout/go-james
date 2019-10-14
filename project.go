package james

import (
	"path/filepath"
)

type Project struct {
	Path string
}

func NewProject(path string) Project {

	if absPath, err := filepath.Abs(path); err == nil {
		path = absPath
	}

	return Project{
		Path: path,
	}

}
