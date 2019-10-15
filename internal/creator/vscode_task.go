package creator

const visualStudioDirName = ".vscode"
const visualStudioCodeTasksFileName = "tasks.json"

type visualStudioCodeTask struct {
	Label          string   `json:"label"`
	Command        string   `json:"command"`
	Type           string   `json:"type"`
	Group          string   `json:"group"`
	ProblemMatcher []string `json:"problemMatcher"`
}

func newVisualStudioCodeTask(label string, command string) visualStudioCodeTask {
	return visualStudioCodeTask{
		Label:          label,
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

func newVisualStudioCodeTaskList() *visualStudioCodeTasks {
	result := &visualStudioCodeTasks{
		Version: "2.0.0",
		Tasks: []visualStudioCodeTask{
			newVisualStudioCodeTask("build", "go-james build"),
			newVisualStudioCodeTask("build (verbose)", "go-james build -v"),
			newVisualStudioCodeTask("clean", "go-james clean"),
			newVisualStudioCodeTask("tests", "go-james test"),
			newVisualStudioCodeTask("run", "go-james run"),
			newVisualStudioCodeTask("install", "go-james install"),
			newVisualStudioCodeTask("uninstall", "go-james uninstall"),
			newVisualStudioCodeTask("run (debug)", "DEBUG=1 go-james run"),
		},
	}
	return result
}
