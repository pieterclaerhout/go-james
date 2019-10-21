package creator

import (
	"strings"

	"github.com/pieterclaerhout/go-james/internal/common"
	"github.com/pieterclaerhout/go-james/internal/config"
)

type readme struct {
	Project common.Project
	Config  config.Config
}

func newReadme(project common.Project, cfg config.Config) readme {
	return readme{
		Project: project,
		Config:  cfg,
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

	for _, badge := range r.Project.Badges() {
		b.WriteString(badge.MarkdownString())
		b.WriteString(" ")
	}

	return strings.TrimSpace(b.String())

}
