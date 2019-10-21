package common

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirkon/goproxy/gomod"
)

// Project defines a Go project based on Go modules
type Project struct {
	FileSystem

	Path        string
	packageName string
}

// NewProject returns a new Project instance
func NewProject(path string, packageName string) Project {

	if absPath, err := filepath.Abs(path); err == nil {
		path = absPath
	}

	return Project{
		Path:        path,
		packageName: packageName,
	}

}

// RelPath returns a relative path inside the project
func (project Project) RelPath(subpath ...string) string {
	fullpath := []string{}
	fullpath = append(fullpath, project.Path)
	fullpath = append(fullpath, subpath...)
	return filepath.Join(fullpath...)
}

// Package gets the main package of the project from the go.mod file
func (project Project) Package() (string, error) {

	if project.packageName != "" {
		return project.packageName, nil
	}

	goModPath := project.RelPath(GoModFileName)
	if !project.FileExists(goModPath) {
		return "", errors.New("Failed to determine the package name for this project")
	}

	b, err := ioutil.ReadFile(goModPath)
	if err != nil {
		return "", errors.Wrap(err, "Failed to read the go.mod file")
	}

	mod, err := gomod.Parse(goModPath, b)
	if err != nil {
		return "", errors.Wrap(err, "Failed to parse the go.mod file")
	}

	project.packageName = strings.TrimSuffix(mod.Name, "/")

	return project.packageName, nil

}

// Badges returns the list of badges which should be shown for the project
func (project Project) Badges() []Badge {

	badges := []Badge{}

	packageName, err := project.Package()
	if err != nil {
		return badges
	}
	relativePackageName := strings.TrimPrefix(packageName, "github.com/")

	goReportBadge := Badge{
		Title: "Go Report Card",
		Link:  "https://goreportcard.com/report/" + packageName,
		Image: "https://goreportcard.com/badge/" + packageName,
	}
	badges = append(badges, goReportBadge)

	docsBadge := Badge{
		Title: "Documentation",
		Link:  "http://godoc.org/" + packageName,
		Image: "https://godoc.org/" + packageName + "?status.svg",
	}
	badges = append(badges, docsBadge)

	if strings.HasPrefix(packageName, "github.com/") {

		licenseBadge := Badge{
			Title: "License",
			Link:  "https://" + packageName + "/raw/master/LICENSE",
			Image: "https://img.shields.io/badge/license-Apache%20v2-orange.svg",
		}
		badges = append(badges, licenseBadge)

		versionBadge := Badge{
			Title: "GitHub Version",
			Link:  "https://" + packageName + "/releases",
			Image: "https://img.shields.io/github/v/release/" + relativePackageName,
		}
		badges = append(badges, versionBadge)

		issuesBadge := Badge{
			Title: "GitHub issues",
			Link:  "https://" + packageName + "/issues",
			Image: "https://img.shields.io/github/issues/" + relativePackageName + ".svg",
		}
		badges = append(badges, issuesBadge)

		lastCommitBadge := Badge{
			Title: "GitHub last commit",
			Link:  "https://" + packageName,
			Image: "https://img.shields.io/github/last-commit/" + relativePackageName + ".svg",
		}
		badges = append(badges, lastCommitBadge)

	}

	return badges

}

// ShortPackageName returns the package name for the main package
func (project Project) ShortPackageName() string {

	packageName, err := project.Package()
	if err != nil {
		packageName = project.Path
	}

	return PackageNameToShort(packageName)

}
