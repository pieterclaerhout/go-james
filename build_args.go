package james

import (
	"os"
)

// BuildArgs are the arguments which are passed to the pre_build.go and post_build.go scripts
type BuildArgs struct {
	ProjectPath        string   `json:"project_path,omitempty"`        // The root path of the project
	OutputPath         string   `json:"output_path,omitempty"`         // The absolute build path
	GOOS               string   `json:"GOOS,omitempty"`                // The GOOS for which was compiled
	GOARCH             string   `json:"GOARCH,omitempty"`              // The GOARCH for which was compiled
	ProjectName        string   `json:"project_name,omitempty"`        // The project name
	ProjectDescription string   `json:"project_description,omitempty"` // The project description
	ProjectCopyright   string   `json:"project_copyright,omitempty"`   // The project copyright
	Version            string   `json:"version,omitempty"`             // The version info
	Revision           string   `json:"revision,omitempty"`            // The project revision
	Branch             string   `json:"branch,omitempty"`              // The project branch
	RawBuildCommand    []string `json:"raw_build_command,omitempty"`   // The raw build command which is executed
}

// AsMap returns the build arguments as a map
func (b BuildArgs) AsMap() map[string]string {
	return map[string]string{
		"GO_JAMES_PROJECT_PATH":        b.ProjectPath,
		"GO_JAMES_OUTPUT_PATH":         b.OutputPath,
		"GO_JAMES_GOOS":                b.GOOS,
		"GO_JAMES_GOARCH":              b.GOARCH,
		"GO_JAMES_PROJECT_NAME":        b.ProjectName,
		"GO_JAMES_PROJECT_DESCRIPTION": b.ProjectDescription,
		"GO_JAMES_PROJECT_COPYRIGHT":   b.ProjectCopyright,
		"GO_JAMES_VERSION":             b.Version,
		"GO_JAMES_REVISION":            b.Revision,
		"GO_JAMES_BRANCH":              b.Branch,
	}
}

// ParseBuildArgs parses the build arguments from os.Args
func ParseBuildArgs() (BuildArgs, error) {
	return parseBuildArgsFromArguments(os.Args)
}

// parseBuildArgsFromArguments does heavy-lifting of the parsing
func parseBuildArgsFromArguments(args []string) (BuildArgs, error) {
	var buildArgs BuildArgs
	if err := parseArgsInto(args, &buildArgs); err != nil {
		return buildArgs, err
	}
	return buildArgs, nil
}
