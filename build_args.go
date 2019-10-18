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
