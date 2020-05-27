package common

import (
	"strings"
)

// Version is used to get version information
type Version struct {
	CommandRunner
}

// Revision returns the Git revision in it's short format
func (version Version) Revision(project Project) string {

	// cmdLine := []string{"git", "rev-parse", "--short", "HEAD"}
	cmdLine := []string{"git", "describe", "--always", "--abbrev=6", "--dirty"}

	result, err := version.RunReturnOutput(cmdLine, project.Path, map[string]string{})
	if err != nil && strings.Contains(result, "not a git repository") {
		return ""
	}

	return strings.TrimSpace(result)

}

// BranchName returns the name of the current branch
func (version Version) BranchName(project Project) string {

	cmdLine := []string{"git", "rev-parse", "--abbrev-ref", "HEAD"}

	result, err := version.RunReturnOutput(cmdLine, project.Path, map[string]string{})
	if err != nil && strings.Contains(result, "not a git repository") {
		return ""
	}

	return strings.TrimSpace(result)

}
