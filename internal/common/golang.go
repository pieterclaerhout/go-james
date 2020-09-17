package common

import (
	"go/build"
	"os"
	"path/filepath"
)

// Golang is what can be injected into a subcommand when you need Go specific items
type Golang struct{}

// GoPath returns a relative path inside $GOPATH
func (g Golang) GoPath(subpath ...string) string {

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	subpath = append([]string{gopath}, subpath...)

	return filepath.Join(subpath...)

}

// GoBin returns a relative path inside $GOPATH/bin
func (g Golang) GoBin(subpath ...string) string {
	subpath = append([]string{"bin"}, subpath...)
	return g.GoPath(subpath...)
}

// IsDebug returns true of the DEBUG env var is set
func (g Golang) IsDebug() bool {
	return os.Getenv("DEBUG") != ""
}
