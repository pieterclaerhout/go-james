package creator

import (
	"strings"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

type dockerIgnore struct {
	PatternsToIgnore []string
}

func newDockerIgnore(_ config.Config) dockerIgnore {
	return dockerIgnore{
		PatternsToIgnore: []string{
			".git",
			".gitignore",
			".vscode",
			".idea",
			"/" + common.BuildDirName,
		},
	}
}

func (g dockerIgnore) string() string {
	var b strings.Builder
	for _, pattern := range g.PatternsToIgnore {
		b.WriteString(pattern)
		b.WriteString("\n")
	}
	return b.String()
}
