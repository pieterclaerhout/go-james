package common

import "strings"

// Badge is used to link with a badge to a specific URL
type Badge struct {
	Title string // The title of the badge
	Link  string // The link to link to
	Image string // The src of the image for the badge
}

// MarkdownString returns the badge as a Markdown string
func (badge Badge) MarkdownString() string {
	var b strings.Builder
	b.WriteString("[![")
	b.WriteString(badge.Title)
	b.WriteString("](")
	b.WriteString(badge.Image)
	b.WriteString(")](")
	b.WriteString(badge.Link)
	b.WriteString(")")
	return b.String()
}
