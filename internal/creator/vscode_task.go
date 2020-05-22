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
	Version string                 `json:"version"`
	Tasks   []visualStudioCodeTask `json:"tasks"`
}

func newVisualStudioCodeTaskList(cfg config.Config, createGitRepo bool) *visualStudioCodeTasks {
	result := &visualStudioCodeTasks{
		Version: "2.0.0",
		Tasks: []visualStudioCodeTask{
			newVisualStudioCodeTask(cfg, "build", "go-james build -v"),
			newVisualStudioCodeTask(cfg, "build darwin/amd64", "go-james build -v --goos=darwin --goarch=amd64"),
			newVisualStudioCodeTask(cfg, "build linux/amd64", "go-james build -v --goos=linux --goarch=amd64"),
			newVisualStudioCodeTask(cfg, "build windows/amd64", "go-james build -v --goos=windows --goarch=amd64"),
			newVisualStudioCodeTask(cfg, "clean", "go-james clean"),
			newVisualStudioCodeTask(cfg, "tests", "go-james test"),
			newVisualStudioCodeTask(cfg, "run", "go-james run"),
			newVisualStudioCodeTask(cfg, "staticcheck", "go-james staticcheck"),
			newVisualStudioCodeTask(cfg, "docker-image", "go-james docker-image"),
			newVisualStudioCodeTask(cfg, "install", "go-james install"),
			newVisualStudioCodeTask(cfg, "uninstall", "go-james uninstall"),
			newVisualStudioCodeTask(cfg, "run (debug)", "DEBUG=1 go-james run"),
		},
	}
	if createGitRepo {
		result.Tasks = append(result.Tasks,
			newVisualStudioCodeTask(cfg, "push to github", "git push --set-upstream origin"),
		)
	}
	return result
}
