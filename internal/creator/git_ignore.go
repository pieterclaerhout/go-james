package creator

import (
	"strings"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

type gitIgnore struct {
	PatternsToIgnore []string
}

func newGitIgnore(_ config.Config) gitIgnore {
	return gitIgnore{
		PatternsToIgnore: []string{"/" + common.BuildDirName},
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
