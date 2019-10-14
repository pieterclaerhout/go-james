package creator

import (
	"strings"

	"github.com/pieterclaerhout/go-james/internal/config"
)

const gitIgnoreFileName = ".gitignore"

type gitIgnore struct {
	PatternsToIgnore []string
}

func newGitIgnore(cfg config.Config) gitIgnore {
	return gitIgnore{
		PatternsToIgnore: []string{"/build"},
	}
}

func (g gitIgnore) string() string {
	var b strings.Builder
	for _, pattern := range g.PatternsToIgnore {
		b.WriteString(pattern)
		b.WriteString("\n")
	}
	return b.String()
}
