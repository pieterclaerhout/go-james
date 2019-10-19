package creator

import (
	"github.com/pieterclaerhout/go-james/internal/config"
)

const visualStudioCodeTasksFileName = "tasks.json"

type visualStudioCodeTask struct {
	Label          string   `json:"label"`
	Command        string   `json:"command"`
	Type           string   `json:"type"`
	Group          string   `json:"group"`
	ProblemMatcher []string `json:"problemMatcher"`
}

func newVisualStudioCodeTask(cfg config.Config, label string, command string) visualStudioCodeTask {
	return visualStudioCodeTask{
		Label:          cfg.Project.Name + " | " + label,
		Command:        command,
		Type:           "shell",
		Group:          "build",
		ProblemMatcher: []string{"$go"},
	}
}

type visualStudioCodeTasks struct {
	Version string
	Tasks   []visualStudioCodeTask `json:"tasks"`
}

func newVisualStudioCodeTaskList(cfg config.Config) *visualStudioCodeTasks {
	result := &visualStudioCodeTasks{
		Version: "2.0.0",
		Tasks: []visualStudioCodeTask{
			newVisualStudioCodeTask(cfg, "build", "go-james build"),
			newVisualStudioCodeTask(cfg, "build (verbose)", "go-james build -v"),
			newVisualStudioCodeTask(cfg, "clean", "go-james clean"),
			newVisualStudioCodeTask(cfg, "tests", "go-james test"),
			newVisualStudioCodeTask(cfg, "run", "go-james run"),
			newVisualStudioCodeTask(cfg, "install", "go-james install"),
			newVisualStudioCodeTask(cfg, "uninstall", "go-james uninstall"),
			newVisualStudioCodeTask(cfg, "run (debug)", "DEBUG=1 go-james run"),
		},
	}
	return result
}
