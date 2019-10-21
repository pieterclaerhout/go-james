package creator

import (
	"strings"

	"github.com/pieterclaerhout/go-james/internal/config"
)

type readme struct {
	Config config.Config
}

func newReadme(cfg config.Config) readme {
	return readme{
		Config: cfg,
	}
}

func (r readme) markdownString() string {

	var b strings.Builder

	b.WriteString("# " + r.Config.Project.Name + "\n")
	b.WriteString("\n")

	if r.Config.Project.Description != "" {
		b.WriteString(r.Config.Project.Description + "\n")
		b.WriteString("\n")
	}

	for _, badge := range r.Config.Badges() {
		b.WriteString(badge.MarkdownString())
		b.WriteString(" ")
	}

	return strings.TrimSpace(b.String())

}
